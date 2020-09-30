"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import ipaddress
import logging

import grpc
from lte.protos.mobilityd_pb2 import AllocateIPRequest, IPAddress, IPBlock, \
    ListAddedIPBlocksResponse, ListAllocatedIPsResponse, RemoveIPBlockResponse, \
    SubscriberIPTable, GWInfo, ListGWInfoResponse, AllocateIPAddressResponse
from lte.protos.mobilityd_pb2_grpc import MobilityServiceServicer, \
    add_MobilityServiceServicer_to_server
from lte.protos.subscriberdb_pb2 import SubscriberID
from magma.common.rpc_utils import return_void
from magma.mobilityd.ip_descriptor import IPDesc
from magma.subscriberdb.sid import SIDUtils
from .ip_address_man import IPAddressManager, IPNotInUseError, \
    MappingNotFoundError
from .ip_allocator_base import DuplicateIPAssignmentError, \
    DuplicatedIPAllocationError, IPBlockNotFoundError, NoAvailableIPError, \
    OverlappedIPBlocksError

from .ip_allocator_base import DuplicatedIPAllocationError
from .ipv6_allocator_pool import IPv6AllocatorPool, MaxRandIIDCalculationError


def _get_ip_block(ip_block_str):
    """ Convert string into ipaddress.ip_network. Support both IPv4 or IPv6
    addresses.

        Args:
            ip_block_str(string): network address, e.g. "192.168.0.0/24".

        Returns:
            ip_block(ipaddress.ip_network)
    """
    try:
        ip_block = ipaddress.ip_network(ip_block_str)
    except ValueError:
        logging.error("Invalid IP block format: %s", ip_block_str)
        return None
    return ip_block


class MobilityServiceRpcServicer(MobilityServiceServicer):
    """ gRPC based server for the IPAllocator. """

    def __init__(self, mconfig, config, subscriberdb_rpc_stub=None):
        # TODO: consider adding gateway mconfig to decide whether to
        # persist to Redis

        self._ip_address_man = IPAddressManager(config=config,
                                                subscriberdb_rpc_stub=subscriberdb_rpc_stub,
                                                mconfig=mconfig)

        # Load IPv4 block from the configurable mconfig file
        # No dynamic reloading support for now, assume restart after updates
        ipv4_block = _get_ip_block(mconfig.ip_block)
        if ipv4_block is not None:
            try:
                self.add_ip_block(ipv4_block)
            except OverlappedIPBlocksError:
                logging.error("Overlapped IPv4 block: %s", ipv4_block)

        # TODO: Add IPv6 allocator to IP Address manager, for now support
        # POOL allocation
        self._ipv6_allocator = IPv6AllocatorPool(config=config)
        ipv6_block = _get_ip_block(config['ipv6_prefix_block'])
        if ipv6_block is not None:
            try:
                self.add_ip_block(ipv6_block)
            except OverlappedIPBlocksError:
                logging.error("Overlapped IPv6 block: %s", ipv6_block)

    def add_to_server(self, server):
        """ Add the servicer to a gRPC server """
        add_MobilityServiceServicer_to_server(self, server)

    def add_ip_block(self, ip_block):
        """ Add IP block to the IP allocator.

            Args:
                ip_block (ipaddress.ip_network): ip network to add
                e.g. ipaddress.ip_network("10.0.0.0/24")

            Raise:
                OverlappedIPBlocksError: if the given IP block overlaps with
                existing ones
                IPVersionNotSupportedError: if given IP version of the IP block
                is not supported
        """
        if ip_block.version == 4:
            self._ipv4_allocator.add_ip_block(ip_block)
            logging.info("Added block %s to the IPv4 address pool", ip_block)
        elif ip_block.version == 6:
            logging.info("Assigned IPv6 block: %s", ip_block)
            self._ipv6_allocator.add_ip_block(ip_block)

    @return_void
    def AddIPBlock(self, ipblock_msg, context):
        """ Add a range of IP addresses into the free IP pool

        Args:
            ipblock_msg (IPBlock): ip block to add. ipblock_msg has the
            type IPBlock, a protobuf message type for the gRPC interface.
            Internal representation of ip blocks uses the ipaddress.ip_network
            type and is named as ipblock.
        """
        ipblock = self._ipblock_msg_to_ipblock(ipblock_msg, context)
        if ipblock is None:
            return

        try:
            self.add_ip_block(ipblock)
        except OverlappedIPBlocksError:
            context.set_details('Overlapped ip block: %s' % ipblock)
            context.set_code(grpc.StatusCode.FAILED_PRECONDITION)
        except IPVersionNotSupportedError:
            self._unimplemented_ip_version_error(context)

    def ListAddedIPv4Blocks(self, void, context):
        """ Return a list of IPv4 blocks assigned """
        resp = ListAddedIPBlocksResponse()

        ip_blocks = self._ip_address_man.list_added_ip_blocks()
        ip_block_msg_list = [IPBlock(version=IPAddress.IPV4,
                                     net_address=block.network_address.packed,
                                     prefix_len=block.prefixlen)
                             for block in ip_blocks]
        resp.ip_block_list.extend(ip_block_msg_list)

        return resp

    def ListAllocatedIPs(self, ipblock_msg, context):
        """ Return a list of IPs allocated from a IP block

        Args:
            ipblock_msg (IPBlock): ip block to add. ipblock_msg has the
            type IPBlock, a protobuf message type for the gRPC interface.
            Internal representation of ip blocks uses the ipaddress.ip_network
            type and is named as ipblock.
        """
        resp = ListAllocatedIPsResponse()

        ipblock = self._ipblock_msg_to_ipblock(ipblock_msg, context)
        if ipblock is None:
            return resp

        if ipblock_msg.version == IPBlock.IPV4:
            try:
                ips = self._ip_address_man.list_allocated_ips(ipblock)
                ip_msg_list = [IPAddress(version=IPAddress.IPV4,
                                         address=ip.packed) for ip in ips]

                resp.ip_list.extend(ip_msg_list)
            except IPBlockNotFoundError:
                context.set_details('IP block not found: %s' % ipblock)
                context.set_code(grpc.StatusCode.FAILED_PRECONDITION)
        else:
            self._unimplemented_ip_version_error(context)

        return resp

    def AllocateIPAddress(self, request, context):
        """ Allocate an IP address from the free IP pool """
        ip_addr = IPAddress()
        composite_sid = SIDUtils.to_str(request.sid)
        if request.apn:
            composite_sid = composite_sid + "." + request.apn

        if request.version == AllocateIPRequest.IPV4:
            try:
                ip, vlan = self._ipv4_allocator.alloc_ip_address(composite_sid)
                logging.info("Allocated IPv4 %s for sid %s for apn %s"
                             % (ip, SIDUtils.to_str(request.sid), request.apn))
                ip_addr.version = IPAddress.IPV4
                ip_addr.address = ip.packed
                return AllocateIPAddressResponse(ip_addr=ip_addr,
                                                 vlan=str(vlan))
            except NoAvailableIPError:
                context.set_details('No free IPv4 IP available')
                context.set_code(grpc.StatusCode.RESOURCE_EXHAUSTED)
            except DuplicatedIPAllocationError:
                context.set_details('IP has been allocated for this subscriber')
                context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            except DuplicateIPAssignmentError:
                context.set_details('IP has been allocated for other subscriber')
                context.set_code(grpc.StatusCode.ALREADY_EXISTS)
        elif request.version == AllocateIPRequest.IPV6:
            try:
                ip = self._ipv6_allocator.alloc_ip_address(composite_sid, 0)
                logging.info("Allocated IPv6 %s for sid %s for apn %s"
                             % (ip, SIDUtils.to_str(request.sid), request.apn))
                ip_addr.version = IPAddress.IPV6
                ip_addr.address = ip.packed
            except MaxRandIIDCalculationError:
                context.set_details('Could not calculate IID for IPv6 address')
                context.set_code(grpc.StatusCode.RESOURCE_EXHAUSTED)
            return AllocateIPAddressResponse(ip_addr=ip_addr, vlan="")
        return AllocateIPAddressResponse()

    @return_void
    def ReleaseIPAddress(self, request, context):
        """ Release an allocated IP address """

        ip = ipaddress.ip_address(request.ip.address)
        composite_sid = SIDUtils.to_str(request.sid)
        if request.apn:
            composite_sid = composite_sid + "." + request.apn

        if request.ip.version == IPAddress.IPV4:
            try:
                self._ipv4_allocator.release_ip_address(
                    composite_sid, ip)
                logging.info("Released IPv4 %s for sid %s"
                             % (ip, SIDUtils.to_str(request.sid)))
            except IPNotInUseError:
                context.set_details('IP %s not in use' % ip)
                context.set_code(grpc.StatusCode.NOT_FOUND)
            except MappingNotFoundError:
                context.set_details('(SID, IP) map not found: (%s, %s)'
                                    % (SIDUtils.to_str(request.sid), ip))
                context.set_code(grpc.StatusCode.NOT_FOUND)
        elif request.ip.version == IPAddress.IPV6:
            ip_desc = IPDesc(ip=ip, sid=composite_sid)
            try:
                self._ipv6_allocator.release_ip(ip_desc)
                logging.info("Released IPv6 %s for sid %s"
                             % (ip, SIDUtils.to_str(request.sid)))
            except IPNotInUseError:
                context.set_details('IP %s not in use' % ip)
                context.set_code(grpc.StatusCode.NOT_FOUND)

    def RemoveIPBlock(self, request, context):
        """ Attempt to remove IP blocks and return the removed blocks """
        if request.version == IPAddress.IPV4:
            allocator = self._ipv4_allocator
            version = IPAddress.IPV4
        else:
            allocator = self._ipv6_allocator
            version = IPAddress.IPV6

        removed_blocks = allocator.remove_ip_blocks(
            *[self._ipblock_msg_to_ipblock(ipblock_msg, context)
              for ipblock_msg in request.ip_blocks],
            force=request.force)
        removed_block_msgs = [IPBlock(version=version,
                                      net_address=block.network_address.packed,
                                      prefix_len=block.prefixlen)
                                      for block in removed_blocks]

        resp = RemoveIPBlockResponse()
        resp.ip_blocks.extend(removed_block_msgs)
        return resp

    def GetIPForSubscriber(self, request, context):
        composite_sid = SIDUtils.to_str(request.sid)
        if request.apn:
            composite_sid = composite_sid + "." + request.apn

        if request.version == IPAddress.IPV4:
            composite_sid += ",ipv4"
        elif request.version == IPAddress.IPV6:
            composite_sid += ",ipv6"

        ip = self._ip_address_man.get_ip_for_sid(composite_sid)
        if ip is None:
            context.set_details('SID %s not found'
                                % SIDUtils.to_str(request.sid))
            context.set_code(grpc.StatusCode.NOT_FOUND)
            return IPAddress()

        version = IPAddress.IPV4 if ip.version == 4 else IPAddress.IPV6
        return IPAddress(version=version, address=ip.packed)

    def GetSubscriberIDFromIP(self, ip_addr, context):
        sent_ip = ipaddress.ip_address(ip_addr.address)
        sid = self._ip_address_man.get_sid_for_ip(sent_ip)

        if sid is None:
            context.set_details('IP address %s not found' % str(sent_ip))
            context.set_code(grpc.StatusCode.NOT_FOUND)
            return SubscriberID()
        else:
            # handle composite key case
            sid, *rest = sid.partition('.')
            return SIDUtils.to_pb(sid)

    def GetSubscriberIPTable(self, void, context):
        """ Get the full subscriber table """
        logging.debug("Listing subscriber IP table")
        resp = SubscriberIPTable()

        csid_ip_pairs = self._ip_address_man.get_sid_ip_table()
        for composite_sid, ip in csid_ip_pairs:
            # handle composite sid to sid and apn mapping
            sid, _, apn_part = composite_sid.partition('.')
            apn, _ = apn_part.split(',')
            sid_pb = SIDUtils.to_pb(sid)
            version = IPAddress.IPV4 if ip.version == 4 else IPAddress.IPV6
            ip_msg = IPAddress(version=version, address=ip.packed)
            resp.entries.add(sid=sid_pb, ip=ip_msg, apn=apn)
        return resp

    def ListGatewayInfo(self, void, context):
        resp = ListGWInfoResponse()
        gw_info_list = self._ip_address_man.list_gateway_info()
        if gw_info_list:
            resp.gw_list.extend(gw_info_list)
        return resp

    @return_void
    def SetGatewayInfo(self, info, context):
        self._ip_address_man.set_gateway_info(info)

    def _get_allocate_ipv4_response(self, composite_sid, context, ip_addr,
                                    request):
        try:
            ip, vlan = self._ip_address_man.alloc_ip_address(composite_sid)
            logging.info("Allocated IPv4 %s for sid %s for apn %s"
                         % (ip, SIDUtils.to_str(request.sid), request.apn))
            ip_addr.version = IPAddress.IPV4
            ip_addr.address = ip.packed
            return AllocateIPAddressResponse(ip_list=[ip_addr],
                                             vlan=str(vlan))
        except NoAvailableIPError:
            context.set_details('No free IPv4 IP available')
            context.set_code(grpc.StatusCode.RESOURCE_EXHAUSTED)
        except DuplicatedIPAllocationError:
            context.set_details(
                'IP has been allocated for this subscriber')
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
        except DuplicateIPAssignmentError:
            context.set_details(
                'IP has been allocated for other subscriber')
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
        return AllocateIPAddressResponse()

    def _get_allocate_ipv6_response(self, composite_sid, context, ip_addr,
                                    request):
        try:
            ip = self._ip_address_man.alloc_ipv6_address(composite_sid)
            logging.info("Allocated IPv6 %s for sid %s for apn %s"
                         % (ip, SIDUtils.to_str(request.sid), request.apn))
            ip_addr.version = IPAddress.IPV6
            ip_addr.address = ip.packed
            return AllocateIPAddressResponse(ip_list=[ip_addr], vlan="")
        except MaxCalculationError:
            context.set_details('Could not calculate IID for IPv6 address')
            context.set_code(grpc.StatusCode.RESOURCE_EXHAUSTED)
        return AllocateIPAddressResponse()

    def _ipblock_msg_to_ipblock(self, ipblock_msg, context):
        """ convert IPBlock to ipaddress.ip_network """
        try:
            ip = ipaddress.ip_address(ipblock_msg.net_address)
            subnet = "%s/%s" % (str(ip), ipblock_msg.prefix_len)
            ipblock = ipaddress.ip_network(subnet)
            return ipblock
        except ValueError:
            context.set_details('Invalid IPBlock format: version=%s,'
                                'net_address=%s, prefix_len=%s' %
                                (ipblock_msg.version, ipblock_msg.net_address,
                                 ipblock_msg.prefix_len))
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            return None

    def _unimplemented_ip_version_error(self, context):
        context.set_details("IPv6 is not yet supported")
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)


class IPVersionNotSupportedError(Exception):
    """ Exception thrown when an IP version is not supported """
    pass


class UnknownIPAllocatorError(Exception):
    """ Exception thrown when an IP version is not supported """
    pass

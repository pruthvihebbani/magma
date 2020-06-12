// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cellular_service.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// (0)Turning off NonEPS service, (1)Both CSFB and SMS, (2)only SMS
type GatewayNonEPSConfig_NonEPSServiceControl int32

const (
	GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_OFF      GatewayNonEPSConfig_NonEPSServiceControl = 0
	GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_CSFB_SMS GatewayNonEPSConfig_NonEPSServiceControl = 1
	GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_SMS      GatewayNonEPSConfig_NonEPSServiceControl = 2
)

var GatewayNonEPSConfig_NonEPSServiceControl_name = map[int32]string{
	0: "NON_EPS_SERVICE_CONTROL_OFF",
	1: "NON_EPS_SERVICE_CONTROL_CSFB_SMS",
	2: "NON_EPS_SERVICE_CONTROL_SMS",
}

var GatewayNonEPSConfig_NonEPSServiceControl_value = map[string]int32{
	"NON_EPS_SERVICE_CONTROL_OFF":      0,
	"NON_EPS_SERVICE_CONTROL_CSFB_SMS": 1,
	"NON_EPS_SERVICE_CONTROL_SMS":      2,
}

func (x GatewayNonEPSConfig_NonEPSServiceControl) String() string {
	return proto.EnumName(GatewayNonEPSConfig_NonEPSServiceControl_name, int32(x))
}

func (GatewayNonEPSConfig_NonEPSServiceControl) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{3, 0}
}

type GatewayNonEPSConfig_CSFBRat int32

const (
	GatewayNonEPSConfig_CSFBRAT_2G GatewayNonEPSConfig_CSFBRat = 0
	GatewayNonEPSConfig_CSFBRAT_3G GatewayNonEPSConfig_CSFBRat = 1
)

var GatewayNonEPSConfig_CSFBRat_name = map[int32]string{
	0: "CSFBRAT_2G",
	1: "CSFBRAT_3G",
}

var GatewayNonEPSConfig_CSFBRat_value = map[string]int32{
	"CSFBRAT_2G": 0,
	"CSFBRAT_3G": 1,
}

func (x GatewayNonEPSConfig_CSFBRat) String() string {
	return proto.EnumName(GatewayNonEPSConfig_CSFBRat_name, int32(x))
}

func (GatewayNonEPSConfig_CSFBRat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{3, 1}
}

type NetworkEPCConfig_NetworkServices int32

const (
	NetworkEPCConfig_METERING    NetworkEPCConfig_NetworkServices = 0 // Deprecated: Do not use.
	NetworkEPCConfig_DPI         NetworkEPCConfig_NetworkServices = 1
	NetworkEPCConfig_ENFORCEMENT NetworkEPCConfig_NetworkServices = 2
)

var NetworkEPCConfig_NetworkServices_name = map[int32]string{
	0: "METERING",
	1: "DPI",
	2: "ENFORCEMENT",
}

var NetworkEPCConfig_NetworkServices_value = map[string]int32{
	"METERING":    0,
	"DPI":         1,
	"ENFORCEMENT": 2,
}

func (x NetworkEPCConfig_NetworkServices) String() string {
	return proto.EnumName(NetworkEPCConfig_NetworkServices_name, int32(x))
}

func (NetworkEPCConfig_NetworkServices) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{6, 0}
}

type CellularGatewayConfig struct {
	Ran                   *GatewayRANConfig    `protobuf:"bytes,1,opt,name=ran,proto3" json:"ran,omitempty"`
	Epc                   *GatewayEPCConfig    `protobuf:"bytes,2,opt,name=epc,proto3" json:"epc,omitempty"`
	NonEpsService         *GatewayNonEPSConfig `protobuf:"bytes,3,opt,name=non_eps_service,json=nonEpsService,proto3" json:"non_eps_service,omitempty"`
	AttachedEnodebSerials []string             `protobuf:"bytes,4,rep,name=attached_enodeb_serials,json=attachedEnodebSerials,proto3" json:"attached_enodeb_serials,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *CellularGatewayConfig) Reset()         { *m = CellularGatewayConfig{} }
func (m *CellularGatewayConfig) String() string { return proto.CompactTextString(m) }
func (*CellularGatewayConfig) ProtoMessage()    {}
func (*CellularGatewayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{0}
}

func (m *CellularGatewayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellularGatewayConfig.Unmarshal(m, b)
}
func (m *CellularGatewayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellularGatewayConfig.Marshal(b, m, deterministic)
}
func (m *CellularGatewayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellularGatewayConfig.Merge(m, src)
}
func (m *CellularGatewayConfig) XXX_Size() int {
	return xxx_messageInfo_CellularGatewayConfig.Size(m)
}
func (m *CellularGatewayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CellularGatewayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CellularGatewayConfig proto.InternalMessageInfo

func (m *CellularGatewayConfig) GetRan() *GatewayRANConfig {
	if m != nil {
		return m.Ran
	}
	return nil
}

func (m *CellularGatewayConfig) GetEpc() *GatewayEPCConfig {
	if m != nil {
		return m.Epc
	}
	return nil
}

func (m *CellularGatewayConfig) GetNonEpsService() *GatewayNonEPSConfig {
	if m != nil {
		return m.NonEpsService
	}
	return nil
}

func (m *CellularGatewayConfig) GetAttachedEnodebSerials() []string {
	if m != nil {
		return m.AttachedEnodebSerials
	}
	return nil
}

type GatewayRANConfig struct {
	// Physical cell ID (0-504)
	Pci int32 `protobuf:"varint,1,opt,name=pci,proto3" json:"pci,omitempty"`
	// Enable eNodeB
	TransmitEnabled      bool     `protobuf:"varint,2,opt,name=transmit_enabled,json=transmitEnabled,proto3" json:"transmit_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayRANConfig) Reset()         { *m = GatewayRANConfig{} }
func (m *GatewayRANConfig) String() string { return proto.CompactTextString(m) }
func (*GatewayRANConfig) ProtoMessage()    {}
func (*GatewayRANConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{1}
}

func (m *GatewayRANConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayRANConfig.Unmarshal(m, b)
}
func (m *GatewayRANConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayRANConfig.Marshal(b, m, deterministic)
}
func (m *GatewayRANConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayRANConfig.Merge(m, src)
}
func (m *GatewayRANConfig) XXX_Size() int {
	return xxx_messageInfo_GatewayRANConfig.Size(m)
}
func (m *GatewayRANConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayRANConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayRANConfig proto.InternalMessageInfo

func (m *GatewayRANConfig) GetPci() int32 {
	if m != nil {
		return m.Pci
	}
	return 0
}

func (m *GatewayRANConfig) GetTransmitEnabled() bool {
	if m != nil {
		return m.TransmitEnabled
	}
	return false
}

type GatewayEPCConfig struct {
	NatEnabled bool `protobuf:"varint,1,opt,name=nat_enabled,json=natEnabled,proto3" json:"nat_enabled,omitempty"`
	// IP block is only set if nat_enabled is false
	// An IP block is a range of IP addresses specified by a network address and
	// a prefix-length of the netmask. For example,
	//    IPv4 IP block:      "192.168.0.0/24"
	//    IPv6 IP block:      "2401:db00:1116:301b::/64"
	IpBlock string `protobuf:"bytes,2,opt,name=ip_block,json=ipBlock,proto3" json:"ip_block,omitempty"`
	// Primary DNS server
	DnsPrimary string `protobuf:"bytes,10,opt,name=dns_primary,json=dnsPrimary,proto3" json:"dns_primary,omitempty"`
	// Secondary DNS server
	DnsSecondary         string   `protobuf:"bytes,11,opt,name=dns_secondary,json=dnsSecondary,proto3" json:"dns_secondary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayEPCConfig) Reset()         { *m = GatewayEPCConfig{} }
func (m *GatewayEPCConfig) String() string { return proto.CompactTextString(m) }
func (*GatewayEPCConfig) ProtoMessage()    {}
func (*GatewayEPCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{2}
}

func (m *GatewayEPCConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayEPCConfig.Unmarshal(m, b)
}
func (m *GatewayEPCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayEPCConfig.Marshal(b, m, deterministic)
}
func (m *GatewayEPCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayEPCConfig.Merge(m, src)
}
func (m *GatewayEPCConfig) XXX_Size() int {
	return xxx_messageInfo_GatewayEPCConfig.Size(m)
}
func (m *GatewayEPCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayEPCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayEPCConfig proto.InternalMessageInfo

func (m *GatewayEPCConfig) GetNatEnabled() bool {
	if m != nil {
		return m.NatEnabled
	}
	return false
}

func (m *GatewayEPCConfig) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

func (m *GatewayEPCConfig) GetDnsPrimary() string {
	if m != nil {
		return m.DnsPrimary
	}
	return ""
}

func (m *GatewayEPCConfig) GetDnsSecondary() string {
	if m != nil {
		return m.DnsSecondary
	}
	return ""
}

type GatewayNonEPSConfig struct {
	// Mobile country code for CSFB
	CsfbMcc string `protobuf:"bytes,1,opt,name=csfb_mcc,json=csfbMcc,proto3" json:"csfb_mcc,omitempty"`
	// Mobile network code for CSFB
	CsfbMnc string `protobuf:"bytes,2,opt,name=csfb_mnc,json=csfbMnc,proto3" json:"csfb_mnc,omitempty"`
	// Location area code. 16-bit
	Lac int32 `protobuf:"varint,3,opt,name=lac,proto3" json:"lac,omitempty"`
	// RAT type (2G/3G), used by eNB
	CsfbRat GatewayNonEPSConfig_CSFBRat `protobuf:"varint,4,opt,name=csfb_rat,json=csfbRat,proto3,enum=magma.cellular.GatewayNonEPSConfig_CSFBRat" json:"csfb_rat,omitempty"`
	// 2G RAT frequencies /ARFCNs for redirection, used by eNB
	Arfcn_2G []int32 `protobuf:"varint,5,rep,packed,name=arfcn_2g,json=arfcn2g,proto3" json:"arfcn_2g,omitempty"`
	// For indicating one of the three modes
	NonEpsServiceControl GatewayNonEPSConfig_NonEPSServiceControl `protobuf:"varint,6,opt,name=non_eps_service_control,json=nonEpsServiceControl,proto3,enum=magma.cellular.GatewayNonEPSConfig_NonEPSServiceControl" json:"non_eps_service_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *GatewayNonEPSConfig) Reset()         { *m = GatewayNonEPSConfig{} }
func (m *GatewayNonEPSConfig) String() string { return proto.CompactTextString(m) }
func (*GatewayNonEPSConfig) ProtoMessage()    {}
func (*GatewayNonEPSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{3}
}

func (m *GatewayNonEPSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayNonEPSConfig.Unmarshal(m, b)
}
func (m *GatewayNonEPSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayNonEPSConfig.Marshal(b, m, deterministic)
}
func (m *GatewayNonEPSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayNonEPSConfig.Merge(m, src)
}
func (m *GatewayNonEPSConfig) XXX_Size() int {
	return xxx_messageInfo_GatewayNonEPSConfig.Size(m)
}
func (m *GatewayNonEPSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayNonEPSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayNonEPSConfig proto.InternalMessageInfo

func (m *GatewayNonEPSConfig) GetCsfbMcc() string {
	if m != nil {
		return m.CsfbMcc
	}
	return ""
}

func (m *GatewayNonEPSConfig) GetCsfbMnc() string {
	if m != nil {
		return m.CsfbMnc
	}
	return ""
}

func (m *GatewayNonEPSConfig) GetLac() int32 {
	if m != nil {
		return m.Lac
	}
	return 0
}

func (m *GatewayNonEPSConfig) GetCsfbRat() GatewayNonEPSConfig_CSFBRat {
	if m != nil {
		return m.CsfbRat
	}
	return GatewayNonEPSConfig_CSFBRAT_2G
}

func (m *GatewayNonEPSConfig) GetArfcn_2G() []int32 {
	if m != nil {
		return m.Arfcn_2G
	}
	return nil
}

func (m *GatewayNonEPSConfig) GetNonEpsServiceControl() GatewayNonEPSConfig_NonEPSServiceControl {
	if m != nil {
		return m.NonEpsServiceControl
	}
	return GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_OFF
}

type CellularNetworkConfig struct {
	Ran                  *NetworkRANConfig `protobuf:"bytes,1,opt,name=ran,proto3" json:"ran,omitempty"`
	Epc                  *NetworkEPCConfig `protobuf:"bytes,2,opt,name=epc,proto3" json:"epc,omitempty"`
	FegNetworkId         string            `protobuf:"bytes,3,opt,name=feg_network_id,json=fegNetworkId,proto3" json:"feg_network_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CellularNetworkConfig) Reset()         { *m = CellularNetworkConfig{} }
func (m *CellularNetworkConfig) String() string { return proto.CompactTextString(m) }
func (*CellularNetworkConfig) ProtoMessage()    {}
func (*CellularNetworkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{4}
}

func (m *CellularNetworkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellularNetworkConfig.Unmarshal(m, b)
}
func (m *CellularNetworkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellularNetworkConfig.Marshal(b, m, deterministic)
}
func (m *CellularNetworkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellularNetworkConfig.Merge(m, src)
}
func (m *CellularNetworkConfig) XXX_Size() int {
	return xxx_messageInfo_CellularNetworkConfig.Size(m)
}
func (m *CellularNetworkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CellularNetworkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CellularNetworkConfig proto.InternalMessageInfo

func (m *CellularNetworkConfig) GetRan() *NetworkRANConfig {
	if m != nil {
		return m.Ran
	}
	return nil
}

func (m *CellularNetworkConfig) GetEpc() *NetworkEPCConfig {
	if m != nil {
		return m.Epc
	}
	return nil
}

func (m *CellularNetworkConfig) GetFegNetworkId() string {
	if m != nil {
		return m.FegNetworkId
	}
	return ""
}

type NetworkRANConfig struct {
	Earfcndl int32 `protobuf:"varint,1,opt,name=earfcndl,proto3" json:"earfcndl,omitempty"`
	// Bandwidth in MHz, from set {1.4, 3, 5, 10, 15, 20}
	BandwidthMhz           int32                       `protobuf:"varint,2,opt,name=bandwidth_mhz,json=bandwidthMhz,proto3" json:"bandwidth_mhz,omitempty"`
	SubframeAssignment     int32                       `protobuf:"varint,3,opt,name=subframe_assignment,json=subframeAssignment,proto3" json:"subframe_assignment,omitempty"`
	SpecialSubframePattern int32                       `protobuf:"varint,5,opt,name=special_subframe_pattern,json=specialSubframePattern,proto3" json:"special_subframe_pattern,omitempty"`
	TddConfig              *NetworkRANConfig_TDDConfig `protobuf:"bytes,6,opt,name=tdd_config,json=tddConfig,proto3" json:"tdd_config,omitempty"`
	FddConfig              *NetworkRANConfig_FDDConfig `protobuf:"bytes,7,opt,name=fdd_config,json=fddConfig,proto3" json:"fdd_config,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                    `json:"-"`
	XXX_unrecognized       []byte                      `json:"-"`
	XXX_sizecache          int32                       `json:"-"`
}

func (m *NetworkRANConfig) Reset()         { *m = NetworkRANConfig{} }
func (m *NetworkRANConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkRANConfig) ProtoMessage()    {}
func (*NetworkRANConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{5}
}

func (m *NetworkRANConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRANConfig.Unmarshal(m, b)
}
func (m *NetworkRANConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRANConfig.Marshal(b, m, deterministic)
}
func (m *NetworkRANConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRANConfig.Merge(m, src)
}
func (m *NetworkRANConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkRANConfig.Size(m)
}
func (m *NetworkRANConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRANConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRANConfig proto.InternalMessageInfo

func (m *NetworkRANConfig) GetEarfcndl() int32 {
	if m != nil {
		return m.Earfcndl
	}
	return 0
}

func (m *NetworkRANConfig) GetBandwidthMhz() int32 {
	if m != nil {
		return m.BandwidthMhz
	}
	return 0
}

func (m *NetworkRANConfig) GetSubframeAssignment() int32 {
	if m != nil {
		return m.SubframeAssignment
	}
	return 0
}

func (m *NetworkRANConfig) GetSpecialSubframePattern() int32 {
	if m != nil {
		return m.SpecialSubframePattern
	}
	return 0
}

func (m *NetworkRANConfig) GetTddConfig() *NetworkRANConfig_TDDConfig {
	if m != nil {
		return m.TddConfig
	}
	return nil
}

func (m *NetworkRANConfig) GetFddConfig() *NetworkRANConfig_FDDConfig {
	if m != nil {
		return m.FddConfig
	}
	return nil
}

type NetworkRANConfig_FDDConfig struct {
	Earfcndl             int32    `protobuf:"varint,1,opt,name=earfcndl,proto3" json:"earfcndl,omitempty"`
	Earfcnul             int32    `protobuf:"varint,2,opt,name=earfcnul,proto3" json:"earfcnul,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkRANConfig_FDDConfig) Reset()         { *m = NetworkRANConfig_FDDConfig{} }
func (m *NetworkRANConfig_FDDConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkRANConfig_FDDConfig) ProtoMessage()    {}
func (*NetworkRANConfig_FDDConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{5, 0}
}

func (m *NetworkRANConfig_FDDConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRANConfig_FDDConfig.Unmarshal(m, b)
}
func (m *NetworkRANConfig_FDDConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRANConfig_FDDConfig.Marshal(b, m, deterministic)
}
func (m *NetworkRANConfig_FDDConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRANConfig_FDDConfig.Merge(m, src)
}
func (m *NetworkRANConfig_FDDConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkRANConfig_FDDConfig.Size(m)
}
func (m *NetworkRANConfig_FDDConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRANConfig_FDDConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRANConfig_FDDConfig proto.InternalMessageInfo

func (m *NetworkRANConfig_FDDConfig) GetEarfcndl() int32 {
	if m != nil {
		return m.Earfcndl
	}
	return 0
}

func (m *NetworkRANConfig_FDDConfig) GetEarfcnul() int32 {
	if m != nil {
		return m.Earfcnul
	}
	return 0
}

type NetworkRANConfig_TDDConfig struct {
	Earfcndl int32 `protobuf:"varint,1,opt,name=earfcndl,proto3" json:"earfcndl,omitempty"`
	// TDD subframe config parameter. See http://niviuk.free.fr/lte_tdd.php
	SubframeAssignment int32 `protobuf:"varint,2,opt,name=subframe_assignment,json=subframeAssignment,proto3" json:"subframe_assignment,omitempty"`
	// TDD subframe config parameter. See http://niviuk.free.fr/lte_tdd.php
	SpecialSubframePattern int32    `protobuf:"varint,3,opt,name=special_subframe_pattern,json=specialSubframePattern,proto3" json:"special_subframe_pattern,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *NetworkRANConfig_TDDConfig) Reset()         { *m = NetworkRANConfig_TDDConfig{} }
func (m *NetworkRANConfig_TDDConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkRANConfig_TDDConfig) ProtoMessage()    {}
func (*NetworkRANConfig_TDDConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{5, 1}
}

func (m *NetworkRANConfig_TDDConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRANConfig_TDDConfig.Unmarshal(m, b)
}
func (m *NetworkRANConfig_TDDConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRANConfig_TDDConfig.Marshal(b, m, deterministic)
}
func (m *NetworkRANConfig_TDDConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRANConfig_TDDConfig.Merge(m, src)
}
func (m *NetworkRANConfig_TDDConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkRANConfig_TDDConfig.Size(m)
}
func (m *NetworkRANConfig_TDDConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRANConfig_TDDConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRANConfig_TDDConfig proto.InternalMessageInfo

func (m *NetworkRANConfig_TDDConfig) GetEarfcndl() int32 {
	if m != nil {
		return m.Earfcndl
	}
	return 0
}

func (m *NetworkRANConfig_TDDConfig) GetSubframeAssignment() int32 {
	if m != nil {
		return m.SubframeAssignment
	}
	return 0
}

func (m *NetworkRANConfig_TDDConfig) GetSpecialSubframePattern() int32 {
	if m != nil {
		return m.SpecialSubframePattern
	}
	return 0
}

type NetworkEPCConfig struct {
	// always 3 digits
	// Mobile country code
	Mcc string `protobuf:"bytes,1,opt,name=mcc,proto3" json:"mcc,omitempty"`
	// 2-3 digits
	// Mobile network code
	Mnc string `protobuf:"bytes,2,opt,name=mnc,proto3" json:"mnc,omitempty"`
	// Tracking area code. 16-bit
	Tac int32 `protobuf:"varint,3,opt,name=tac,proto3" json:"tac,omitempty"`
	// Operator configuration field for LTE
	LteAuthOp []byte `protobuf:"bytes,4,opt,name=lte_auth_op,json=lteAuthOp,proto3" json:"lte_auth_op,omitempty"`
	// Authentication management field for LTE
	LteAuthAmf    []byte                                           `protobuf:"bytes,5,opt,name=lte_auth_amf,json=lteAuthAmf,proto3" json:"lte_auth_amf,omitempty"`
	SubProfiles   map[string]*NetworkEPCConfig_SubscriptionProfile `protobuf:"bytes,6,rep,name=sub_profiles,json=subProfiles,proto3" json:"sub_profiles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DefaultRuleId string                                           `protobuf:"bytes,7,opt,name=default_rule_id,json=defaultRuleId,proto3" json:"default_rule_id,omitempty"`
	// Enable relaying S6a messages via FeG RPC
	RelayEnabled bool `protobuf:"varint,8,opt,name=relay_enabled,json=relayEnabled,proto3" json:"relay_enabled,omitempty"`
	// Ordered list of network service that will be enabled
	NetworkServices []NetworkEPCConfig_NetworkServices `protobuf:"varint,9,rep,packed,name=network_services,json=networkServices,proto3,enum=magma.cellular.NetworkEPCConfig_NetworkServices" json:"network_services,omitempty"`
	// If relay_enabled is false, this determines whether cloud subscriberdb
	// or local subscriberdb is used for authentication requests.
	CloudSubscriberdbEnabled bool     `protobuf:"varint,10,opt,name=cloud_subscriberdb_enabled,json=cloudSubscriberdbEnabled,proto3" json:"cloud_subscriberdb_enabled,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *NetworkEPCConfig) Reset()         { *m = NetworkEPCConfig{} }
func (m *NetworkEPCConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkEPCConfig) ProtoMessage()    {}
func (*NetworkEPCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{6}
}

func (m *NetworkEPCConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkEPCConfig.Unmarshal(m, b)
}
func (m *NetworkEPCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkEPCConfig.Marshal(b, m, deterministic)
}
func (m *NetworkEPCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkEPCConfig.Merge(m, src)
}
func (m *NetworkEPCConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkEPCConfig.Size(m)
}
func (m *NetworkEPCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkEPCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkEPCConfig proto.InternalMessageInfo

func (m *NetworkEPCConfig) GetMcc() string {
	if m != nil {
		return m.Mcc
	}
	return ""
}

func (m *NetworkEPCConfig) GetMnc() string {
	if m != nil {
		return m.Mnc
	}
	return ""
}

func (m *NetworkEPCConfig) GetTac() int32 {
	if m != nil {
		return m.Tac
	}
	return 0
}

func (m *NetworkEPCConfig) GetLteAuthOp() []byte {
	if m != nil {
		return m.LteAuthOp
	}
	return nil
}

func (m *NetworkEPCConfig) GetLteAuthAmf() []byte {
	if m != nil {
		return m.LteAuthAmf
	}
	return nil
}

func (m *NetworkEPCConfig) GetSubProfiles() map[string]*NetworkEPCConfig_SubscriptionProfile {
	if m != nil {
		return m.SubProfiles
	}
	return nil
}

func (m *NetworkEPCConfig) GetDefaultRuleId() string {
	if m != nil {
		return m.DefaultRuleId
	}
	return ""
}

func (m *NetworkEPCConfig) GetRelayEnabled() bool {
	if m != nil {
		return m.RelayEnabled
	}
	return false
}

func (m *NetworkEPCConfig) GetNetworkServices() []NetworkEPCConfig_NetworkServices {
	if m != nil {
		return m.NetworkServices
	}
	return nil
}

func (m *NetworkEPCConfig) GetCloudSubscriberdbEnabled() bool {
	if m != nil {
		return m.CloudSubscriberdbEnabled
	}
	return false
}

type NetworkEPCConfig_SubscriptionProfile struct {
	// Maximum uplink bit rate (AMBR-UL)
	MaxUlBitRate uint64 `protobuf:"varint,1,opt,name=max_ul_bit_rate,json=maxUlBitRate,proto3" json:"max_ul_bit_rate,omitempty"`
	// Maximum downlink bit rate (AMBR-DL)
	MaxDlBitRate         uint64   `protobuf:"varint,2,opt,name=max_dl_bit_rate,json=maxDlBitRate,proto3" json:"max_dl_bit_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkEPCConfig_SubscriptionProfile) Reset()         { *m = NetworkEPCConfig_SubscriptionProfile{} }
func (m *NetworkEPCConfig_SubscriptionProfile) String() string { return proto.CompactTextString(m) }
func (*NetworkEPCConfig_SubscriptionProfile) ProtoMessage()    {}
func (*NetworkEPCConfig_SubscriptionProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{6, 0}
}

func (m *NetworkEPCConfig_SubscriptionProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Unmarshal(m, b)
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Marshal(b, m, deterministic)
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Merge(m, src)
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_Size() int {
	return xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Size(m)
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile proto.InternalMessageInfo

func (m *NetworkEPCConfig_SubscriptionProfile) GetMaxUlBitRate() uint64 {
	if m != nil {
		return m.MaxUlBitRate
	}
	return 0
}

func (m *NetworkEPCConfig_SubscriptionProfile) GetMaxDlBitRate() uint64 {
	if m != nil {
		return m.MaxDlBitRate
	}
	return 0
}

type CellularEnodebConfig struct {
	// EARFCN (0-65535)
	Earfcndl int32 `protobuf:"varint,1,opt,name=earfcndl,proto3" json:"earfcndl,omitempty"`
	// Subframe Assignment (0-6)
	SubframeAssignment int32 `protobuf:"varint,2,opt,name=subframe_assignment,json=subframeAssignment,proto3" json:"subframe_assignment,omitempty"`
	// Special Subframe Pattern (0-9)
	SpecialSubframePattern int32 `protobuf:"varint,3,opt,name=special_subframe_pattern,json=specialSubframePattern,proto3" json:"special_subframe_pattern,omitempty"`
	// Physical cell ID (0-504)
	Pci int32 `protobuf:"varint,4,opt,name=pci,proto3" json:"pci,omitempty"`
	// Enable eNodeB
	TransmitEnabled bool `protobuf:"varint,5,opt,name=transmit_enabled,json=transmitEnabled,proto3" json:"transmit_enabled,omitempty"`
	// eNodeB device class - limited support
	DeviceClass string `protobuf:"bytes,6,opt,name=device_class,json=deviceClass,proto3" json:"device_class,omitempty"`
	// Cellular ID needs to be set to avoid collision
	CellId int32 `protobuf:"varint,7,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	// Bandwidth in MHz
	BandwidthMhz int32 `protobuf:"varint,8,opt,name=bandwidth_mhz,json=bandwidthMhz,proto3" json:"bandwidth_mhz,omitempty"`
	// Cellular tracking area code
	Tac                  int32    `protobuf:"varint,9,opt,name=tac,proto3" json:"tac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CellularEnodebConfig) Reset()         { *m = CellularEnodebConfig{} }
func (m *CellularEnodebConfig) String() string { return proto.CompactTextString(m) }
func (*CellularEnodebConfig) ProtoMessage()    {}
func (*CellularEnodebConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7419489971880e58, []int{7}
}

func (m *CellularEnodebConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellularEnodebConfig.Unmarshal(m, b)
}
func (m *CellularEnodebConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellularEnodebConfig.Marshal(b, m, deterministic)
}
func (m *CellularEnodebConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellularEnodebConfig.Merge(m, src)
}
func (m *CellularEnodebConfig) XXX_Size() int {
	return xxx_messageInfo_CellularEnodebConfig.Size(m)
}
func (m *CellularEnodebConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CellularEnodebConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CellularEnodebConfig proto.InternalMessageInfo

func (m *CellularEnodebConfig) GetEarfcndl() int32 {
	if m != nil {
		return m.Earfcndl
	}
	return 0
}

func (m *CellularEnodebConfig) GetSubframeAssignment() int32 {
	if m != nil {
		return m.SubframeAssignment
	}
	return 0
}

func (m *CellularEnodebConfig) GetSpecialSubframePattern() int32 {
	if m != nil {
		return m.SpecialSubframePattern
	}
	return 0
}

func (m *CellularEnodebConfig) GetPci() int32 {
	if m != nil {
		return m.Pci
	}
	return 0
}

func (m *CellularEnodebConfig) GetTransmitEnabled() bool {
	if m != nil {
		return m.TransmitEnabled
	}
	return false
}

func (m *CellularEnodebConfig) GetDeviceClass() string {
	if m != nil {
		return m.DeviceClass
	}
	return ""
}

func (m *CellularEnodebConfig) GetCellId() int32 {
	if m != nil {
		return m.CellId
	}
	return 0
}

func (m *CellularEnodebConfig) GetBandwidthMhz() int32 {
	if m != nil {
		return m.BandwidthMhz
	}
	return 0
}

func (m *CellularEnodebConfig) GetTac() int32 {
	if m != nil {
		return m.Tac
	}
	return 0
}

func init() {
	proto.RegisterEnum("magma.cellular.GatewayNonEPSConfig_NonEPSServiceControl", GatewayNonEPSConfig_NonEPSServiceControl_name, GatewayNonEPSConfig_NonEPSServiceControl_value)
	proto.RegisterEnum("magma.cellular.GatewayNonEPSConfig_CSFBRat", GatewayNonEPSConfig_CSFBRat_name, GatewayNonEPSConfig_CSFBRat_value)
	proto.RegisterEnum("magma.cellular.NetworkEPCConfig_NetworkServices", NetworkEPCConfig_NetworkServices_name, NetworkEPCConfig_NetworkServices_value)
	proto.RegisterType((*CellularGatewayConfig)(nil), "magma.cellular.CellularGatewayConfig")
	proto.RegisterType((*GatewayRANConfig)(nil), "magma.cellular.GatewayRANConfig")
	proto.RegisterType((*GatewayEPCConfig)(nil), "magma.cellular.GatewayEPCConfig")
	proto.RegisterType((*GatewayNonEPSConfig)(nil), "magma.cellular.GatewayNonEPSConfig")
	proto.RegisterType((*CellularNetworkConfig)(nil), "magma.cellular.CellularNetworkConfig")
	proto.RegisterType((*NetworkRANConfig)(nil), "magma.cellular.NetworkRANConfig")
	proto.RegisterType((*NetworkRANConfig_FDDConfig)(nil), "magma.cellular.NetworkRANConfig.FDDConfig")
	proto.RegisterType((*NetworkRANConfig_TDDConfig)(nil), "magma.cellular.NetworkRANConfig.TDDConfig")
	proto.RegisterType((*NetworkEPCConfig)(nil), "magma.cellular.NetworkEPCConfig")
	proto.RegisterMapType((map[string]*NetworkEPCConfig_SubscriptionProfile)(nil), "magma.cellular.NetworkEPCConfig.SubProfilesEntry")
	proto.RegisterType((*NetworkEPCConfig_SubscriptionProfile)(nil), "magma.cellular.NetworkEPCConfig.SubscriptionProfile")
	proto.RegisterType((*CellularEnodebConfig)(nil), "magma.cellular.CellularEnodebConfig")
}

func init() { proto.RegisterFile("cellular_service.proto", fileDescriptor_7419489971880e58) }

var fileDescriptor_7419489971880e58 = []byte{
	// 1150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0xbf, 0x24, 0x4d, 0x93, 0x4c, 0x72, 0x8d, 0xb5, 0xf7, 0xcf, 0x04, 0x89, 0x0b, 0xb9, 0x03,
	0xf5, 0x40, 0x0a, 0x90, 0x43, 0xe8, 0x84, 0xe0, 0xa1, 0x4d, 0xdd, 0x2a, 0x40, 0x93, 0x68, 0x1d,
	0x78, 0x80, 0x87, 0xd5, 0xda, 0x5e, 0xa7, 0x56, 0xed, 0xb5, 0x65, 0xaf, 0xef, 0xae, 0xf7, 0xc0,
	0x13, 0x9f, 0x00, 0xf1, 0x15, 0x78, 0xe5, 0x9b, 0xf0, 0x0d, 0x78, 0xe2, 0x93, 0xa0, 0x5d, 0xaf,
	0x9d, 0x6b, 0xd4, 0xeb, 0x15, 0x9e, 0x78, 0xb2, 0x67, 0xe6, 0xf7, 0x9b, 0x9d, 0x9d, 0x9d, 0x99,
	0x5d, 0xb8, 0xef, 0xb2, 0x30, 0xcc, 0x43, 0x9a, 0x92, 0x8c, 0xa5, 0xcf, 0x03, 0x97, 0x8d, 0x93,
	0x34, 0x16, 0x31, 0xda, 0x8b, 0xe8, 0x3a, 0xa2, 0xe3, 0xd2, 0x3a, 0xfa, 0xa5, 0x0e, 0xf7, 0xa6,
	0x5a, 0x38, 0xa1, 0x82, 0xbd, 0xa0, 0x17, 0xd3, 0x98, 0xfb, 0xc1, 0x1a, 0x4d, 0xa0, 0x91, 0x52,
	0x6e, 0xd6, 0x86, 0xb5, 0xfd, 0xee, 0x64, 0x38, 0xbe, 0xcc, 0x1b, 0x6b, 0x2c, 0x3e, 0x98, 0x17,
	0x70, 0x2c, 0xc1, 0x92, 0xc3, 0x12, 0xd7, 0xac, 0x5f, 0xcb, 0xb1, 0x96, 0xd3, 0x92, 0xc3, 0x12,
	0x17, 0x7d, 0x0b, 0x7d, 0x1e, 0x73, 0xc2, 0x92, 0xac, 0x0c, 0xd5, 0x6c, 0x28, 0xfe, 0xa3, 0x37,
	0xf0, 0xe7, 0x31, 0xb7, 0x96, 0xb6, 0x76, 0x71, 0x9b, 0xc7, 0xdc, 0x4a, 0x32, 0xbb, 0x60, 0xa2,
	0x2f, 0xe0, 0x01, 0x15, 0x82, 0xba, 0x67, 0xcc, 0x23, 0x8c, 0xc7, 0x1e, 0x73, 0xa4, 0xd3, 0x80,
	0x86, 0x99, 0xb9, 0x33, 0x6c, 0xec, 0x77, 0xf0, 0xbd, 0xd2, 0x6c, 0x29, 0xab, 0x5d, 0x18, 0x47,
	0x0b, 0x30, 0xb6, 0x77, 0x84, 0x0c, 0x68, 0x24, 0x6e, 0xa0, 0x12, 0xd0, 0xc4, 0xf2, 0x17, 0x3d,
	0x01, 0x43, 0xa4, 0x94, 0x67, 0x51, 0x20, 0x08, 0xe3, 0xd4, 0x09, 0x99, 0xa7, 0xf6, 0xda, 0xc6,
	0xfd, 0x52, 0x6f, 0x15, 0xea, 0xd1, 0x6f, 0xb5, 0xca, 0x63, 0xb5, 0x5f, 0xf4, 0x10, 0xba, 0x9c,
	0x6e, 0xa8, 0x35, 0x45, 0x05, 0x4e, 0x4b, 0x16, 0x7a, 0x07, 0xda, 0x41, 0x42, 0x9c, 0x30, 0x76,
	0xcf, 0x95, 0xe3, 0x0e, 0x6e, 0x05, 0xc9, 0xa1, 0x14, 0x25, 0xd7, 0xe3, 0x19, 0x49, 0xd2, 0x20,
	0xa2, 0xe9, 0x85, 0x09, 0xca, 0x0a, 0x1e, 0xcf, 0x96, 0x85, 0x06, 0x3d, 0x82, 0xdb, 0x12, 0x90,
	0x31, 0x37, 0xe6, 0x9e, 0x84, 0x74, 0x15, 0xa4, 0xe7, 0xf1, 0xcc, 0x2e, 0x75, 0xa3, 0xbf, 0x1a,
	0x70, 0xe7, 0x8a, 0x34, 0xca, 0x85, 0xdd, 0xcc, 0x77, 0x48, 0xe4, 0xba, 0x2a, 0xac, 0x0e, 0x6e,
	0x49, 0xf9, 0xd4, 0x75, 0x37, 0x26, 0xee, 0x96, 0x31, 0x29, 0x13, 0x77, 0x65, 0x86, 0x42, 0xea,
	0xaa, 0xe3, 0x6a, 0x62, 0xf9, 0x8b, 0x8e, 0x35, 0x38, 0xa5, 0xc2, 0xdc, 0x19, 0xd6, 0xf6, 0xf7,
	0x26, 0x1f, 0xdf, 0xe0, 0x14, 0xc7, 0x53, 0xfb, 0xf8, 0x10, 0x53, 0x51, 0x78, 0xc6, 0x54, 0xc8,
	0x45, 0x69, 0xea, 0xbb, 0x9c, 0x4c, 0xd6, 0x66, 0x73, 0xd8, 0xd8, 0x6f, 0xe2, 0x96, 0x92, 0x27,
	0x6b, 0x14, 0xc3, 0x83, 0xad, 0x7a, 0x21, 0x6e, 0xcc, 0x45, 0x1a, 0x87, 0xe6, 0xae, 0x5a, 0xf1,
	0xd9, 0x4d, 0x56, 0x2c, 0x04, 0x5d, 0x36, 0xd3, 0x82, 0x8f, 0xef, 0x5e, 0x2a, 0x26, 0xad, 0x1d,
	0xfd, 0x0c, 0x77, 0xaf, 0x42, 0xa3, 0x87, 0xf0, 0xee, 0x7c, 0x31, 0x27, 0xd6, 0xd2, 0x26, 0xb6,
	0x85, 0x7f, 0x98, 0x4d, 0x2d, 0x32, 0x5d, 0xcc, 0x57, 0x78, 0xf1, 0x1d, 0x59, 0x1c, 0x1f, 0x1b,
	0xb7, 0xd0, 0x63, 0x18, 0xbe, 0x09, 0x20, 0x37, 0x4c, 0xec, 0x53, 0xdb, 0xa8, 0x5d, 0xe7, 0x46,
	0x02, 0xea, 0xa3, 0x27, 0xd0, 0xd2, 0xf9, 0x41, 0x7b, 0x00, 0xea, 0xf7, 0x60, 0x45, 0x26, 0x27,
	0xc6, 0xad, 0xd7, 0xe5, 0xa7, 0x27, 0x46, 0x6d, 0xf4, 0x7b, 0x6d, 0xd3, 0xcd, 0x73, 0x26, 0x5e,
	0xc4, 0xe9, 0xf9, 0x8d, 0xba, 0x59, 0x63, 0xff, 0x55, 0x37, 0x6b, 0xce, 0x56, 0x37, 0x3f, 0x86,
	0x3d, 0x9f, 0xad, 0x09, 0x2f, 0x8c, 0x24, 0xf0, 0x54, 0x75, 0x74, 0x70, 0xcf, 0x67, 0x6b, 0xcd,
	0x98, 0x79, 0xa3, 0x3f, 0x76, 0xc0, 0xd8, 0x5e, 0x13, 0x0d, 0xa0, 0xcd, 0xd4, 0x21, 0x7b, 0xa1,
	0x6e, 0xba, 0x4a, 0x96, 0xc5, 0xed, 0x50, 0xee, 0xbd, 0x08, 0x3c, 0x71, 0x46, 0xa2, 0xb3, 0x57,
	0x2a, 0xa8, 0x26, 0xee, 0x55, 0xca, 0xd3, 0xb3, 0x57, 0xe8, 0x13, 0xb8, 0x93, 0xe5, 0x8e, 0x9f,
	0xd2, 0x88, 0x11, 0x9a, 0x65, 0xc1, 0x9a, 0x47, 0x8c, 0x0b, 0x5d, 0x9e, 0xa8, 0x34, 0x1d, 0x54,
	0x16, 0xf4, 0x0c, 0xcc, 0x2c, 0x61, 0x6e, 0x40, 0x43, 0x52, 0x11, 0x13, 0x2a, 0x04, 0x4b, 0xb9,
	0xd9, 0x54, 0xac, 0xfb, 0xda, 0x6e, 0x6b, 0xf3, 0xb2, 0xb0, 0xa2, 0x19, 0x80, 0xf0, 0x3c, 0x59,
	0x78, 0x7e, 0xb0, 0x56, 0x75, 0xd7, 0x9d, 0x7c, 0xf4, 0xb6, 0xac, 0x8e, 0x57, 0x47, 0x47, 0x3a,
	0x57, 0x1d, 0xe1, 0x79, 0x7a, 0xdb, 0x33, 0x00, 0x7f, 0xe3, 0xaa, 0x75, 0x43, 0x57, 0xc7, 0x1b,
	0x57, 0x7e, 0xe9, 0x6a, 0x30, 0x85, 0x4e, 0xa5, 0xbf, 0x36, 0x9d, 0x95, 0x2d, 0x0f, 0x75, 0x26,
	0x2b, 0x79, 0xf0, 0x6b, 0x0d, 0x3a, 0xab, 0x1b, 0x79, 0x79, 0x43, 0xbe, 0xeb, 0xff, 0x29, 0xdf,
	0x8d, 0xeb, 0xf2, 0x3d, 0xfa, 0xbb, 0x59, 0x15, 0xcc, 0x66, 0x9c, 0x1a, 0xd0, 0xd8, 0xcc, 0x2b,
	0xf9, 0xab, 0x34, 0xd5, 0x98, 0x92, 0xbf, 0x52, 0x23, 0x36, 0x23, 0x4a, 0x50, 0x17, 0xbd, 0x07,
	0xdd, 0x50, 0x30, 0x42, 0x73, 0x71, 0x46, 0xe2, 0x44, 0x4d, 0xa9, 0x1e, 0xee, 0x84, 0x82, 0x1d,
	0xe4, 0xe2, 0x6c, 0x91, 0xa0, 0x21, 0xf4, 0x2a, 0x3b, 0x8d, 0x7c, 0x55, 0x08, 0x3d, 0x0c, 0x1a,
	0x70, 0x10, 0xf9, 0x68, 0x05, 0xbd, 0x2c, 0x77, 0x48, 0x92, 0xc6, 0x7e, 0x10, 0xb2, 0xcc, 0xdc,
	0x1d, 0x36, 0xf6, 0xbb, 0x93, 0xcf, 0xde, 0xd6, 0x20, 0x63, 0x3b, 0x77, 0x96, 0x9a, 0x63, 0x71,
	0x91, 0x5e, 0xe0, 0x6e, 0xb6, 0xd1, 0xa0, 0x0f, 0xa1, 0xef, 0x31, 0x9f, 0xe6, 0xa1, 0x20, 0x69,
	0x1e, 0x32, 0xd9, 0x3a, 0x2d, 0xb5, 0x8f, 0xdb, 0x5a, 0x8d, 0xf3, 0x90, 0xcd, 0x3c, 0xd9, 0x0a,
	0x29, 0x0b, 0xe9, 0x45, 0x75, 0x8d, 0xb4, 0xd5, 0x35, 0xd2, 0x53, 0xca, 0xf2, 0x22, 0xf9, 0x09,
	0x8c, 0xb2, 0x05, 0xf5, 0x90, 0xcc, 0xcc, 0xce, 0xb0, 0xb1, 0xbf, 0x37, 0xf9, 0xf4, 0xad, 0x61,
	0x6a, 0x85, 0x9e, 0x76, 0x19, 0xee, 0xf3, 0xcb, 0x0a, 0xf4, 0x15, 0x0c, 0xdc, 0x30, 0xce, 0x3d,
	0x79, 0x88, 0x99, 0x9b, 0x06, 0x0e, 0x4b, 0x3d, 0xa7, 0x0a, 0x07, 0x54, 0x38, 0xa6, 0x42, 0xd8,
	0xaf, 0x01, 0x74, 0x68, 0x03, 0x17, 0xee, 0x68, 0x75, 0x22, 0x82, 0x98, 0xeb, 0xfd, 0xa3, 0x0f,
	0xa0, 0x1f, 0xd1, 0x97, 0x24, 0x0f, 0x89, 0x13, 0x08, 0x79, 0x7f, 0x30, 0x75, 0xb0, 0x3b, 0xb8,
	0x17, 0xd1, 0x97, 0xdf, 0x87, 0x87, 0x81, 0xc0, 0x54, 0x54, 0x30, 0xef, 0x35, 0x58, 0xbd, 0x82,
	0x1d, 0x95, 0xb0, 0x81, 0x00, 0x63, 0x3b, 0xdb, 0xb2, 0x14, 0xce, 0xd9, 0x45, 0x59, 0x2e, 0xe7,
	0xec, 0x02, 0x7d, 0x03, 0xcd, 0xe7, 0x34, 0xcc, 0x99, 0x1e, 0x71, 0x9f, 0xdf, 0xe4, 0x04, 0xb7,
	0x03, 0xc7, 0x85, 0x8b, 0x2f, 0xeb, 0xcf, 0x6a, 0xa3, 0xaf, 0xa1, 0xbf, 0x95, 0x3c, 0x64, 0x40,
	0xfb, 0xd4, 0x5a, 0x59, 0x78, 0x36, 0x3f, 0x31, 0x6e, 0x0d, 0xea, 0xed, 0x1a, 0x6a, 0x41, 0xe3,
	0x68, 0x39, 0x33, 0x6a, 0xa8, 0x0f, 0x5d, 0x6b, 0x7e, 0xbc, 0xc0, 0x53, 0xeb, 0xd4, 0x9a, 0xaf,
	0x8c, 0xfa, 0xe8, 0xcf, 0x3a, 0xdc, 0x2d, 0xa7, 0x77, 0xf1, 0x3c, 0xf9, 0x5f, 0x35, 0x61, 0xf9,
	0x20, 0xda, 0xb9, 0xfe, 0x41, 0xd4, 0xbc, 0xf2, 0x41, 0x84, 0xde, 0x87, 0x9e, 0xc7, 0x8a, 0xdb,
	0x3a, 0xa4, 0x59, 0xa6, 0x66, 0x66, 0x07, 0x77, 0x0b, 0xdd, 0x54, 0xaa, 0xd0, 0x03, 0x68, 0xc9,
	0xd4, 0x97, 0x95, 0xdf, 0xc4, 0xbb, 0x52, 0x2c, 0x4a, 0xfe, 0xf2, 0xf4, 0x6f, 0x5f, 0x31, 0xfd,
	0x75, 0xa7, 0x77, 0xaa, 0x4e, 0x3f, 0x6c, 0xff, 0xb8, 0xab, 0x1e, 0xbd, 0x99, 0x53, 0x7c, 0x9f,
	0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x65, 0xf1, 0x0f, 0xa0, 0x16, 0x0b, 0x00, 0x00,
}

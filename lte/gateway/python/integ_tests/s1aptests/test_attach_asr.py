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

import unittest
import s1ap_types
import time

from integ_tests.s1aptests import s1ap_wrapper
from integ_tests.s1aptests.s1ap_utils import SpgwUtil
from integ_tests.s1aptests.s1ap_utils import SessionManagerUtil
from integ_tests.s1aptests.ovs.rest_api import get_datapath, get_flows


class TestAttachASR(unittest.TestCase):
    SPGW_TABLE = 0
    LOCAL_PORT = "LOCAL"

    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()
        self._sessionManager_util = SessionManagerUtil()
        self._spgw_util = SpgwUtil()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_attach_asr_tcp_data(self):
        """ attach + send ASR Req to session manager with a"""
        """ single UE """
        num_ues = 1
        self._s1ap_wrapper.configUEDevice(num_ues)
        datapath = get_datapath()

        req = self._s1ap_wrapper.ue_req
        print(
            "********************** Running End to End attach for ",
            "UE id ",
            req.ue_id,
        )
        # Now actually complete the attach
        self._s1ap_wrapper._s1_util.attach(
            req.ue_id,
            s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
            s1ap_types.ueAttachAccept_t,
        )

        # Wait on EMM Information from MME
        self._s1ap_wrapper._s1_util.receive_emm_info()

        print("Sleeping for 5 seconds")
        time.sleep(5)

        print(
            "********************** Sending ASR for IMSI",
            "".join([str(i) for i in req.imsi]),
        )
        self._sessionManager_util.create_AbortSessionRequest(
            "IMSI" + "".join([str(i) for i in req.imsi]),
        )

        print("Sleeping for 5 seconds")
        time.sleep(5)

        print("Checking that uplink/downlink flows were deleted")
        flows = get_flows(
            datapath, {"table_id": self.SPGW_TABLE, "priority": 0}
        )
        self.assertEqual(
            len(flows), 2, "There should only be 2 default table 0 flows"
        )


if __name__ == "__main__":
    unittest.main()

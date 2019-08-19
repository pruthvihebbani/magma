/**
 * Copyright (c) 2016-present, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree. An additional grant
 * of patent rights can be found in the PATENTS file in the same directory.
 */

#include <memory>

#include <glog/logging.h>
#include <gmock/gmock.h>
#include <grpc++/grpc++.h>
#include <gtest/gtest.h>

#include <lte/protos/sctpd.grpc.pb.h>

#include "sctpd_event_handler.h"
#include "sctpd_uplink_client.h"

using ::testing::_;
using ::testing::AllOf;
using ::testing::Eq;
using ::testing::NotNull;
using ::testing::Property;
using ::testing::Return;
using ::testing::Test;

namespace magma {
namespace sctpd {

class MockSctpdUplinkClient final : public SctpdUplinkClient {
 public:
  MockSctpdUplinkClient(std::shared_ptr<Channel> channel):
    SctpdUplinkClient(channel)
  {
    ON_CALL(*this, sendUl(_, _)).WillByDefault(Return(0));
    ON_CALL(*this, newAssoc(_, _)).WillByDefault(Return(0));
    ON_CALL(*this, closeAssoc(_, _)).WillByDefault(Return(0));
  }

  MOCK_METHOD2(sendUl, int(const SendUlReq &, SendUlRes *));
  MOCK_METHOD2(newAssoc, int(const NewAssocReq &, NewAssocRes *));
  MOCK_METHOD2(closeAssoc, int(const CloseAssocReq &, CloseAssocRes *));
};

class EventHandlerTest : public ::testing::Test {
 protected:
  virtual void SetUp()
  {
    send_ul_req.set_assoc_id(123);
    send_ul_req.set_stream(321);
    send_ul_req.set_payload("testtest");

    new_assoc_req.set_assoc_id(1234);
    new_assoc_req.set_instreams(16);
    new_assoc_req.set_outstreams(32);

    close_assoc_req.set_assoc_id(12345);
    close_assoc_req.set_is_reset(true);

    auto channel = grpc::CreateChannel("", grpc::InsecureChannelCredentials());

    _uplink_client = std::make_unique<MockSctpdUplinkClient>(channel);
    _handler = std::make_unique<SctpdEventHandler>(*_uplink_client);
  }

  SendUlReq send_ul_req;
  NewAssocReq new_assoc_req;
  CloseAssocReq close_assoc_req;

  std::unique_ptr<MockSctpdUplinkClient> _uplink_client;
  std::unique_ptr<SctpdEventHandler> _handler;
};

TEST_F(EventHandlerTest, test_event_handler_new_assoc)
{
  auto correct_assoc_id =
    Property(&NewAssocReq::assoc_id, Eq(new_assoc_req.assoc_id()));
  auto correct_instreams =
    Property(&NewAssocReq::instreams, Eq(new_assoc_req.instreams()));
  auto correct_outstreams =
    Property(&NewAssocReq::outstreams, Eq(new_assoc_req.outstreams()));
  auto correct_new_assoc_req =
    AllOf(correct_assoc_id, correct_instreams, correct_outstreams);

  EXPECT_CALL(*_uplink_client, newAssoc(correct_new_assoc_req, NotNull()))
    .Times(1);

  _handler->HandleNewAssoc(
    new_assoc_req.assoc_id(),
    new_assoc_req.instreams(),
    new_assoc_req.outstreams());
}

TEST_F(EventHandlerTest, test_event_handler_close_assoc)
{
  auto correct_assoc_id =
    Property(&CloseAssocReq::assoc_id, Eq(close_assoc_req.assoc_id()));
  auto correct_reset =
    Property(&CloseAssocReq::is_reset, Eq(close_assoc_req.is_reset()));
  auto correct_close_assoc_req = AllOf(correct_assoc_id, correct_reset);

  EXPECT_CALL(*_uplink_client, closeAssoc(correct_close_assoc_req, NotNull()))
    .Times(1);

  _handler->HandleCloseAssoc(
    close_assoc_req.assoc_id(), close_assoc_req.is_reset());
}

TEST_F(EventHandlerTest, test_event_handler_send_ul)
{
  auto correct_assoc_id =
    Property(&SendUlReq::assoc_id, Eq(send_ul_req.assoc_id()));
  auto correct_stream = Property(&SendUlReq::stream, Eq(send_ul_req.stream()));
  auto correct_payload =
    Property(&SendUlReq::payload, Eq(send_ul_req.payload()));
  auto correct_send_ul_req =
    AllOf(correct_assoc_id, correct_stream, correct_payload);

  EXPECT_CALL(*_uplink_client, sendUl(correct_send_ul_req, NotNull())).Times(1);

  _handler->HandleRecv(
    send_ul_req.assoc_id(), send_ul_req.stream(), send_ul_req.payload());
}

} // namespace sctpd
} // namespace magma

int main(int argc, char **argv)
{
  ::testing::InitGoogleTest(&argc, argv);
  FLAGS_logtostderr = 1;
  FLAGS_v = 10;
  return RUN_ALL_TESTS();
}

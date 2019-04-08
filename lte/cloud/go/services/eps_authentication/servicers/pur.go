/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers

import (
	"errors"

	"magma/feg/cloud/go/protos"
	"magma/lte/cloud/go/services/eps_authentication/metrics"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (srv *EPSAuthServer) PurgeUE(ctx context.Context, purge *protos.PurgeUERequest) (*protos.PurgeUEAnswer, error) {
	glog.V(2).Infof("received PUR from: %s", purge.GetUserName())
	metrics.PURequests.Inc()
	if err := validatePUR(purge); err != nil {
		glog.V(2).Infof("PUR is invalid: %v", err.Error())
		metrics.InvalidRequests.Inc()
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	networkID, err := getNetworkID(ctx)
	if err != nil {
		glog.V(2).Infof("could not lookup networkID: %v", err.Error())
		metrics.NetworkIDErrors.Inc()
		return nil, err
	}
	_, errorCode, err := srv.lookupSubscriber(purge.UserName, networkID)
	if err != nil {
		glog.V(2).Infof("failed to lookup subscriber '%s': %v", purge.UserName, err.Error())
		metrics.UnknownSubscribers.Inc()
		return &protos.PurgeUEAnswer{ErrorCode: errorCode}, err
	}
	return &protos.PurgeUEAnswer{ErrorCode: protos.ErrorCode_SUCCESS}, nil
}

// validatePUR returns an error iff the PUR is invalid.
func validatePUR(purge *protos.PurgeUERequest) error {
	if purge == nil {
		return errors.New("received a nil PurgeUERequest")
	}
	if len(purge.UserName) == 0 {
		return errors.New("user name was empty")
	}
	return nil
}

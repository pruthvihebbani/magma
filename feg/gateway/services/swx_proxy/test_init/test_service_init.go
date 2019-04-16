/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package test_init

import (
	"fmt"
	"testing"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/mconfig"
	"magma/feg/gateway/registry"
	"magma/feg/gateway/services/swx_proxy/cache"
	"magma/feg/gateway/services/swx_proxy/servicers"
	"magma/feg/gateway/services/swx_proxy/servicers/test"
	"magma/orc8r/cloud/go/service"
	"magma/orc8r/cloud/go/test_utils"
)

func StartTestService(t *testing.T) (*service.Service, error) {
	return StartTestServiceWithCache(t, cache.New())
}

func StartTestServiceWithCache(t *testing.T, cache *cache.Impl) (*service.Service, error) {
	srv, lis := test_utils.NewTestService(t, registry.ModuleName, registry.SWX_PROXY)

	config := servicers.GetSwxProxyConfig()
	serverAddr, err := test.StartTestSwxServer(config.ServerCfg.Protocol, config.ServerCfg.Addr)
	if err != nil {
		return nil, err
	}
	// Update server config with chosen port of swx test server
	config.ServerCfg.Addr = serverAddr
	service, err := servicers.NewSwxProxyWithCache(config, cache)
	if err != nil {
		return nil, err
	}
	protos.RegisterSwxProxyServer(srv.GrpcServer, service)
	go srv.RunTest(lis)
	return srv, nil
}

func InitTestMconfig(t *testing.T, addr string, verify bool) error {
	// Create tmp mconfig test file & load configs from it
	fegConfigFmt := `{
		"configsByKey": {
			"swx_proxy": {
				"@type": "type.googleapis.com/magma.mconfig.SwxConfig",
				"logLevel": "INFO",
				"server": {
					"protocol": "sctp",
					"address": "%s",
					"retransmits": 3,
					"watchdogInterval": 1,
					"retryCount": 5,
					"productName": "magma_test",
					"realm": "openair4G.eur",
					"host": "magma-oai.openair4G.eur"
				},
				"verifyAuthorization": %t
			}
		}
	}`
	return mconfig.CreateLoadTempConfig(fmt.Sprintf(fegConfigFmt, addr, verify))
}

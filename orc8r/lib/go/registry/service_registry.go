/*
 Copyright (c) Facebook, Inc. and its affiliates.
 All rights reserved.

 This source code is licensed under the BSD-style license found in the
 LICENSE file in the root directory of this source tree.
*/

package registry

import (
	"fmt"
	"strings"

	"magma/orc8r/lib/go/service/config"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

const (
	serviceRegistryFilename = "service_registry"
)

type rawMapType = map[interface{}]interface{}

// LoadServiceRegistryConfig reads service registry config file from
// /etc/magma/configs/{moduleName} or override config.
func LoadServiceRegistryConfig(moduleName string) ([]ServiceLocation, error) {
	cfg, err := config.GetServiceConfig(moduleName, serviceRegistryFilename)
	if err != nil {
		// file does not exist
		return nil, err
	}
	rawMap, err := getRawMap(cfg)
	if err != nil {
		// file is empty
		return nil, err
	}
	locations, err := convertToServiceLocations(rawMap)
	if err != nil {
		glog.Errorf("Failed to load in service registry for %s:%s.yml: %v", moduleName, serviceRegistryFilename, err)
		return nil, err
	}

	return locations, nil
}

// LoadServiceRegistryConfigs reads service registry config file from all
// modules under /etc/magma/configs/.
// Also, prefers override configs on a per-module basis.
func LoadServiceRegistryConfigs() ([]ServiceLocation, error) {
	configsByModule, err := config.GetServiceConfigs(serviceRegistryFilename)
	if err != nil {
		return nil, err
	}

	var ret []ServiceLocation

	for module, cfg := range configsByModule {
		rawMap, err := getRawMap(cfg)
		if err != nil {
			glog.Warningf("Ignoring malformed service registry %s:%s.yml", module, serviceRegistryFilename)
			continue
		}
		locations, err := convertToServiceLocations(rawMap)
		if err != nil {
			return nil, errors.Wrapf(err, "load service registry for %s:%s.yml", module, serviceRegistryFilename)
		}
		ret = append(ret, locations...)
	}

	return ret, nil
}

func getRawMap(serviceRegistry *config.ConfigMap) (map[interface{}]interface{}, error) {
	services, ok := serviceRegistry.RawMap["services"]
	if !ok {
		return nil, fmt.Errorf("the field services does not exist")
	}
	rawMap, ok := services.(rawMapType)
	if !ok {
		return nil, fmt.Errorf("unable to convert service registry to map")
	}
	return rawMap, nil
}

func convertToServiceLocations(rawMap rawMapType) ([]ServiceLocation, error) {
	serviceLocations := make([]ServiceLocation, 0, len(rawMap))
	for k, v := range rawMap {
		name, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("the name of the service is not a string: %v", k)
		}
		rawMap, ok := v.(rawMapType)
		if !ok {
			return nil, fmt.Errorf("the value associated with key:%v is not a map: %v", k, v)
		}
		configMap := &config.ConfigMap{RawMap: rawMap}
		host, err := configMap.GetString("host")
		if err != nil {
			// Check old/py format: 'ip_address'
			var ipErr error
			if host, ipErr = configMap.GetString("ip_address"); ipErr != nil {
				return nil, err
			}
		}
		port, err := configMap.GetInt("port")
		if err != nil {
			return nil, err
		}
		// echoPort is an optional field used for services which run an echo
		// server
		echoPort, err := configMap.GetInt("echo_port")
		if err != nil {
			echoPort = 0
		}
		proxyAliases := getProxyAliases(rawMap)
		serviceLocations = append(serviceLocations, ServiceLocation{Name: strings.ToUpper(name), Host: host, Port: port, EchoPort: echoPort, ProxyAliases: proxyAliases})
	}
	return serviceLocations, nil
}

func getProxyAliases(rawMap map[interface{}]interface{}) map[string]int {
	proxyAliases := map[string]int{}
	if val, ok := rawMap["proxy_aliases"]; ok {
		rawMap, _ := val.(rawMapType)
		for k, v := range rawMap {
			proxyName, _ := k.(string)
			portMap, _ := v.(rawMapType)
			port, _ := portMap["port"].(int)
			proxyAliases[proxyName] = port
		}
	}
	return proxyAliases
}

/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers

import (
	"net/http"
	"reflect"
	"sort"

	"magma/orc8r/cloud/go/models"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/services/configurator"
	"magma/orc8r/cloud/go/services/device"
	"magma/orc8r/cloud/go/services/orchestrator/obsidian/handlers"
	orc8rmodels "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"magma/orc8r/cloud/go/storage"
	merrors "magma/orc8r/lib/go/errors"
	wifimodels "magma/wifi/cloud/go/services/wifi/obsidian/models"
	"magma/wifi/cloud/go/wifi"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const (
	WifiNetworks                 = "wifi"
	BaseNetworksPath             = obsidian.V1Root + WifiNetworks
	ManageNetworkPath            = BaseNetworksPath + obsidian.UrlSep + ":network_id"
	ManageNetworkNamePath        = ManageNetworkPath + obsidian.UrlSep + "name"
	ManageNetworkDescriptionPath = ManageNetworkPath + obsidian.UrlSep + "description"
	ManageNetworkFeaturesPath    = ManageNetworkPath + obsidian.UrlSep + "features"
	ManageNetworkWifiPath        = ManageNetworkPath + obsidian.UrlSep + "wifi"

	BaseGatewaysPath             = ManageNetworkPath + obsidian.UrlSep + "gateways"
	ManageGatewayPath            = BaseGatewaysPath + obsidian.UrlSep + ":gateway_id"
	ManageGatewayNamePath        = ManageGatewayPath + obsidian.UrlSep + "name"
	ManageGatewayDescriptionPath = ManageGatewayPath + obsidian.UrlSep + "description"
	ManageGatewayConfigPath      = ManageGatewayPath + obsidian.UrlSep + "magmad"
	ManageGatewayDevicePath      = ManageGatewayPath + obsidian.UrlSep + "device"
	ManageGatewayStatePath       = ManageGatewayPath + obsidian.UrlSep + "status"
	ManageGatewayTierPath        = ManageGatewayPath + obsidian.UrlSep + "tier"
	ManageGatewayWifiPath        = ManageGatewayPath + obsidian.UrlSep + "wifi"

	MeshID               = "mesh_id"
	BaseMeshesPath       = ManageNetworkPath + obsidian.UrlSep + "meshes"
	ManageMeshPath       = BaseMeshesPath + obsidian.UrlSep + ":mesh_id"
	ManageMeshNamePath   = ManageMeshPath + obsidian.UrlSep + "name"
	ManageMeshConfigPath = ManageMeshPath + obsidian.UrlSep + "config"
)

// GetHandlers returns all obsidian handlers for Wifi
func GetHandlers() []obsidian.Handler {
	ret := []obsidian.Handler{
		handlers.GetListGatewaysHandler(BaseGatewaysPath, wifi.WifiGatewayType, makeWifiGateways),
		{Path: BaseGatewaysPath, Methods: obsidian.POST, HandlerFunc: createGateway},
		{Path: ManageGatewayPath, Methods: obsidian.GET, HandlerFunc: getGateway},
		{Path: ManageGatewayPath, Methods: obsidian.PUT, HandlerFunc: updateGateway},
		{Path: ManageGatewayPath, Methods: obsidian.DELETE, HandlerFunc: deleteGateway},
		{Path: ManageGatewayStatePath, Methods: obsidian.GET, HandlerFunc: handlers.GetStateHandler},

		{Path: BaseMeshesPath, Methods: obsidian.GET, HandlerFunc: listMeshes},
		{Path: BaseMeshesPath, Methods: obsidian.POST, HandlerFunc: createMesh},
		{Path: ManageMeshPath, Methods: obsidian.GET, HandlerFunc: getMesh},
		{Path: ManageMeshPath, Methods: obsidian.PUT, HandlerFunc: updateMesh},
		{Path: ManageMeshPath, Methods: obsidian.DELETE, HandlerFunc: deleteMesh},
	}
	ret = append(ret, handlers.GetTypedNetworkCRUDHandlers(BaseNetworksPath, ManageNetworkPath, wifi.WifiNetworkType, &wifimodels.WifiNetwork{})...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkNamePath, new(models.NetworkName), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkDescriptionPath, new(models.NetworkDescription), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkFeaturesPath, &orc8rmodels.NetworkFeatures{}, orc8r.NetworkFeaturesConfig)...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkWifiPath, &wifimodels.NetworkWifiConfigs{}, wifi.WifiNetworkType)...)

	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayNamePath, new(models.GatewayName))...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayDescriptionPath, new(models.GatewayDescription))...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayConfigPath, &orc8rmodels.MagmadGatewayConfigs{})...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayTierPath, new(orc8rmodels.TierID))...)
	ret = append(ret, handlers.GetGatewayDeviceHandlers(ManageGatewayDevicePath)...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayWifiPath, &wifimodels.GatewayWifiConfigs{})...)

	ret = append(ret, handlers.GetPartialEntityHandlers(ManageMeshNamePath, MeshID, new(wifimodels.MeshName))...)
	ret = append(ret, handlers.GetPartialEntityHandlers(ManageMeshConfigPath, MeshID, &wifimodels.MeshWifiConfigs{})...)

	return ret
}

type wifiAndMagmadGatewayEntities struct {
	wifiGatewayEnt, magmadEnt configurator.NetworkEntity
}

func makeWifiGateways(
	entsByTK map[storage.TypeAndKey]configurator.NetworkEntity,
	devicesByID map[string]interface{},
	statusesByID map[string]*orc8rmodels.GatewayStatus,
) map[string]handlers.GatewayModel {
	gatewayEntsByKey := map[string]*wifiAndMagmadGatewayEntities{}
	for tk, ent := range entsByTK {
		existing, found := gatewayEntsByKey[tk.Key]
		if !found {
			existing = &wifiAndMagmadGatewayEntities{}
			gatewayEntsByKey[tk.Key] = existing
		}
		switch ent.Type {
		case orc8r.MagmadGatewayType:
			existing.magmadEnt = ent
		case wifi.WifiGatewayType:
			existing.wifiGatewayEnt = ent
		}
	}

	ret := make(map[string]handlers.GatewayModel, len(gatewayEntsByKey))
	for key, wMEnts := range gatewayEntsByKey {
		hwID := wMEnts.magmadEnt.PhysicalID
		var devCasted *orc8rmodels.GatewayDevice
		if devicesByID[hwID] != nil {
			devCasted = devicesByID[hwID].(*orc8rmodels.GatewayDevice)
		}
		ret[key] = (&wifimodels.WifiGateway{}).FromBackendModels(wMEnts.magmadEnt, wMEnts.wifiGatewayEnt, devCasted, statusesByID[hwID])
	}
	return ret
}

func createGateway(c echo.Context) error {
	if nerr := handlers.CreateMagmadGatewayFromModel(c, &wifimodels.MutableWifiGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusCreated)
}

func getGateway(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	magmadModel, nerr := handlers.LoadMagmadGatewayModel(nid, gid)
	if nerr != nil {
		return nerr
	}

	ent, err := configurator.LoadEntity(
		nid, wifi.WifiGatewayType, gid,
		configurator.EntityLoadCriteria{LoadConfig: true, LoadAssocsFromThis: true},
	)
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to load wifi gateway"), http.StatusInternalServerError)
	}

	ret := &wifimodels.WifiGateway{
		ID:          magmadModel.ID,
		Name:        magmadModel.Name,
		Description: magmadModel.Description,
		Device:      magmadModel.Device,
		Status:      magmadModel.Status,
		Tier:        magmadModel.Tier,
		Magmad:      magmadModel.Magmad,
	}
	if ent.Config != nil {
		ret.Wifi = ent.Config.(*wifimodels.GatewayWifiConfigs)
	}

	return c.JSON(http.StatusOK, ret)
}

func updateGateway(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}
	if nerr = handlers.UpdateMagmadGatewayFromModel(c, nid, gid, &wifimodels.MutableWifiGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusNoContent)
}

func deleteGateway(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}
	gwEnt, err := configurator.LoadEntity(
		nid, orc8r.MagmadGatewayType, gid, configurator.EntityLoadCriteria{},
	)

	err = configurator.DeleteEntities(
		nid,
		[]storage.TypeAndKey{
			{Type: orc8r.MagmadGatewayType, Key: gid},
			{Type: wifi.WifiGatewayType, Key: gid},
		},
	)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	if gwEnt.PhysicalID != "" {
		err = device.DeleteDevice(nid, orc8r.AccessGatewayRecordType, gwEnt.PhysicalID)
		if err != nil {
			return obsidian.HttpError(errors.Wrap(err, "failed to delete device for gateway"), http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func listMeshes(c echo.Context) error {
	nid, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}

	ids, err := configurator.ListEntityKeys(nid, wifi.MeshEntityType)
	if err != nil {
		if err == merrors.ErrNotFound {
			return obsidian.HttpError(err, http.StatusNotFound)
		}
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	sort.Strings(ids)
	return c.JSON(http.StatusOK, ids)
}

func createMesh(c echo.Context) error {
	nid, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}

	payload := &wifimodels.WifiMesh{}
	if err := c.Bind(payload); err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}
	if err := payload.ValidateModel(); err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}

	gwIDs := []storage.TypeAndKey{}
	for _, gwID := range payload.GatewayIds {
		gwIDs = append(gwIDs, storage.TypeAndKey{Key: string(gwID), Type: orc8r.MagmadGatewayType})
	}

	_, err := configurator.CreateEntity(nid, configurator.NetworkEntity{
		Type:         wifi.MeshEntityType,
		Key:          string(payload.ID),
		Name:         string(payload.Name),
		Config:       payload.Config,
		Associations: gwIDs,
	})

	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

func getMesh(c echo.Context) error {
	nid, mid, nerr := GetNetworkAndMeshIDs(c)
	if nerr != nil {
		return nerr
	}

	ent, err := configurator.LoadEntity(nid, wifi.MeshEntityType, mid, configurator.FullEntityLoadCriteria())
	switch {
	case err == merrors.ErrNotFound:
		return echo.ErrNotFound
	case err != nil:
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	ret := (&wifimodels.WifiMesh{}).FromBackendModels(ent)
	return c.JSON(http.StatusOK, ret)
}

func updateMesh(c echo.Context) error {
	nid, mid, nerr := GetNetworkAndMeshIDs(c)
	if nerr != nil {
		return nerr
	}

	payload := &wifimodels.WifiMesh{}
	if err := c.Bind(payload); err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}
	if err := payload.ValidateModel(); err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}
	if string(payload.ID) != mid {
		return echo.NewHTTPError(http.StatusBadRequest, "mesh ID in body must match mesh_id in path")
	}

	ent, err := configurator.LoadEntity(nid, wifi.MeshEntityType, mid, configurator.FullEntityLoadCriteria())
	switch {
	case err == merrors.ErrNotFound:
		return echo.ErrNotFound
	case err != nil:
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	oldGWIds := []string{}
	newGWIds := []string{}
	for _, assoc := range ent.Associations {
		oldGWIds = append(oldGWIds, assoc.Key)
	}
	for _, gwId := range payload.GatewayIds {
		newGWIds = append(newGWIds, string(gwId))
	}
	if !reflect.DeepEqual(oldGWIds, newGWIds) {
		return echo.NewHTTPError(http.StatusBadRequest, "can't update gateways here! please update the individual gateways instead.")
	}

	_, err = configurator.UpdateEntities(nid, payload.ToUpdateCriteria())
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

func deleteMesh(c echo.Context) error {
	nid, mid, nerr := GetNetworkAndMeshIDs(c)
	if nerr != nil {
		return nerr
	}

	ent, err := configurator.LoadEntity(nid, wifi.MeshEntityType, mid, configurator.FullEntityLoadCriteria())
	switch {
	case err == merrors.ErrNotFound:
		return echo.ErrNotFound
	case err != nil:
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	// Don't allow the deletion if there are still associated gateways
	if len(ent.Associations) != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "can't delete a mesh with gateways!")
	}

	err = configurator.DeleteEntity(nid, wifi.MeshEntityType, mid)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

func GetNetworkAndMeshIDs(c echo.Context) (string, string, *echo.HTTPError) {
	vals, err := obsidian.GetParamValues(c, "network_id", "mesh_id")
	if err != nil {
		return "", "", err
	}
	return vals[0], vals[1], nil
}

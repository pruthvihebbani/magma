// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkFederationConfigs Federation configuration for a network
// swagger:model network_federation_configs
type NetworkFederationConfigs struct {

	// eap aka
	EapAka *NetworkFederationConfigsEapAka `json:"eap_aka,omitempty"`

	// gx
	Gx *NetworkFederationConfigsGx `json:"gx,omitempty"`

	// gy
	Gy *NetworkFederationConfigsGy `json:"gy,omitempty"`

	// health
	Health *NetworkFederationConfigsHealth `json:"health,omitempty"`

	// hss
	Hss *NetworkFederationConfigsHss `json:"hss,omitempty"`

	// s6a
	S6a *NetworkFederationConfigsS6a `json:"s6a,omitempty"`

	// served network ids
	ServedNetworkIds []string `json:"served_network_ids"`

	// swx
	Swx *NetworkFederationConfigsSwx `json:"swx,omitempty"`
}

// Validate validates this network federation configs
func (m *NetworkFederationConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEapAka(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS6a(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFederationConfigs) validateEapAka(formats strfmt.Registry) error {

	if swag.IsZero(m.EapAka) { // not required
		return nil
	}

	if m.EapAka != nil {
		if err := m.EapAka.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_aka")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigs) validateGx(formats strfmt.Registry) error {

	if swag.IsZero(m.Gx) { // not required
		return nil
	}

	if m.Gx != nil {
		if err := m.Gx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gx")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigs) validateGy(formats strfmt.Registry) error {

	if swag.IsZero(m.Gy) { // not required
		return nil
	}

	if m.Gy != nil {
		if err := m.Gy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gy")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigs) validateHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigs) validateHss(formats strfmt.Registry) error {

	if swag.IsZero(m.Hss) { // not required
		return nil
	}

	if m.Hss != nil {
		if err := m.Hss.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hss")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigs) validateS6a(formats strfmt.Registry) error {

	if swag.IsZero(m.S6a) { // not required
		return nil
	}

	if m.S6a != nil {
		if err := m.S6a.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s6a")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigs) validateSwx(formats strfmt.Registry) error {

	if swag.IsZero(m.Swx) { // not required
		return nil
	}

	if m.Swx != nil {
		if err := m.Swx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigs) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkFederationConfigsEapAka network federation configs eap aka
// swagger:model NetworkFederationConfigsEapAka
type NetworkFederationConfigsEapAka struct {

	// plmn ids
	PlmnIds []string `json:"plmn_ids"`

	// timeout
	Timeout *EapAkaTimeouts `json:"timeout,omitempty"`
}

// Validate validates this network federation configs eap aka
func (m *NetworkFederationConfigsEapAka) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlmnIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFederationConfigsEapAka) validatePlmnIds(formats strfmt.Registry) error {

	if swag.IsZero(m.PlmnIds) { // not required
		return nil
	}

	for i := 0; i < len(m.PlmnIds); i++ {

		if err := validate.MinLength("eap_aka"+"."+"plmn_ids"+"."+strconv.Itoa(i), "body", string(m.PlmnIds[i]), 5); err != nil {
			return err
		}

		if err := validate.MaxLength("eap_aka"+"."+"plmn_ids"+"."+strconv.Itoa(i), "body", string(m.PlmnIds[i]), 6); err != nil {
			return err
		}

		if err := validate.Pattern("eap_aka"+"."+"plmn_ids"+"."+strconv.Itoa(i), "body", string(m.PlmnIds[i]), `^(\d{5,6})$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *NetworkFederationConfigsEapAka) validateTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if m.Timeout != nil {
		if err := m.Timeout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_aka" + "." + "timeout")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigsEapAka) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigsEapAka) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigsEapAka
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkFederationConfigsGx network federation configs gx
// swagger:model NetworkFederationConfigsGx
type NetworkFederationConfigsGx struct {

	// server
	Server *DiameterClientConfigs `json:"server,omitempty"`
}

// Validate validates this network federation configs gx
func (m *NetworkFederationConfigsGx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFederationConfigsGx) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gx" + "." + "server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigsGx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigsGx) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigsGx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkFederationConfigsGy network federation configs gy
// swagger:model NetworkFederationConfigsGy
type NetworkFederationConfigsGy struct {

	// init method
	// Enum: [1 2]
	InitMethod *uint32 `json:"init_method,omitempty"`

	// server
	Server *DiameterClientConfigs `json:"server,omitempty"`
}

// Validate validates this network federation configs gy
func (m *NetworkFederationConfigsGy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInitMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkFederationConfigsGyTypeInitMethodPropEnum []interface{}

func init() {
	var res []uint32
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkFederationConfigsGyTypeInitMethodPropEnum = append(networkFederationConfigsGyTypeInitMethodPropEnum, v)
	}
}

// prop value enum
func (m *NetworkFederationConfigsGy) validateInitMethodEnum(path, location string, value uint32) error {
	if err := validate.Enum(path, location, value, networkFederationConfigsGyTypeInitMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkFederationConfigsGy) validateInitMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.InitMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateInitMethodEnum("gy"+"."+"init_method", "body", *m.InitMethod); err != nil {
		return err
	}

	return nil
}

func (m *NetworkFederationConfigsGy) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gy" + "." + "server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigsGy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigsGy) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigsGy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkFederationConfigsHealth network federation configs health
// swagger:model NetworkFederationConfigsHealth
type NetworkFederationConfigsHealth struct {

	// cloud disable period secs
	CloudDisablePeriodSecs uint32 `json:"cloud_disable_period_secs,omitempty"`

	// cpu utilization threshold
	CPUUtilizationThreshold float32 `json:"cpu_utilization_threshold,omitempty" magma_alt_name:"CpuUtilizationThreshold"`

	// FeG services for the health service to monitor
	HealthServices []string `json:"health_services"`

	// local disable period secs
	LocalDisablePeriodSecs uint32 `json:"local_disable_period_secs,omitempty"`

	// memory available threshold
	MemoryAvailableThreshold float32 `json:"memory_available_threshold,omitempty"`

	// minimum request threshold
	MinimumRequestThreshold uint32 `json:"minimum_request_threshold,omitempty"`

	// request failure threshold
	RequestFailureThreshold float32 `json:"request_failure_threshold,omitempty"`

	// update failure threshold
	UpdateFailureThreshold uint32 `json:"update_failure_threshold,omitempty"`

	// update interval secs
	UpdateIntervalSecs uint32 `json:"update_interval_secs,omitempty"`
}

// Validate validates this network federation configs health
func (m *NetworkFederationConfigsHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var networkFederationConfigsHealthHealthServicesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["S6A_PROXY","SESSION_PROXY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkFederationConfigsHealthHealthServicesItemsEnum = append(networkFederationConfigsHealthHealthServicesItemsEnum, v)
	}
}

func (m *NetworkFederationConfigsHealth) validateHealthServicesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkFederationConfigsHealthHealthServicesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkFederationConfigsHealth) validateHealthServices(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthServices) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthServices); i++ {

		// value enum
		if err := m.validateHealthServicesItemsEnum("health"+"."+"health_services"+"."+strconv.Itoa(i), "body", m.HealthServices[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigsHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigsHealth) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigsHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkFederationConfigsHss network federation configs hss
// swagger:model NetworkFederationConfigsHss
type NetworkFederationConfigsHss struct {

	// default sub profile
	DefaultSubProfile *SubscriptionProfile `json:"default_sub_profile,omitempty"`

	// lte auth amf
	// Format: byte
	LteAuthAmf strfmt.Base64 `json:"lte_auth_amf,omitempty"`

	// lte auth op
	// Format: byte
	LteAuthOp strfmt.Base64 `json:"lte_auth_op,omitempty"`

	// server
	Server *DiameterServerConfigs `json:"server,omitempty"`

	// stream subscribers
	StreamSubscribers bool `json:"stream_subscribers,omitempty"`

	// sub profiles
	SubProfiles map[string]SubscriptionProfile `json:"sub_profiles,omitempty"`
}

// Validate validates this network federation configs hss
func (m *NetworkFederationConfigsHss) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultSubProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLteAuthAmf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLteAuthOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFederationConfigsHss) validateDefaultSubProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultSubProfile) { // not required
		return nil
	}

	if m.DefaultSubProfile != nil {
		if err := m.DefaultSubProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hss" + "." + "default_sub_profile")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigsHss) validateLteAuthAmf(formats strfmt.Registry) error {

	if swag.IsZero(m.LteAuthAmf) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *NetworkFederationConfigsHss) validateLteAuthOp(formats strfmt.Registry) error {

	if swag.IsZero(m.LteAuthOp) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *NetworkFederationConfigsHss) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hss" + "." + "server")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkFederationConfigsHss) validateSubProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.SubProfiles) { // not required
		return nil
	}

	for k := range m.SubProfiles {

		if err := validate.Required("hss"+"."+"sub_profiles"+"."+k, "body", m.SubProfiles[k]); err != nil {
			return err
		}
		if val, ok := m.SubProfiles[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigsHss) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigsHss) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigsHss
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkFederationConfigsS6a network federation configs s6a
// swagger:model NetworkFederationConfigsS6a
type NetworkFederationConfigsS6a struct {

	// server
	Server *DiameterClientConfigs `json:"server,omitempty"`
}

// Validate validates this network federation configs s6a
func (m *NetworkFederationConfigsS6a) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFederationConfigsS6a) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s6a" + "." + "server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigsS6a) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigsS6a) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigsS6a
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkFederationConfigsSwx network federation configs swx
// swagger:model NetworkFederationConfigsSwx
type NetworkFederationConfigsSwx struct {

	// cache TTL seconds
	CacheTTLSeconds uint32 `json:"cache_TTL_seconds,omitempty"`

	// server
	Server *DiameterClientConfigs `json:"server,omitempty"`

	// verify authorization
	VerifyAuthorization bool `json:"verify_authorization,omitempty"`
}

// Validate validates this network federation configs swx
func (m *NetworkFederationConfigsSwx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkFederationConfigsSwx) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swx" + "." + "server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkFederationConfigsSwx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkFederationConfigsSwx) UnmarshalBinary(b []byte) error {
	var res NetworkFederationConfigsSwx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

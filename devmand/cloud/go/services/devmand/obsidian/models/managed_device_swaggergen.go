// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ManagedDevice Management properties for a given device
// swagger:model managed_device
type ManagedDevice struct {

	// channels
	Channels *ManagedDeviceChannels `json:"channels,omitempty"`

	// device config
	DeviceConfig string `json:"device_config,omitempty"`

	// device type
	DeviceType []string `json:"device_type"`

	// host
	Host string `json:"host,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`
}

// Validate validates this managed device
func (m *ManagedDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagedDevice) validateChannels(formats strfmt.Registry) error {

	if swag.IsZero(m.Channels) { // not required
		return nil
	}

	if m.Channels != nil {
		if err := m.Channels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channels")
			}
			return err
		}
	}

	return nil
}

func (m *ManagedDevice) validateDeviceType(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceType) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceType); i++ {

		if err := validate.MinLength("device_type"+"."+strconv.Itoa(i), "body", string(m.DeviceType[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagedDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagedDevice) UnmarshalBinary(b []byte) error {
	var res ManagedDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ManagedDeviceChannels managed device channels
// swagger:model ManagedDeviceChannels
type ManagedDeviceChannels struct {

	// cambium channel
	CambiumChannel *CambiumChannel `json:"cambium_channel,omitempty"`

	// frinx channel
	FrinxChannel *FrinxChannel `json:"frinx_channel,omitempty"`

	// other channel
	OtherChannel *OtherChannel `json:"other_channel,omitempty"`

	// snmp channel
	SnmpChannel *SnmpChannel `json:"snmp_channel,omitempty"`
}

// Validate validates this managed device channels
func (m *ManagedDeviceChannels) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCambiumChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrinxChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnmpChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagedDeviceChannels) validateCambiumChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.CambiumChannel) { // not required
		return nil
	}

	if m.CambiumChannel != nil {
		if err := m.CambiumChannel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channels" + "." + "cambium_channel")
			}
			return err
		}
	}

	return nil
}

func (m *ManagedDeviceChannels) validateFrinxChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.FrinxChannel) { // not required
		return nil
	}

	if m.FrinxChannel != nil {
		if err := m.FrinxChannel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channels" + "." + "frinx_channel")
			}
			return err
		}
	}

	return nil
}

func (m *ManagedDeviceChannels) validateOtherChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.OtherChannel) { // not required
		return nil
	}

	if m.OtherChannel != nil {
		if err := m.OtherChannel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channels" + "." + "other_channel")
			}
			return err
		}
	}

	return nil
}

func (m *ManagedDeviceChannels) validateSnmpChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.SnmpChannel) { // not required
		return nil
	}

	if m.SnmpChannel != nil {
		if err := m.SnmpChannel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channels" + "." + "snmp_channel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagedDeviceChannels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagedDeviceChannels) UnmarshalBinary(b []byte) error {
	var res ManagedDeviceChannels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

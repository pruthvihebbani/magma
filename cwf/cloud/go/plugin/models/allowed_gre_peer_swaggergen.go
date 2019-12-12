// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AllowedGrePeer allowed gre peer
// swagger:model allowed_gre_peer
type AllowedGrePeer struct {

	// ip
	// Required: true
	// Max Length: 49
	// Min Length: 5
	IP string `json:"ip"`

	// key
	// Required: true
	Key *uint32 `json:"key"`
}

// Validate validates this allowed gre peer
func (m *AllowedGrePeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllowedGrePeer) validateIP(formats strfmt.Registry) error {

	if err := validate.RequiredString("ip", "body", string(m.IP)); err != nil {
		return err
	}

	if err := validate.MinLength("ip", "body", string(m.IP), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("ip", "body", string(m.IP), 49); err != nil {
		return err
	}

	return nil
}

func (m *AllowedGrePeer) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllowedGrePeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllowedGrePeer) UnmarshalBinary(b []byte) error {
	var res AllowedGrePeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FlowMatch flow match
// swagger:model flow_match
type FlowMatch struct {

	// direction
	// Required: true
	// Enum: [UPLINK DOWNLINK]
	Direction *string `json:"direction"`

	// ip dst
	IPDst *IPAddress `json:"ip_dst,omitempty" magma_alt_name:"IpDst"`

	// ip proto
	// Required: true
	// Enum: [IPPROTO_IP IPPROTO_TCP IPPROTO_UDP IPPROTO_ICMP]
	IPProto *string `json:"ip_proto"`

	// ip src
	IPSrc *IPAddress `json:"ip_src,omitempty" magma_alt_name:"IpSrc"`

	// ipv4 dst
	IPV4Dst string `json:"ipv4_dst,omitempty" magma_alt_name:"Ipv4Dst"`

	// ipv4 src
	IPV4Src string `json:"ipv4_src,omitempty" magma_alt_name:"Ipv4Src"`

	// tcp dst
	TCPDst uint32 `json:"tcp_dst,omitempty" magma_alt_name:"TcpDst"`

	// tcp src
	TCPSrc uint32 `json:"tcp_src,omitempty" magma_alt_name:"TcpSrc"`

	// udp dst
	UDPDst uint32 `json:"udp_dst,omitempty" magma_alt_name:"UdpDst"`

	// udp src
	UDPSrc uint32 `json:"udp_src,omitempty" magma_alt_name:"UdpSrc"`
}

// Validate validates this flow match
func (m *FlowMatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPDst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPProto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPSrc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var flowMatchTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UPLINK","DOWNLINK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowMatchTypeDirectionPropEnum = append(flowMatchTypeDirectionPropEnum, v)
	}
}

const (

	// FlowMatchDirectionUPLINK captures enum value "UPLINK"
	FlowMatchDirectionUPLINK string = "UPLINK"

	// FlowMatchDirectionDOWNLINK captures enum value "DOWNLINK"
	FlowMatchDirectionDOWNLINK string = "DOWNLINK"
)

// prop value enum
func (m *FlowMatch) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, flowMatchTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlowMatch) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *FlowMatch) validateIPDst(formats strfmt.Registry) error {

	if swag.IsZero(m.IPDst) { // not required
		return nil
	}

	if m.IPDst != nil {
		if err := m.IPDst.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_dst")
			}
			return err
		}
	}

	return nil
}

var flowMatchTypeIPProtoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPPROTO_IP","IPPROTO_TCP","IPPROTO_UDP","IPPROTO_ICMP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		flowMatchTypeIPProtoPropEnum = append(flowMatchTypeIPProtoPropEnum, v)
	}
}

const (

	// FlowMatchIPProtoIPPROTOIP captures enum value "IPPROTO_IP"
	FlowMatchIPProtoIPPROTOIP string = "IPPROTO_IP"

	// FlowMatchIPProtoIPPROTOTCP captures enum value "IPPROTO_TCP"
	FlowMatchIPProtoIPPROTOTCP string = "IPPROTO_TCP"

	// FlowMatchIPProtoIPPROTOUDP captures enum value "IPPROTO_UDP"
	FlowMatchIPProtoIPPROTOUDP string = "IPPROTO_UDP"

	// FlowMatchIPProtoIPPROTOICMP captures enum value "IPPROTO_ICMP"
	FlowMatchIPProtoIPPROTOICMP string = "IPPROTO_ICMP"
)

// prop value enum
func (m *FlowMatch) validateIPProtoEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, flowMatchTypeIPProtoPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FlowMatch) validateIPProto(formats strfmt.Registry) error {

	if err := validate.Required("ip_proto", "body", m.IPProto); err != nil {
		return err
	}

	// value enum
	if err := m.validateIPProtoEnum("ip_proto", "body", *m.IPProto); err != nil {
		return err
	}

	return nil
}

func (m *FlowMatch) validateIPSrc(formats strfmt.Registry) error {

	if swag.IsZero(m.IPSrc) { // not required
		return nil
	}

	if m.IPSrc != nil {
		if err := m.IPSrc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_src")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowMatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowMatch) UnmarshalBinary(b []byte) error {
	var res FlowMatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

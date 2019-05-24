// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DiskPartition disk partition
// swagger:model disk_partition
type DiskPartition struct {

	// Name of the device
	Device string `json:"device,omitempty"`

	// Free disk space of the device in bytes
	Free uint64 `json:"free,omitempty"`

	// Mount point of the device
	MountPoint string `json:"mount_point,omitempty"`

	// Total disk space of the device in bytes
	Total uint64 `json:"total,omitempty"`

	// Used disk space of the device in bytes
	Used uint64 `json:"used,omitempty"`
}

// Validate validates this disk partition
func (m *DiskPartition) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiskPartition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiskPartition) UnmarshalBinary(b []byte) error {
	var res DiskPartition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

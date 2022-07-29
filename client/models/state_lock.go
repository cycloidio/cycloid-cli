// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StateLock State Lock
//
// The Lock management of a State in the Inventory of the Project's environment
// swagger:model StateLock
type StateLock struct {

	// created
	Created string `json:"created,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// info
	Info string `json:"info,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// who
	Who string `json:"who,omitempty"`
}

// Validate validates this state lock
func (m *StateLock) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateLock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateLock) UnmarshalBinary(b []byte) error {
	var res StateLock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
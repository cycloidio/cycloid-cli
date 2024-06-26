// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMwareVsphere Representation of VMwareVsphere external backend.
// Must be matched with a credential of the "vmware" type.
//
// swagger:model VMwareVsphere
type VMwareVsphere struct {

	// Whether verification of SSL certificate should be disabled.
	//
	AllowUnverifiedSsl bool `json:"allow_unverified_ssl,omitempty"`

	// FQDN or IP address of the vCenter server.
	//
	Server string `json:"server,omitempty"`
}

// Engine gets the engine of this subtype
func (m *VMwareVsphere) Engine() string {
	return "VMwareVsphere"
}

// SetEngine sets the engine of this subtype
func (m *VMwareVsphere) SetEngine(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *VMwareVsphere) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Whether verification of SSL certificate should be disabled.
		//
		AllowUnverifiedSsl bool `json:"allow_unverified_ssl,omitempty"`

		// FQDN or IP address of the vCenter server.
		//
		Server string `json:"server,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Engine string `json:"engine"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result VMwareVsphere

	if base.Engine != result.Engine() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid engine value: %q", base.Engine)
	}

	result.AllowUnverifiedSsl = data.AllowUnverifiedSsl
	result.Server = data.Server

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m VMwareVsphere) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Whether verification of SSL certificate should be disabled.
		//
		AllowUnverifiedSsl bool `json:"allow_unverified_ssl,omitempty"`

		// FQDN or IP address of the vCenter server.
		//
		Server string `json:"server,omitempty"`
	}{

		AllowUnverifiedSsl: m.AllowUnverifiedSsl,

		Server: m.Server,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Engine string `json:"engine"`
	}{

		Engine: m.Engine(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this v mware vsphere
func (m *VMwareVsphere) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this v mware vsphere based on the context it is used
func (m *VMwareVsphere) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VMwareVsphere) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMwareVsphere) UnmarshalBinary(b []byte) error {
	var res VMwareVsphere
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

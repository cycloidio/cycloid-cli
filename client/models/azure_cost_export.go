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

// AzureCostExport Representation of AzureCostExport external backend.
// Must be matched with a credential of the "azure" type.
//
// swagger:model AzureCostExport
type AzureCostExport struct {

	// Endpoint of blob storage service containing the export files.
	BlobServiceURL string `json:"blob_service_url,omitempty"`

	// Name of the export
	Name string `json:"name,omitempty"`

	// Scope of the export
	Scope string `json:"scope,omitempty"`
}

// Engine gets the engine of this subtype
func (m *AzureCostExport) Engine() string {
	return "AzureCostExport"
}

// SetEngine sets the engine of this subtype
func (m *AzureCostExport) SetEngine(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AzureCostExport) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Endpoint of blob storage service containing the export files.
		BlobServiceURL string `json:"blob_service_url,omitempty"`

		// Name of the export
		Name string `json:"name,omitempty"`

		// Scope of the export
		Scope string `json:"scope,omitempty"`
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

	var result AzureCostExport

	if base.Engine != result.Engine() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid engine value: %q", base.Engine)
	}

	result.BlobServiceURL = data.BlobServiceURL
	result.Name = data.Name
	result.Scope = data.Scope

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AzureCostExport) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Endpoint of blob storage service containing the export files.
		BlobServiceURL string `json:"blob_service_url,omitempty"`

		// Name of the export
		Name string `json:"name,omitempty"`

		// Scope of the export
		Scope string `json:"scope,omitempty"`
	}{

		BlobServiceURL: m.BlobServiceURL,

		Name: m.Name,

		Scope: m.Scope,
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

// Validate validates this azure cost export
func (m *AzureCostExport) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this azure cost export based on the context it is used
func (m *AzureCostExport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureCostExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCostExport) UnmarshalBinary(b []byte) error {
	var res AzureCostExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

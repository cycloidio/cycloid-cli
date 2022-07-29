// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GCPStorage Representation of GCP remote tf state for external backend.
// Must be matched with a credential of the "gcp" type.
//
// swagger:model GCPStorage
type GCPStorage struct {

	// The GCP bucket containing objects
	//
	// Required: true
	Bucket *string `json:"bucket"`

	// The GCP object uniquely identifying an object in a bucket,
	// will be required if the EB is not default
	//
	Object string `json:"object,omitempty"`
}

// Engine gets the engine of this subtype
func (m *GCPStorage) Engine() string {
	return "GCPStorage"
}

// SetEngine sets the engine of this subtype
func (m *GCPStorage) SetEngine(val string) {

}

// Bucket gets the bucket of this subtype

// Object gets the object of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GCPStorage) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The GCP bucket containing objects
		//
		// Required: true
		Bucket *string `json:"bucket"`

		// The GCP object uniquely identifying an object in a bucket,
		// will be required if the EB is not default
		//
		Object string `json:"object,omitempty"`
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

	var result GCPStorage

	if base.Engine != result.Engine() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid engine value: %q", base.Engine)
	}

	result.Bucket = data.Bucket

	result.Object = data.Object

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GCPStorage) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The GCP bucket containing objects
		//
		// Required: true
		Bucket *string `json:"bucket"`

		// The GCP object uniquely identifying an object in a bucket,
		// will be required if the EB is not default
		//
		Object string `json:"object,omitempty"`
	}{

		Bucket: m.Bucket,

		Object: m.Object,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Engine string `json:"engine"`
	}{

		Engine: m.Engine(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this g c p storage
func (m *GCPStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GCPStorage) validateBucket(formats strfmt.Registry) error {

	if err := validate.Required("bucket", "body", m.Bucket); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GCPStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPStorage) UnmarshalBinary(b []byte) error {
	var res GCPStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
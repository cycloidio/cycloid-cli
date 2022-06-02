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

// SwiftRemoteTFState Deprecated. Please use SwiftStorage.
// Representation of Swift remote tf state for external backend.
// Must be matched with a credential of the "swift" type.
//
// swagger:model SwiftRemoteTFState
type SwiftRemoteTFState struct {

	// The Swift container containing objects
	//
	// Required: true
	Container *string `json:"container"`

	// The swift object uniquely identifying an object in a container
	//
	Object string `json:"object,omitempty"`

	// The Swift region where the resource exists
	//
	// Required: true
	Region *string `json:"region"`

	// Set this to `true` to not verify SSL certificates
	//
	SkipVerifySsl bool `json:"skip_verify_ssl,omitempty"`
}

// Engine gets the engine of this subtype
func (m *SwiftRemoteTFState) Engine() string {
	return "SwiftRemoteTFState"
}

// SetEngine sets the engine of this subtype
func (m *SwiftRemoteTFState) SetEngine(val string) {

}

// Container gets the container of this subtype

// Object gets the object of this subtype

// Region gets the region of this subtype

// SkipVerifySsl gets the skip verify ssl of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SwiftRemoteTFState) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The Swift container containing objects
		//
		// Required: true
		Container *string `json:"container"`

		// The swift object uniquely identifying an object in a container
		//
		Object string `json:"object,omitempty"`

		// The Swift region where the resource exists
		//
		// Required: true
		Region *string `json:"region"`

		// Set this to `true` to not verify SSL certificates
		//
		SkipVerifySsl bool `json:"skip_verify_ssl,omitempty"`
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

	var result SwiftRemoteTFState

	if base.Engine != result.Engine() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid engine value: %q", base.Engine)
	}

	result.Container = data.Container

	result.Object = data.Object

	result.Region = data.Region

	result.SkipVerifySsl = data.SkipVerifySsl

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SwiftRemoteTFState) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The Swift container containing objects
		//
		// Required: true
		Container *string `json:"container"`

		// The swift object uniquely identifying an object in a container
		//
		Object string `json:"object,omitempty"`

		// The Swift region where the resource exists
		//
		// Required: true
		Region *string `json:"region"`

		// Set this to `true` to not verify SSL certificates
		//
		SkipVerifySsl bool `json:"skip_verify_ssl,omitempty"`
	}{

		Container: m.Container,

		Object: m.Object,

		Region: m.Region,

		SkipVerifySsl: m.SkipVerifySsl,
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

// Validate validates this swift remote t f state
func (m *SwiftRemoteTFState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwiftRemoteTFState) validateContainer(formats strfmt.Registry) error {

	if err := validate.Required("container", "body", m.Container); err != nil {
		return err
	}

	return nil
}

func (m *SwiftRemoteTFState) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwiftRemoteTFState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwiftRemoteTFState) UnmarshalBinary(b []byte) error {
	var res SwiftRemoteTFState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

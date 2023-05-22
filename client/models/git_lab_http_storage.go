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

// GitLabHTTPStorage Representation of GitLab HTTP storage for external backend.
// Must be matched with a credential of the "basic_auth" type.
//
// swagger:model GitLabHTTPStorage
type GitLabHTTPStorage struct {

	// The URL endpoint to use
	//
	// Required: true
	URL *string `json:"url"`
}

// Engine gets the engine of this subtype
func (m *GitLabHTTPStorage) Engine() string {
	return "GitLabHTTPStorage"
}

// SetEngine sets the engine of this subtype
func (m *GitLabHTTPStorage) SetEngine(val string) {

}

// URL gets the url of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GitLabHTTPStorage) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The URL endpoint to use
		//
		// Required: true
		URL *string `json:"url"`
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

	var result GitLabHTTPStorage

	if base.Engine != result.Engine() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid engine value: %q", base.Engine)
	}

	result.URL = data.URL

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GitLabHTTPStorage) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The URL endpoint to use
		//
		// Required: true
		URL *string `json:"url"`
	}{

		URL: m.URL,
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

// Validate validates this git lab HTTP storage
func (m *GitLabHTTPStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitLabHTTPStorage) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitLabHTTPStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitLabHTTPStorage) UnmarshalBinary(b []byte) error {
	var res GitLabHTTPStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

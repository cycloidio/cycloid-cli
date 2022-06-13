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

// AWSRemoteTFState Deprecated. Please use AWSStorage.
// Representation of AWS remote tf state for external backend.
// Must be matched with a credential of the "aws" type.
//
// swagger:model AWSRemoteTFState
type AWSRemoteTFState struct {

	// The AWS bucket containing objects
	//
	// Required: true
	Bucket *string `json:"bucket"`

	// A custom endpoint for the S3 API (default: s3.amazonaws.com)
	//
	Endpoint string `json:"endpoint,omitempty"`

	// The S3 Key uniquely identifies an object in a bucket, will
	// be required if the EB is not the default one.
	//
	Key string `json:"key,omitempty"`

	// The AWS region where the resource exists
	//
	// Required: true
	Region *string `json:"region"`

	// Always use path-style S3 URLs (https://<HOST>/<BUCKET> instead of https://<BUCKET>.<HOST>)
	//
	S3ForcePathStyle bool `json:"s3_force_path_style,omitempty"`

	// Set this to `true` to not verify SSL certificates
	//
	SkipVerifySsl bool `json:"skip_verify_ssl,omitempty"`
}

// Engine gets the engine of this subtype
func (m *AWSRemoteTFState) Engine() string {
	return "AWSRemoteTFState"
}

// SetEngine sets the engine of this subtype
func (m *AWSRemoteTFState) SetEngine(val string) {

}

// Bucket gets the bucket of this subtype

// Endpoint gets the endpoint of this subtype

// Key gets the key of this subtype

// Region gets the region of this subtype

// S3ForcePathStyle gets the s3 force path style of this subtype

// SkipVerifySsl gets the skip verify ssl of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AWSRemoteTFState) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The AWS bucket containing objects
		//
		// Required: true
		Bucket *string `json:"bucket"`

		// A custom endpoint for the S3 API (default: s3.amazonaws.com)
		//
		Endpoint string `json:"endpoint,omitempty"`

		// The S3 Key uniquely identifies an object in a bucket, will
		// be required if the EB is not the default one.
		//
		Key string `json:"key,omitempty"`

		// The AWS region where the resource exists
		//
		// Required: true
		Region *string `json:"region"`

		// Always use path-style S3 URLs (https://<HOST>/<BUCKET> instead of https://<BUCKET>.<HOST>)
		//
		S3ForcePathStyle bool `json:"s3_force_path_style,omitempty"`

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

	var result AWSRemoteTFState

	if base.Engine != result.Engine() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid engine value: %q", base.Engine)
	}

	result.Bucket = data.Bucket

	result.Endpoint = data.Endpoint

	result.Key = data.Key

	result.Region = data.Region

	result.S3ForcePathStyle = data.S3ForcePathStyle

	result.SkipVerifySsl = data.SkipVerifySsl

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AWSRemoteTFState) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The AWS bucket containing objects
		//
		// Required: true
		Bucket *string `json:"bucket"`

		// A custom endpoint for the S3 API (default: s3.amazonaws.com)
		//
		Endpoint string `json:"endpoint,omitempty"`

		// The S3 Key uniquely identifies an object in a bucket, will
		// be required if the EB is not the default one.
		//
		Key string `json:"key,omitempty"`

		// The AWS region where the resource exists
		//
		// Required: true
		Region *string `json:"region"`

		// Always use path-style S3 URLs (https://<HOST>/<BUCKET> instead of https://<BUCKET>.<HOST>)
		//
		S3ForcePathStyle bool `json:"s3_force_path_style,omitempty"`

		// Set this to `true` to not verify SSL certificates
		//
		SkipVerifySsl bool `json:"skip_verify_ssl,omitempty"`
	}{

		Bucket: m.Bucket,

		Endpoint: m.Endpoint,

		Key: m.Key,

		Region: m.Region,

		S3ForcePathStyle: m.S3ForcePathStyle,

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

// Validate validates this a w s remote t f state
func (m *AWSRemoteTFState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
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

func (m *AWSRemoteTFState) validateBucket(formats strfmt.Registry) error {

	if err := validate.Required("bucket", "body", m.Bucket); err != nil {
		return err
	}

	return nil
}

func (m *AWSRemoteTFState) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AWSRemoteTFState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSRemoteTFState) UnmarshalBinary(b []byte) error {
	var res AWSRemoteTFState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

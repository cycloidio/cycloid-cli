// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// CIVersion CIVersion
//
// # Represent a version of a resource
//
// swagger:model CIVersion
type CIVersion map[string]string

// Validate validates this c i version
func (m CIVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this c i version based on context it is used
func (m CIVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

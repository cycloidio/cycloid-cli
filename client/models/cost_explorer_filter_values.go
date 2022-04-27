// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CostExplorerFilterValues CostExplorerFilterValues
//
// A collection of the values for whitch the cost can be filtered, i.e. all
// the providers, services, reginos, etc... for one org.
//
// swagger:model CostExplorerFilterValues
type CostExplorerFilterValues struct {

	// categories
	Categories []string `json:"categories"`

	// currencies
	Currencies []string `json:"currencies"`

	// environments
	Environments []string `json:"environments"`

	// instance types
	InstanceTypes []string `json:"instance_types"`

	// linked accounts
	LinkedAccounts []string `json:"linked_accounts"`

	// master accounts
	MasterAccounts []string `json:"master_accounts"`

	// projects
	Projects []string `json:"projects"`

	// providers
	Providers []string `json:"providers"`

	// regions
	Regions []string `json:"regions"`

	// resources
	Resources []string `json:"resources"`

	// services
	Services []string `json:"services"`

	// tags
	Tags interface{} `json:"tags,omitempty"`
}

// Validate validates this cost explorer filter values
func (m *CostExplorerFilterValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CostExplorerFilterValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostExplorerFilterValues) UnmarshalBinary(b []byte) error {
	var res CostExplorerFilterValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

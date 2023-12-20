// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CloudCostManagementFilterValues CloudCostManagementFilterValues
//
// A collection of the values for which the cost can be filtered, i.e. all
// the providers, services, regions, etc... for one org.
//
// swagger:model CloudCostManagementFilterValues
type CloudCostManagementFilterValues struct {

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

	// linked accounts with cp
	LinkedAccountsWithCp interface{} `json:"linked_accounts_with_cp,omitempty"`

	// master accounts
	MasterAccounts []string `json:"master_accounts"`

	// master accounts with cp
	MasterAccountsWithCp interface{} `json:"master_accounts_with_cp,omitempty"`

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
	Tags map[string][]string `json:"tags,omitempty"`
}

// Validate validates this cloud cost management filter values
func (m *CloudCostManagementFilterValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudCostManagementFilterValues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudCostManagementFilterValues) UnmarshalBinary(b []byte) error {
	var res CloudCostManagementFilterValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

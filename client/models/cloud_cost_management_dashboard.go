// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudCostManagementDashboard CloudCostManagementDashboard
//
// The dashboard of the Cloud Cost Management, it contains
//   - a histogram of the cost in the period aggregated by provider and
//     by time granularity
//   - a histogram of the cost aggregated by the top projects and providers
//     and filtered by the top projects
//   - a list of resources and relative cost for each top projects
//   - a map containing  properties that can be specified filtering the
//     returned results, with a set of valid values for each.
//
// swagger:model CloudCostManagementDashboard
type CloudCostManagementDashboard struct {

	// filter values
	// Required: true
	FilterValues *CloudCostManagementFilterValues `json:"filter_values"`

	// project resources
	// Required: true
	ProjectResources []*CloudCostManagementProjectResources `json:"project_resources"`

	// projects
	// Required: true
	Projects *CloudCostManagementHistogram `json:"projects"`

	// providers
	// Required: true
	Providers *CloudCostManagementHistogram `json:"providers"`

	// total co2e
	TotalCo2e float64 `json:"total_co2e,omitempty"`

	// total cost
	TotalCost float64 `json:"total_cost,omitempty"`

	// total kwh
	TotalKwh float64 `json:"total_kwh,omitempty"`
}

// Validate validates this cloud cost management dashboard
func (m *CloudCostManagementDashboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudCostManagementDashboard) validateFilterValues(formats strfmt.Registry) error {

	if err := validate.Required("filter_values", "body", m.FilterValues); err != nil {
		return err
	}

	if m.FilterValues != nil {
		if err := m.FilterValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter_values")
			}
			return err
		}
	}

	return nil
}

func (m *CloudCostManagementDashboard) validateProjectResources(formats strfmt.Registry) error {

	if err := validate.Required("project_resources", "body", m.ProjectResources); err != nil {
		return err
	}

	for i := 0; i < len(m.ProjectResources); i++ {
		if swag.IsZero(m.ProjectResources[i]) { // not required
			continue
		}

		if m.ProjectResources[i] != nil {
			if err := m.ProjectResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudCostManagementDashboard) validateProjects(formats strfmt.Registry) error {

	if err := validate.Required("projects", "body", m.Projects); err != nil {
		return err
	}

	if m.Projects != nil {
		if err := m.Projects.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projects")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("projects")
			}
			return err
		}
	}

	return nil
}

func (m *CloudCostManagementDashboard) validateProviders(formats strfmt.Registry) error {

	if err := validate.Required("providers", "body", m.Providers); err != nil {
		return err
	}

	if m.Providers != nil {
		if err := m.Providers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("providers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("providers")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloud cost management dashboard based on the context it is used
func (m *CloudCostManagementDashboard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilterValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProviders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudCostManagementDashboard) contextValidateFilterValues(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterValues != nil {

		if err := m.FilterValues.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter_values")
			}
			return err
		}
	}

	return nil
}

func (m *CloudCostManagementDashboard) contextValidateProjectResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProjectResources); i++ {

		if m.ProjectResources[i] != nil {

			if swag.IsZero(m.ProjectResources[i]) { // not required
				return nil
			}

			if err := m.ProjectResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("project_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("project_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudCostManagementDashboard) contextValidateProjects(ctx context.Context, formats strfmt.Registry) error {

	if m.Projects != nil {

		if err := m.Projects.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projects")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("projects")
			}
			return err
		}
	}

	return nil
}

func (m *CloudCostManagementDashboard) contextValidateProviders(ctx context.Context, formats strfmt.Registry) error {

	if m.Providers != nil {

		if err := m.Providers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("providers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("providers")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudCostManagementDashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudCostManagementDashboard) UnmarshalBinary(b []byte) error {
	var res CloudCostManagementDashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

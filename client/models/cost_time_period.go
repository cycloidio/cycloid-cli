// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CostTimePeriod CostTimePeriod
//
// This object contains the items described in https://docs.aws.amazon.com/sdk-for-go/api/service/costexplorer/#DateInterval It defines the beginning and the end of the time frame for which, the API should gather costs.
// swagger:model CostTimePeriod
type CostTimePeriod struct {

	// begin
	// Required: true
	// Max Length: 10
	// Min Length: 10
	// Pattern: ^[0-9]{4}-[0-9]{2}-[0-9]{2}$
	Begin *string `json:"begin"`

	// end
	// Required: true
	// Max Length: 10
	// Min Length: 10
	// Pattern: ^[0-9]{4}-[0-9]{2}-[0-9]{2}$
	End *string `json:"end"`
}

// Validate validates this cost time period
func (m *CostTimePeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBegin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CostTimePeriod) validateBegin(formats strfmt.Registry) error {

	if err := validate.Required("begin", "body", m.Begin); err != nil {
		return err
	}

	if err := validate.MinLength("begin", "body", string(*m.Begin), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("begin", "body", string(*m.Begin), 10); err != nil {
		return err
	}

	if err := validate.Pattern("begin", "body", string(*m.Begin), `^[0-9]{4}-[0-9]{2}-[0-9]{2}$`); err != nil {
		return err
	}

	return nil
}

func (m *CostTimePeriod) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	if err := validate.MinLength("end", "body", string(*m.End), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("end", "body", string(*m.End), 10); err != nil {
		return err
	}

	if err := validate.Pattern("end", "body", string(*m.End), `^[0-9]{4}-[0-9]{2}-[0-9]{2}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CostTimePeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostTimePeriod) UnmarshalBinary(b []byte) error {
	var res CostTimePeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

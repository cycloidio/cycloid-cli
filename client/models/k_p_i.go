// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KPI KPI
//
// A KPI
// swagger:model KPI
type KPI struct {

	// canonical
	// Required: true
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	Canonical *string `json:"canonical"`

	// The config represent some extra parameters which are required for the configuration of certain KPIs. Please refer to the documentation for more details.
	//
	// Required: true
	Config interface{} `json:"config"`

	// created at
	// Required: true
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at"`

	// The data is represented in CSV format. Depending on the widget configured for the KPI, the format may vary. For more information please refer to our product documentation.
	DataSet []interface{} `json:"data_set"`

	// description
	// Required: true
	Description *string `json:"description"`

	// environment canonical
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$
	EnvironmentCanonical string `json:"environment_canonical,omitempty"`

	// id
	// Required: true
	// Minimum: 1
	ID *uint32 `json:"id"`

	// job name
	JobName string `json:"job_name,omitempty"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// pipeline name
	PipelineName string `json:"pipeline_name,omitempty"`

	// project canonical
	// Max Length: 100
	// Min Length: 3
	// Pattern: ^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$
	ProjectCanonical string `json:"project_canonical,omitempty"`

	// type
	// Required: true
	// Enum: [build_avg_time build_frequency build_history code_coverage time_to_release]
	Type *string `json:"type"`

	// updated at
	// Required: true
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at"`

	// widget
	// Required: true
	// Enum: [bars stackbars doughnut history line pie summary]
	Widget *string `json:"widget"`
}

// Validate validates this k p i
func (m *KPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectCanonical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KPI) validateCanonical(formats strfmt.Registry) error {

	if err := validate.Required("canonical", "body", m.Canonical); err != nil {
		return err
	}

	if err := validate.MinLength("canonical", "body", string(*m.Canonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("canonical", "body", string(*m.Canonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("canonical", "body", string(*m.Canonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("created_at", "body", int64(*m.CreatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateEnvironmentCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("environment_canonical", "body", string(m.EnvironmentCanonical), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("environment_canonical", "body", string(m.EnvironmentCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("environment_canonical", "body", string(m.EnvironmentCanonical), `^[\da-zA-Z]+(?:[\da-zA-Z\-._]+[\da-zA-Z]|[\da-zA-Z])$`); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateProjectCanonical(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectCanonical) { // not required
		return nil
	}

	if err := validate.MinLength("project_canonical", "body", string(m.ProjectCanonical), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("project_canonical", "body", string(m.ProjectCanonical), 100); err != nil {
		return err
	}

	if err := validate.Pattern("project_canonical", "body", string(m.ProjectCanonical), `^[a-z0-9]+[a-z0-9\-_]+[a-z0-9]+$`); err != nil {
		return err
	}

	return nil
}

var kPITypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["build_avg_time","build_frequency","build_history","code_coverage","time_to_release"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kPITypeTypePropEnum = append(kPITypeTypePropEnum, v)
	}
}

const (

	// KPITypeBuildAvgTime captures enum value "build_avg_time"
	KPITypeBuildAvgTime string = "build_avg_time"

	// KPITypeBuildFrequency captures enum value "build_frequency"
	KPITypeBuildFrequency string = "build_frequency"

	// KPITypeBuildHistory captures enum value "build_history"
	KPITypeBuildHistory string = "build_history"

	// KPITypeCodeCoverage captures enum value "code_coverage"
	KPITypeCodeCoverage string = "code_coverage"

	// KPITypeTimeToRelease captures enum value "time_to_release"
	KPITypeTimeToRelease string = "time_to_release"
)

// prop value enum
func (m *KPI) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, kPITypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KPI) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *KPI) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.MinimumInt("updated_at", "body", int64(*m.UpdatedAt), 0, false); err != nil {
		return err
	}

	return nil
}

var kPITypeWidgetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bars","stackbars","doughnut","history","line","pie","summary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kPITypeWidgetPropEnum = append(kPITypeWidgetPropEnum, v)
	}
}

const (

	// KPIWidgetBars captures enum value "bars"
	KPIWidgetBars string = "bars"

	// KPIWidgetStackbars captures enum value "stackbars"
	KPIWidgetStackbars string = "stackbars"

	// KPIWidgetDoughnut captures enum value "doughnut"
	KPIWidgetDoughnut string = "doughnut"

	// KPIWidgetHistory captures enum value "history"
	KPIWidgetHistory string = "history"

	// KPIWidgetLine captures enum value "line"
	KPIWidgetLine string = "line"

	// KPIWidgetPie captures enum value "pie"
	KPIWidgetPie string = "pie"

	// KPIWidgetSummary captures enum value "summary"
	KPIWidgetSummary string = "summary"
)

// prop value enum
func (m *KPI) validateWidgetEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, kPITypeWidgetPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KPI) validateWidget(formats strfmt.Registry) error {

	if err := validate.Required("widget", "body", m.Widget); err != nil {
		return err
	}

	// value enum
	if err := m.validateWidgetEnum("widget", "body", *m.Widget); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KPI) UnmarshalBinary(b []byte) error {
	var res KPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

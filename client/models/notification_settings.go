// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NotificationSettings User's notification settings
//
// swagger:model NotificationSettings
type NotificationSettings struct {

	// The frequency of the activity report.
	// Required: true
	// Enum: ["none","daily","weekly"]
	ActivityReportFrequency *string `json:"activity_report_frequency"`

	// Automatically subscribe to the project updates when creating a new project.
	// Required: true
	WatchProjectsOnCreate *bool `json:"watch_projects_on_create"`
}

// Validate validates this notification settings
func (m *NotificationSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityReportFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWatchProjectsOnCreate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var notificationSettingsTypeActivityReportFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","daily","weekly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notificationSettingsTypeActivityReportFrequencyPropEnum = append(notificationSettingsTypeActivityReportFrequencyPropEnum, v)
	}
}

const (

	// NotificationSettingsActivityReportFrequencyNone captures enum value "none"
	NotificationSettingsActivityReportFrequencyNone string = "none"

	// NotificationSettingsActivityReportFrequencyDaily captures enum value "daily"
	NotificationSettingsActivityReportFrequencyDaily string = "daily"

	// NotificationSettingsActivityReportFrequencyWeekly captures enum value "weekly"
	NotificationSettingsActivityReportFrequencyWeekly string = "weekly"
)

// prop value enum
func (m *NotificationSettings) validateActivityReportFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, notificationSettingsTypeActivityReportFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NotificationSettings) validateActivityReportFrequency(formats strfmt.Registry) error {

	if err := validate.Required("activity_report_frequency", "body", m.ActivityReportFrequency); err != nil {
		return err
	}

	// value enum
	if err := m.validateActivityReportFrequencyEnum("activity_report_frequency", "body", *m.ActivityReportFrequency); err != nil {
		return err
	}

	return nil
}

func (m *NotificationSettings) validateWatchProjectsOnCreate(formats strfmt.Registry) error {

	if err := validate.Required("watch_projects_on_create", "body", m.WatchProjectsOnCreate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this notification settings based on context it is used
func (m *NotificationSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationSettings) UnmarshalBinary(b []byte) error {
	var res NotificationSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

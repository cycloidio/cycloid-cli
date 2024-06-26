// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewEvent A new event
//
// A new event to register in the Cycloid platform.
//
// swagger:model NewEvent
type NewEvent struct {

	// The HTML color associated to the event. The valid colors are the ones listed in the  CSS 2 specification: https://www.w3.org/TR/CSS2/syndata.html#value-def-color. Only the keyword are accepted, the hexadecimal values are not valid.
	// Max Length: 20
	// Min Length: 3
	// Pattern: [a-z]+
	Color string `json:"color,omitempty"`

	// [A Font Awesome class name](https://fontawesome.com/icons)
	// Min Length: 3
	Icon string `json:"icon,omitempty"`

	// The message associated to the event.
	// Required: true
	// Min Length: 1
	Message *string `json:"message"`

	// tThe severity associated to the event.
	// Required: true
	// Enum: ["info","warn","err","crit"]
	Severity *string `json:"severity"`

	// The list of tags associated to the event.
	// Required: true
	Tags []*Tag `json:"tags"`

	// The title of the event.
	// Required: true
	// Min Length: 1
	Title *string `json:"title"`

	// The type of the event
	// Required: true
	// Enum: ["Cycloid","AWS","Monitoring","Custom"]
	Type *string `json:"type"`
}

// Validate validates this new event
func (m *NewEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewEvent) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MinLength("color", "body", m.Color, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("color", "body", m.Color, 20); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `[a-z]+`); err != nil {
		return err
	}

	return nil
}

func (m *NewEvent) validateIcon(formats strfmt.Registry) error {
	if swag.IsZero(m.Icon) { // not required
		return nil
	}

	if err := validate.MinLength("icon", "body", m.Icon, 3); err != nil {
		return err
	}

	return nil
}

func (m *NewEvent) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	if err := validate.MinLength("message", "body", *m.Message, 1); err != nil {
		return err
	}

	return nil
}

var newEventTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["info","warn","err","crit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newEventTypeSeverityPropEnum = append(newEventTypeSeverityPropEnum, v)
	}
}

const (

	// NewEventSeverityInfo captures enum value "info"
	NewEventSeverityInfo string = "info"

	// NewEventSeverityWarn captures enum value "warn"
	NewEventSeverityWarn string = "warn"

	// NewEventSeverityErr captures enum value "err"
	NewEventSeverityErr string = "err"

	// NewEventSeverityCrit captures enum value "crit"
	NewEventSeverityCrit string = "crit"
)

// prop value enum
func (m *NewEvent) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, newEventTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NewEvent) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *NewEvent) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewEvent) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	return nil
}

var newEventTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cycloid","AWS","Monitoring","Custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newEventTypeTypePropEnum = append(newEventTypeTypePropEnum, v)
	}
}

const (

	// NewEventTypeCycloid captures enum value "Cycloid"
	NewEventTypeCycloid string = "Cycloid"

	// NewEventTypeAWS captures enum value "AWS"
	NewEventTypeAWS string = "AWS"

	// NewEventTypeMonitoring captures enum value "Monitoring"
	NewEventTypeMonitoring string = "Monitoring"

	// NewEventTypeCustom captures enum value "Custom"
	NewEventTypeCustom string = "Custom"
)

// prop value enum
func (m *NewEvent) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, newEventTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NewEvent) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this new event based on the context it is used
func (m *NewEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewEvent) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewEvent) UnmarshalBinary(b []byte) error {
	var res NewEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

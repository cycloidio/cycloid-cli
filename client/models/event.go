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

// Event An event
//
// A event which has registered an activity in the Cycloid platform.
//
// swagger:model Event
type Event struct {

	// The HTML color associated to the event. The valid colors are the ones listed in the  CSS 2 specification: https://www.w3.org/TR/CSS2/syndata.html#value-def-color. Only the keyword are accepted, the hexadecimal values are not valid.
	// Max Length: 20
	// Min Length: 3
	// Pattern: [a-z]+
	Color string `json:"color,omitempty"`

	// [A Font Awesome class name](https://fontawesome.com/icons)
	// Min Length: 3
	Icon string `json:"icon,omitempty"`

	// The unique ID of the event from the database.
	// Required: true
	ID *uint32 `json:"id"`

	// The message associated to the event.
	// Required: true
	// Min Length: 1
	Message *string `json:"message"`

	// The severity associated to the event.
	// Required: true
	// Enum: ["info","warn","err","crit"]
	Severity *string `json:"severity"`

	// The list of tags associated to the event.
	// Required: true
	Tags []*Tag `json:"tags"`

	// The timestamp when the event was created in milliseconds.
	// Required: true
	Timestamp *uint64 `json:"timestamp"`

	// The title of the event.
	// Required: true
	// Min Length: 1
	Title *string `json:"title"`

	// The type of the event
	// Required: true
	// Enum: ["Cycloid","AWS","Monitoring","Custom"]
	Type *string `json:"type"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

	if err := m.validateTimestamp(formats); err != nil {
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

func (m *Event) validateColor(formats strfmt.Registry) error {
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

func (m *Event) validateIcon(formats strfmt.Registry) error {
	if swag.IsZero(m.Icon) { // not required
		return nil
	}

	if err := validate.MinLength("icon", "body", m.Icon, 3); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	if err := validate.MinLength("message", "body", *m.Message, 1); err != nil {
		return err
	}

	return nil
}

var eventTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["info","warn","err","crit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypeSeverityPropEnum = append(eventTypeSeverityPropEnum, v)
	}
}

const (

	// EventSeverityInfo captures enum value "info"
	EventSeverityInfo string = "info"

	// EventSeverityWarn captures enum value "warn"
	EventSeverityWarn string = "warn"

	// EventSeverityErr captures enum value "err"
	EventSeverityErr string = "err"

	// EventSeverityCrit captures enum value "crit"
	EventSeverityCrit string = "crit"
)

// prop value enum
func (m *Event) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Event) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTags(formats strfmt.Registry) error {

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

func (m *Event) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	return nil
}

var eventTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cycloid","AWS","Monitoring","Custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypeTypePropEnum = append(eventTypeTypePropEnum, v)
	}
}

const (

	// EventTypeCycloid captures enum value "Cycloid"
	EventTypeCycloid string = "Cycloid"

	// EventTypeAWS captures enum value "AWS"
	EventTypeAWS string = "AWS"

	// EventTypeMonitoring captures enum value "Monitoring"
	EventTypeMonitoring string = "Monitoring"

	// EventTypeCustom captures enum value "Custom"
	EventTypeCustom string = "Custom"
)

// prop value enum
func (m *Event) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Event) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this event based on the context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

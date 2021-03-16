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

// FormEntity Forms file's Entity
//
// This describes all the attributes possible to configure a form's variable
// swagger:model FormEntity
type FormEntity struct {

	// The current value that was previously configured for this variable upon creation or update.
	Current interface{} `json:"current,omitempty"`

	// The default to assign to the variable if nothing is returned and that the varialbe is required
	Default interface{} `json:"default,omitempty"`

	// The description helping users understand the interest/impact of such variable/change
	Description string `json:"description,omitempty"`

	// The key is the name of variables for the ansible/pipeline/terraform technologies. If this is a first level variable then: keyX. If you have multiple terraform modules then use: module.Y.keyX to help identify the unique variable.
	// Required: true
	Key *string `json:"key"`

	// The name of the variable displayed to the user
	// Required: true
	Name *string `json:"name"`

	// Whether or not the field is required - that helps distinguish "optional" variables and allows to set default if necessary and present
	Required bool `json:"required,omitempty"`

	// The source is only used for the branch widget to reference the key of the Catalog Repository or Config Repository that the branches have to be read from. Because a branch in itself cannot exist, the user has to indicate from which Catalog Repository or Config Repository he wants to retrieve branches. The source has to reference the key of an entity of a widget: 'CyCRS' or 'CySCS'
	Source string `json:"source,omitempty"`

	// The type of data handled - used to manipulate/validate the input, and also validate default/values
	// Required: true
	// Enum: [integer string array boolean map]
	Type *string `json:"type"`

	// The unit to be displayed for the variable, helping to know what's being manipulated: amount of servers, Go, users, etc.
	Unit string `json:"unit,omitempty"`

	// Values allowed, e.g. [1, 10, 20, 50], this can be of any type but boolean. Note: In case of SliderRange only 2 values should be provided: [min, max], in case of providing them the other way around some validation test will fail.
	Values []interface{} `json:"values"`

	// The widget used to display the data in the most suitable way
	// Required: true
	// Enum: [auto_complete dropdown radios slider_list slider_range number simple_text switch text_area cy_cred cy_scs cy_crs cy_branch]
	Widget *string `json:"widget"`
}

// Validate validates this form entity
func (m *FormEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *FormEntity) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *FormEntity) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var formEntityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["integer","string","array","boolean","map"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		formEntityTypeTypePropEnum = append(formEntityTypeTypePropEnum, v)
	}
}

const (

	// FormEntityTypeInteger captures enum value "integer"
	FormEntityTypeInteger string = "integer"

	// FormEntityTypeString captures enum value "string"
	FormEntityTypeString string = "string"

	// FormEntityTypeArray captures enum value "array"
	FormEntityTypeArray string = "array"

	// FormEntityTypeBoolean captures enum value "boolean"
	FormEntityTypeBoolean string = "boolean"

	// FormEntityTypeMap captures enum value "map"
	FormEntityTypeMap string = "map"
)

// prop value enum
func (m *FormEntity) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, formEntityTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FormEntity) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var formEntityTypeWidgetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto_complete","dropdown","radios","slider_list","slider_range","number","simple_text","switch","text_area","cy_cred","cy_scs","cy_crs","cy_branch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		formEntityTypeWidgetPropEnum = append(formEntityTypeWidgetPropEnum, v)
	}
}

const (

	// FormEntityWidgetAutoComplete captures enum value "auto_complete"
	FormEntityWidgetAutoComplete string = "auto_complete"

	// FormEntityWidgetDropdown captures enum value "dropdown"
	FormEntityWidgetDropdown string = "dropdown"

	// FormEntityWidgetRadios captures enum value "radios"
	FormEntityWidgetRadios string = "radios"

	// FormEntityWidgetSliderList captures enum value "slider_list"
	FormEntityWidgetSliderList string = "slider_list"

	// FormEntityWidgetSliderRange captures enum value "slider_range"
	FormEntityWidgetSliderRange string = "slider_range"

	// FormEntityWidgetNumber captures enum value "number"
	FormEntityWidgetNumber string = "number"

	// FormEntityWidgetSimpleText captures enum value "simple_text"
	FormEntityWidgetSimpleText string = "simple_text"

	// FormEntityWidgetSwitch captures enum value "switch"
	FormEntityWidgetSwitch string = "switch"

	// FormEntityWidgetTextArea captures enum value "text_area"
	FormEntityWidgetTextArea string = "text_area"

	// FormEntityWidgetCyCred captures enum value "cy_cred"
	FormEntityWidgetCyCred string = "cy_cred"

	// FormEntityWidgetCyScs captures enum value "cy_scs"
	FormEntityWidgetCyScs string = "cy_scs"

	// FormEntityWidgetCyCrs captures enum value "cy_crs"
	FormEntityWidgetCyCrs string = "cy_crs"

	// FormEntityWidgetCyBranch captures enum value "cy_branch"
	FormEntityWidgetCyBranch string = "cy_branch"
)

// prop value enum
func (m *FormEntity) validateWidgetEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, formEntityTypeWidgetPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FormEntity) validateWidget(formats strfmt.Registry) error {

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
func (m *FormEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormEntity) UnmarshalBinary(b []byte) error {
	var res FormEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

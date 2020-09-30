// Code generated by go-swagger; DO NOT EDIT.

package organization_config_repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// CreateConfigRepositoryReader is a Reader for the CreateConfigRepository structure.
type CreateConfigRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConfigRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateConfigRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateConfigRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewCreateConfigRepositoryLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateConfigRepositoryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateConfigRepositoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateConfigRepositoryOK creates a CreateConfigRepositoryOK with default headers values
func NewCreateConfigRepositoryOK() *CreateConfigRepositoryOK {
	return &CreateConfigRepositoryOK{}
}

/*CreateConfigRepositoryOK handles this case with default header values.

Success creation
*/
type CreateConfigRepositoryOK struct {
	Payload *CreateConfigRepositoryOKBody
}

func (o *CreateConfigRepositoryOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories][%d] createConfigRepositoryOK  %+v", 200, o.Payload)
}

func (o *CreateConfigRepositoryOK) GetPayload() *CreateConfigRepositoryOKBody {
	return o.Payload
}

func (o *CreateConfigRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateConfigRepositoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfigRepositoryNotFound creates a CreateConfigRepositoryNotFound with default headers values
func NewCreateConfigRepositoryNotFound() *CreateConfigRepositoryNotFound {
	return &CreateConfigRepositoryNotFound{}
}

/*CreateConfigRepositoryNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateConfigRepositoryNotFound struct {
	Payload *models.ErrorPayload
}

func (o *CreateConfigRepositoryNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories][%d] createConfigRepositoryNotFound  %+v", 404, o.Payload)
}

func (o *CreateConfigRepositoryNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfigRepositoryLengthRequired creates a CreateConfigRepositoryLengthRequired with default headers values
func NewCreateConfigRepositoryLengthRequired() *CreateConfigRepositoryLengthRequired {
	return &CreateConfigRepositoryLengthRequired{}
}

/*CreateConfigRepositoryLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreateConfigRepositoryLengthRequired struct {
}

func (o *CreateConfigRepositoryLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories][%d] createConfigRepositoryLengthRequired ", 411)
}

func (o *CreateConfigRepositoryLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConfigRepositoryUnprocessableEntity creates a CreateConfigRepositoryUnprocessableEntity with default headers values
func NewCreateConfigRepositoryUnprocessableEntity() *CreateConfigRepositoryUnprocessableEntity {
	return &CreateConfigRepositoryUnprocessableEntity{}
}

/*CreateConfigRepositoryUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateConfigRepositoryUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *CreateConfigRepositoryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories][%d] createConfigRepositoryUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateConfigRepositoryUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfigRepositoryDefault creates a CreateConfigRepositoryDefault with default headers values
func NewCreateConfigRepositoryDefault(code int) *CreateConfigRepositoryDefault {
	return &CreateConfigRepositoryDefault{
		_statusCode: code,
	}
}

/*CreateConfigRepositoryDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateConfigRepositoryDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the create config repository default response
func (o *CreateConfigRepositoryDefault) Code() int {
	return o._statusCode
}

func (o *CreateConfigRepositoryDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/config_repositories][%d] createConfigRepository default  %+v", o._statusCode, o.Payload)
}

func (o *CreateConfigRepositoryDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateConfigRepositoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateConfigRepositoryOKBody create config repository o k body
swagger:model CreateConfigRepositoryOKBody
*/
type CreateConfigRepositoryOKBody struct {

	// data
	// Required: true
	Data *models.ConfigRepository `json:"data"`
}

// Validate validates this create config repository o k body
func (o *CreateConfigRepositoryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateConfigRepositoryOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createConfigRepositoryOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createConfigRepositoryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateConfigRepositoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateConfigRepositoryOKBody) UnmarshalBinary(b []byte) error {
	var res CreateConfigRepositoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

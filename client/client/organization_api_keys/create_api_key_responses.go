// Code generated by go-swagger; DO NOT EDIT.

package organization_api_keys

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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// CreateAPIKeyReader is a Reader for the CreateAPIKey structure.
type CreateAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewCreateAPIKeyLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateAPIKeyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAPIKeyOK creates a CreateAPIKeyOK with default headers values
func NewCreateAPIKeyOK() *CreateAPIKeyOK {
	return &CreateAPIKeyOK{}
}

/*CreateAPIKeyOK handles this case with default header values.

API key created. The body contains the information of the new API key of the organization.
*/
type CreateAPIKeyOK struct {
	Payload *CreateAPIKeyOKBody
}

func (o *CreateAPIKeyOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/api_keys][%d] createApiKeyOK  %+v", 200, o.Payload)
}

func (o *CreateAPIKeyOK) GetPayload() *CreateAPIKeyOKBody {
	return o.Payload
}

func (o *CreateAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateAPIKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIKeyNotFound creates a CreateAPIKeyNotFound with default headers values
func NewCreateAPIKeyNotFound() *CreateAPIKeyNotFound {
	return &CreateAPIKeyNotFound{}
}

/*CreateAPIKeyNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateAPIKeyNotFound struct {
	Payload *models.ErrorPayload
}

func (o *CreateAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/api_keys][%d] createApiKeyNotFound  %+v", 404, o.Payload)
}

func (o *CreateAPIKeyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIKeyLengthRequired creates a CreateAPIKeyLengthRequired with default headers values
func NewCreateAPIKeyLengthRequired() *CreateAPIKeyLengthRequired {
	return &CreateAPIKeyLengthRequired{}
}

/*CreateAPIKeyLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreateAPIKeyLengthRequired struct {
}

func (o *CreateAPIKeyLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/api_keys][%d] createApiKeyLengthRequired ", 411)
}

func (o *CreateAPIKeyLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAPIKeyUnprocessableEntity creates a CreateAPIKeyUnprocessableEntity with default headers values
func NewCreateAPIKeyUnprocessableEntity() *CreateAPIKeyUnprocessableEntity {
	return &CreateAPIKeyUnprocessableEntity{}
}

/*CreateAPIKeyUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateAPIKeyUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *CreateAPIKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/api_keys][%d] createApiKeyUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateAPIKeyUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateAPIKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIKeyDefault creates a CreateAPIKeyDefault with default headers values
func NewCreateAPIKeyDefault(code int) *CreateAPIKeyDefault {
	return &CreateAPIKeyDefault{
		_statusCode: code,
	}
}

/*CreateAPIKeyDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateAPIKeyDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the create API key default response
func (o *CreateAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *CreateAPIKeyDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/api_keys][%d] createAPIKey default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAPIKeyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateAPIKeyOKBody create API key o k body
swagger:model CreateAPIKeyOKBody
*/
type CreateAPIKeyOKBody struct {

	// data
	// Required: true
	Data *models.APIKey `json:"data"`
}

// Validate validates this create API key o k body
func (o *CreateAPIKeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAPIKeyOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createApiKeyOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createApiKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateAPIKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAPIKeyOKBody) UnmarshalBinary(b []byte) error {
	var res CreateAPIKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

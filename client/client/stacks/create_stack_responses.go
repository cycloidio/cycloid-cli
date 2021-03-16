// Code generated by go-swagger; DO NOT EDIT.

package stacks

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

// CreateStackReader is a Reader for the CreateStack structure.
type CreateStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateStackAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateStackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateStackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateStackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateStackAccepted creates a CreateStackAccepted with default headers values
func NewCreateStackAccepted() *CreateStackAccepted {
	return &CreateStackAccepted{}
}

/*CreateStackAccepted handles this case with default header values.

The information of the stack.
*/
type CreateStackAccepted struct {
	Payload *CreateStackAcceptedBody
}

func (o *CreateStackAccepted) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/stacks][%d] createStackAccepted  %+v", 202, o.Payload)
}

func (o *CreateStackAccepted) GetPayload() *CreateStackAcceptedBody {
	return o.Payload
}

func (o *CreateStackAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateStackAcceptedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStackForbidden creates a CreateStackForbidden with default headers values
func NewCreateStackForbidden() *CreateStackForbidden {
	return &CreateStackForbidden{}
}

/*CreateStackForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateStackForbidden struct {
	Payload *models.ErrorPayload
}

func (o *CreateStackForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/stacks][%d] createStackForbidden  %+v", 403, o.Payload)
}

func (o *CreateStackForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateStackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStackNotFound creates a CreateStackNotFound with default headers values
func NewCreateStackNotFound() *CreateStackNotFound {
	return &CreateStackNotFound{}
}

/*CreateStackNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateStackNotFound struct {
	Payload *models.ErrorPayload
}

func (o *CreateStackNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/stacks][%d] createStackNotFound  %+v", 404, o.Payload)
}

func (o *CreateStackNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStackUnprocessableEntity creates a CreateStackUnprocessableEntity with default headers values
func NewCreateStackUnprocessableEntity() *CreateStackUnprocessableEntity {
	return &CreateStackUnprocessableEntity{}
}

/*CreateStackUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateStackUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *CreateStackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/stacks][%d] createStackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateStackUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateStackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStackDefault creates a CreateStackDefault with default headers values
func NewCreateStackDefault(code int) *CreateStackDefault {
	return &CreateStackDefault{
		_statusCode: code,
	}
}

/*CreateStackDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateStackDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the create stack default response
func (o *CreateStackDefault) Code() int {
	return o._statusCode
}

func (o *CreateStackDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/stacks][%d] createStack default  %+v", o._statusCode, o.Payload)
}

func (o *CreateStackDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateStackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateStackAcceptedBody create stack accepted body
swagger:model CreateStackAcceptedBody
*/
type CreateStackAcceptedBody struct {

	// data
	// Required: true
	Data *models.Stack `json:"data"`
}

// Validate validates this create stack accepted body
func (o *CreateStackAcceptedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateStackAcceptedBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createStackAccepted"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createStackAccepted" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateStackAcceptedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateStackAcceptedBody) UnmarshalBinary(b []byte) error {
	var res CreateStackAcceptedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

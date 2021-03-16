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

// UpdateStackReader is a Reader for the UpdateStack structure.
type UpdateStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateStackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateStackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateStackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateStackOK creates a UpdateStackOK with default headers values
func NewUpdateStackOK() *UpdateStackOK {
	return &UpdateStackOK{}
}

/*UpdateStackOK handles this case with default header values.

Updated the Stack
*/
type UpdateStackOK struct {
	Payload *UpdateStackOKBody
}

func (o *UpdateStackOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/stacks/{stack_ref}][%d] updateStackOK  %+v", 200, o.Payload)
}

func (o *UpdateStackOK) GetPayload() *UpdateStackOKBody {
	return o.Payload
}

func (o *UpdateStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateStackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStackForbidden creates a UpdateStackForbidden with default headers values
func NewUpdateStackForbidden() *UpdateStackForbidden {
	return &UpdateStackForbidden{}
}

/*UpdateStackForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateStackForbidden struct {
	Payload *models.ErrorPayload
}

func (o *UpdateStackForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/stacks/{stack_ref}][%d] updateStackForbidden  %+v", 403, o.Payload)
}

func (o *UpdateStackForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateStackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStackNotFound creates a UpdateStackNotFound with default headers values
func NewUpdateStackNotFound() *UpdateStackNotFound {
	return &UpdateStackNotFound{}
}

/*UpdateStackNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateStackNotFound struct {
	Payload *models.ErrorPayload
}

func (o *UpdateStackNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/stacks/{stack_ref}][%d] updateStackNotFound  %+v", 404, o.Payload)
}

func (o *UpdateStackNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStackUnprocessableEntity creates a UpdateStackUnprocessableEntity with default headers values
func NewUpdateStackUnprocessableEntity() *UpdateStackUnprocessableEntity {
	return &UpdateStackUnprocessableEntity{}
}

/*UpdateStackUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateStackUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *UpdateStackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/stacks/{stack_ref}][%d] updateStackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateStackUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateStackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateStackDefault creates a UpdateStackDefault with default headers values
func NewUpdateStackDefault(code int) *UpdateStackDefault {
	return &UpdateStackDefault{
		_statusCode: code,
	}
}

/*UpdateStackDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateStackDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the update stack default response
func (o *UpdateStackDefault) Code() int {
	return o._statusCode
}

func (o *UpdateStackDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/stacks/{stack_ref}][%d] updateStack default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateStackDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateStackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateStackOKBody update stack o k body
swagger:model UpdateStackOKBody
*/
type UpdateStackOKBody struct {

	// data
	// Required: true
	Data *models.Stack `json:"data"`
}

// Validate validates this update stack o k body
func (o *UpdateStackOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateStackOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateStackOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStackOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateStackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateStackOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateStackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

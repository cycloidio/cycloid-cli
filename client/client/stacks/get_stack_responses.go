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

// GetStackReader is a Reader for the GetStack structure.
type GetStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetStackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStackOK creates a GetStackOK with default headers values
func NewGetStackOK() *GetStackOK {
	return &GetStackOK{}
}

/*GetStackOK handles this case with default header values.

The information of the stack.
*/
type GetStackOK struct {
	Payload *GetStackOKBody
}

func (o *GetStackOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}][%d] getStackOK  %+v", 200, o.Payload)
}

func (o *GetStackOK) GetPayload() *GetStackOKBody {
	return o.Payload
}

func (o *GetStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackForbidden creates a GetStackForbidden with default headers values
func NewGetStackForbidden() *GetStackForbidden {
	return &GetStackForbidden{}
}

/*GetStackForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetStackForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetStackForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}][%d] getStackForbidden  %+v", 403, o.Payload)
}

func (o *GetStackForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackNotFound creates a GetStackNotFound with default headers values
func NewGetStackNotFound() *GetStackNotFound {
	return &GetStackNotFound{}
}

/*GetStackNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetStackNotFound struct {
	Payload *models.ErrorPayload
}

func (o *GetStackNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}][%d] getStackNotFound  %+v", 404, o.Payload)
}

func (o *GetStackNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStackDefault creates a GetStackDefault with default headers values
func NewGetStackDefault(code int) *GetStackDefault {
	return &GetStackDefault{
		_statusCode: code,
	}
}

/*GetStackDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetStackDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get stack default response
func (o *GetStackDefault) Code() int {
	return o._statusCode
}

func (o *GetStackDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/stacks/{stack_ref}][%d] getStack default  %+v", o._statusCode, o.Payload)
}

func (o *GetStackDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetStackOKBody get stack o k body
swagger:model GetStackOKBody
*/
type GetStackOKBody struct {

	// data
	// Required: true
	Data *models.Stack `json:"data"`
}

// Validate validates this get stack o k body
func (o *GetStackOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStackOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getStackOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStackOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStackOKBody) UnmarshalBinary(b []byte) error {
	var res GetStackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

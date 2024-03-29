// Code generated by go-swagger; DO NOT EDIT.

package user

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

// UpdateUserAccountReader is a Reader for the UpdateUserAccount structure.
type UpdateUserAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewUpdateUserAccountConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewUpdateUserAccountLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateUserAccountUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateUserAccountServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateUserAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserAccountOK creates a UpdateUserAccountOK with default headers values
func NewUpdateUserAccountOK() *UpdateUserAccountOK {
	return &UpdateUserAccountOK{}
}

/*UpdateUserAccountOK handles this case with default header values.

The updated user profile information.
*/
type UpdateUserAccountOK struct {
	Payload *UpdateUserAccountOKBody
}

func (o *UpdateUserAccountOK) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateUserAccountOK) GetPayload() *UpdateUserAccountOKBody {
	return o.Payload
}

func (o *UpdateUserAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateUserAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAccountConflict creates a UpdateUserAccountConflict with default headers values
func NewUpdateUserAccountConflict() *UpdateUserAccountConflict {
	return &UpdateUserAccountConflict{}
}

/*UpdateUserAccountConflict handles this case with default header values.

Trying setting an unverified email as the primary
*/
type UpdateUserAccountConflict struct {
}

func (o *UpdateUserAccountConflict) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountConflict ", 409)
}

func (o *UpdateUserAccountConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserAccountLengthRequired creates a UpdateUserAccountLengthRequired with default headers values
func NewUpdateUserAccountLengthRequired() *UpdateUserAccountLengthRequired {
	return &UpdateUserAccountLengthRequired{}
}

/*UpdateUserAccountLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type UpdateUserAccountLengthRequired struct {
}

func (o *UpdateUserAccountLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountLengthRequired ", 411)
}

func (o *UpdateUserAccountLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserAccountUnprocessableEntity creates a UpdateUserAccountUnprocessableEntity with default headers values
func NewUpdateUserAccountUnprocessableEntity() *UpdateUserAccountUnprocessableEntity {
	return &UpdateUserAccountUnprocessableEntity{}
}

/*UpdateUserAccountUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateUserAccountUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateUserAccountUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateUserAccountUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserAccountUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAccountServiceUnavailable creates a UpdateUserAccountServiceUnavailable with default headers values
func NewUpdateUserAccountServiceUnavailable() *UpdateUserAccountServiceUnavailable {
	return &UpdateUserAccountServiceUnavailable{}
}

/*UpdateUserAccountServiceUnavailable handles this case with default header values.

The operation couldn't be executed or completed and it should retried.
*/
type UpdateUserAccountServiceUnavailable struct {
	/*The number of seconds to wait until retry the request
	 */
	RetryAfter uint16

	Payload *models.ErrorPayload
}

func (o *UpdateUserAccountServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateUserAccountServiceUnavailable) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserAccountServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Retry-After
	retryAfter, err := swag.ConvertUint16(response.GetHeader("Retry-After"))
	if err != nil {
		return errors.InvalidType("Retry-After", "header", "uint16", response.GetHeader("Retry-After"))
	}
	o.RetryAfter = retryAfter

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserAccountDefault creates a UpdateUserAccountDefault with default headers values
func NewUpdateUserAccountDefault(code int) *UpdateUserAccountDefault {
	return &UpdateUserAccountDefault{
		_statusCode: code,
	}
}

/*UpdateUserAccountDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateUserAccountDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the update user account default response
func (o *UpdateUserAccountDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserAccountDefault) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccount default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateUserAccountDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateUserAccountOKBody update user account o k body
swagger:model UpdateUserAccountOKBody
*/
type UpdateUserAccountOKBody struct {

	// data
	// Required: true
	Data *models.UserAccount `json:"data"`
}

// Validate validates this update user account o k body
func (o *UpdateUserAccountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserAccountOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateUserAccountOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateUserAccountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserAccountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserAccountOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserAccountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

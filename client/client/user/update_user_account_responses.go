// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cycloidio/cycloid-cli/client/models"
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

/*
UpdateUserAccountOK describes a response with status code 200, with default header values.

The updated user profile information.
*/
type UpdateUserAccountOK struct {
	Payload *UpdateUserAccountOKBody
}

// IsSuccess returns true when this update user account o k response has a 2xx status code
func (o *UpdateUserAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user account o k response has a 3xx status code
func (o *UpdateUserAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user account o k response has a 4xx status code
func (o *UpdateUserAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user account o k response has a 5xx status code
func (o *UpdateUserAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user account o k response a status code equal to that given
func (o *UpdateUserAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update user account o k response
func (o *UpdateUserAccountOK) Code() int {
	return 200
}

func (o *UpdateUserAccountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountOK %s", 200, payload)
}

func (o *UpdateUserAccountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountOK %s", 200, payload)
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

/*
UpdateUserAccountConflict describes a response with status code 409, with default header values.

Trying setting an unverified email as the primary
*/
type UpdateUserAccountConflict struct {
}

// IsSuccess returns true when this update user account conflict response has a 2xx status code
func (o *UpdateUserAccountConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user account conflict response has a 3xx status code
func (o *UpdateUserAccountConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user account conflict response has a 4xx status code
func (o *UpdateUserAccountConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user account conflict response has a 5xx status code
func (o *UpdateUserAccountConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update user account conflict response a status code equal to that given
func (o *UpdateUserAccountConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update user account conflict response
func (o *UpdateUserAccountConflict) Code() int {
	return 409
}

func (o *UpdateUserAccountConflict) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountConflict", 409)
}

func (o *UpdateUserAccountConflict) String() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountConflict", 409)
}

func (o *UpdateUserAccountConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserAccountLengthRequired creates a UpdateUserAccountLengthRequired with default headers values
func NewUpdateUserAccountLengthRequired() *UpdateUserAccountLengthRequired {
	return &UpdateUserAccountLengthRequired{}
}

/*
UpdateUserAccountLengthRequired describes a response with status code 411, with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type UpdateUserAccountLengthRequired struct {
}

// IsSuccess returns true when this update user account length required response has a 2xx status code
func (o *UpdateUserAccountLengthRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user account length required response has a 3xx status code
func (o *UpdateUserAccountLengthRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user account length required response has a 4xx status code
func (o *UpdateUserAccountLengthRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user account length required response has a 5xx status code
func (o *UpdateUserAccountLengthRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this update user account length required response a status code equal to that given
func (o *UpdateUserAccountLengthRequired) IsCode(code int) bool {
	return code == 411
}

// Code gets the status code for the update user account length required response
func (o *UpdateUserAccountLengthRequired) Code() int {
	return 411
}

func (o *UpdateUserAccountLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountLengthRequired", 411)
}

func (o *UpdateUserAccountLengthRequired) String() string {
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountLengthRequired", 411)
}

func (o *UpdateUserAccountLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserAccountUnprocessableEntity creates a UpdateUserAccountUnprocessableEntity with default headers values
func NewUpdateUserAccountUnprocessableEntity() *UpdateUserAccountUnprocessableEntity {
	return &UpdateUserAccountUnprocessableEntity{}
}

/*
UpdateUserAccountUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateUserAccountUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update user account unprocessable entity response has a 2xx status code
func (o *UpdateUserAccountUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user account unprocessable entity response has a 3xx status code
func (o *UpdateUserAccountUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user account unprocessable entity response has a 4xx status code
func (o *UpdateUserAccountUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user account unprocessable entity response has a 5xx status code
func (o *UpdateUserAccountUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update user account unprocessable entity response a status code equal to that given
func (o *UpdateUserAccountUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update user account unprocessable entity response
func (o *UpdateUserAccountUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateUserAccountUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountUnprocessableEntity %s", 422, payload)
}

func (o *UpdateUserAccountUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountUnprocessableEntity %s", 422, payload)
}

func (o *UpdateUserAccountUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserAccountUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

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

/*
UpdateUserAccountServiceUnavailable describes a response with status code 503, with default header values.

The operation couldn't be executed or completed and it should retried.
*/
type UpdateUserAccountServiceUnavailable struct {

	/* The number of seconds to wait until retry the request

	   Format: uint16
	*/
	RetryAfter uint16

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update user account service unavailable response has a 2xx status code
func (o *UpdateUserAccountServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user account service unavailable response has a 3xx status code
func (o *UpdateUserAccountServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user account service unavailable response has a 4xx status code
func (o *UpdateUserAccountServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user account service unavailable response has a 5xx status code
func (o *UpdateUserAccountServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update user account service unavailable response a status code equal to that given
func (o *UpdateUserAccountServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the update user account service unavailable response
func (o *UpdateUserAccountServiceUnavailable) Code() int {
	return 503
}

func (o *UpdateUserAccountServiceUnavailable) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountServiceUnavailable %s", 503, payload)
}

func (o *UpdateUserAccountServiceUnavailable) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccountServiceUnavailable %s", 503, payload)
}

func (o *UpdateUserAccountServiceUnavailable) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserAccountServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		valretryAfter, err := swag.ConvertUint16(hdrRetryAfter)
		if err != nil {
			return errors.InvalidType("Retry-After", "header", "uint16", hdrRetryAfter)
		}
		o.RetryAfter = valretryAfter
	}

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

/*
UpdateUserAccountDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateUserAccountDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update user account default response has a 2xx status code
func (o *UpdateUserAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update user account default response has a 3xx status code
func (o *UpdateUserAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update user account default response has a 4xx status code
func (o *UpdateUserAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update user account default response has a 5xx status code
func (o *UpdateUserAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update user account default response a status code equal to that given
func (o *UpdateUserAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update user account default response
func (o *UpdateUserAccountDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserAccountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccount default %s", o._statusCode, payload)
}

func (o *UpdateUserAccountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user][%d] updateUserAccount default %s", o._statusCode, payload)
}

func (o *UpdateUserAccountDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateUserAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateUserAccountOKBody update user account o k body
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateUserAccountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update user account o k body based on the context it is used
func (o *UpdateUserAccountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserAccountOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateUserAccountOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateUserAccountOK" + "." + "data")
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

// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// EmailVerificationReader is a Reader for the EmailVerification structure.
type EmailVerificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailVerificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEmailVerificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEmailVerificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewEmailVerificationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEmailVerificationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailVerificationNoContent creates a EmailVerificationNoContent with default headers values
func NewEmailVerificationNoContent() *EmailVerificationNoContent {
	return &EmailVerificationNoContent{}
}

/*
EmailVerificationNoContent describes a response with status code 204, with default header values.

Email address has been verified.
*/
type EmailVerificationNoContent struct {
}

// IsSuccess returns true when this email verification no content response has a 2xx status code
func (o *EmailVerificationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this email verification no content response has a 3xx status code
func (o *EmailVerificationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this email verification no content response has a 4xx status code
func (o *EmailVerificationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this email verification no content response has a 5xx status code
func (o *EmailVerificationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this email verification no content response a status code equal to that given
func (o *EmailVerificationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the email verification no content response
func (o *EmailVerificationNoContent) Code() int {
	return 204
}

func (o *EmailVerificationNoContent) Error() string {
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerificationNoContent", 204)
}

func (o *EmailVerificationNoContent) String() string {
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerificationNoContent", 204)
}

func (o *EmailVerificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEmailVerificationNotFound creates a EmailVerificationNotFound with default headers values
func NewEmailVerificationNotFound() *EmailVerificationNotFound {
	return &EmailVerificationNotFound{}
}

/*
EmailVerificationNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type EmailVerificationNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this email verification not found response has a 2xx status code
func (o *EmailVerificationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this email verification not found response has a 3xx status code
func (o *EmailVerificationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this email verification not found response has a 4xx status code
func (o *EmailVerificationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this email verification not found response has a 5xx status code
func (o *EmailVerificationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this email verification not found response a status code equal to that given
func (o *EmailVerificationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the email verification not found response
func (o *EmailVerificationNotFound) Code() int {
	return 404
}

func (o *EmailVerificationNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerificationNotFound %s", 404, payload)
}

func (o *EmailVerificationNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerificationNotFound %s", 404, payload)
}

func (o *EmailVerificationNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *EmailVerificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEmailVerificationUnprocessableEntity creates a EmailVerificationUnprocessableEntity with default headers values
func NewEmailVerificationUnprocessableEntity() *EmailVerificationUnprocessableEntity {
	return &EmailVerificationUnprocessableEntity{}
}

/*
EmailVerificationUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type EmailVerificationUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this email verification unprocessable entity response has a 2xx status code
func (o *EmailVerificationUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this email verification unprocessable entity response has a 3xx status code
func (o *EmailVerificationUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this email verification unprocessable entity response has a 4xx status code
func (o *EmailVerificationUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this email verification unprocessable entity response has a 5xx status code
func (o *EmailVerificationUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this email verification unprocessable entity response a status code equal to that given
func (o *EmailVerificationUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the email verification unprocessable entity response
func (o *EmailVerificationUnprocessableEntity) Code() int {
	return 422
}

func (o *EmailVerificationUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerificationUnprocessableEntity %s", 422, payload)
}

func (o *EmailVerificationUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerificationUnprocessableEntity %s", 422, payload)
}

func (o *EmailVerificationUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *EmailVerificationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEmailVerificationDefault creates a EmailVerificationDefault with default headers values
func NewEmailVerificationDefault(code int) *EmailVerificationDefault {
	return &EmailVerificationDefault{
		_statusCode: code,
	}
}

/*
EmailVerificationDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type EmailVerificationDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this email verification default response has a 2xx status code
func (o *EmailVerificationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this email verification default response has a 3xx status code
func (o *EmailVerificationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this email verification default response has a 4xx status code
func (o *EmailVerificationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this email verification default response has a 5xx status code
func (o *EmailVerificationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this email verification default response a status code equal to that given
func (o *EmailVerificationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the email verification default response
func (o *EmailVerificationDefault) Code() int {
	return o._statusCode
}

func (o *EmailVerificationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerification default %s", o._statusCode, payload)
}

func (o *EmailVerificationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /user/email/verification/{verification_token}][%d] emailVerification default %s", o._statusCode, payload)
}

func (o *EmailVerificationDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *EmailVerificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

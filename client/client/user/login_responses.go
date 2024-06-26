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

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewLoginUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*
LoginOK describes a response with status code 200, with default header values.

The token which represents the session of the user.
*/
type LoginOK struct {
	Payload *LoginOKBody
}

// IsSuccess returns true when this login o k response has a 2xx status code
func (o *LoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this login o k response has a 3xx status code
func (o *LoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login o k response has a 4xx status code
func (o *LoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this login o k response has a 5xx status code
func (o *LoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this login o k response a status code equal to that given
func (o *LoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the login o k response
func (o *LoginOK) Code() int {
	return 200
}

func (o *LoginOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] loginOK %s", 200, payload)
}

func (o *LoginOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] loginOK %s", 200, payload)
}

func (o *LoginOK) GetPayload() *LoginOKBody {
	return o.Payload
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LoginOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUnauthorized creates a LoginUnauthorized with default headers values
func NewLoginUnauthorized() *LoginUnauthorized {
	return &LoginUnauthorized{}
}

/*
LoginUnauthorized describes a response with status code 401, with default header values.

The user cannot be authenticated with the credentials which she/he has used.
*/
type LoginUnauthorized struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this login unauthorized response has a 2xx status code
func (o *LoginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this login unauthorized response has a 3xx status code
func (o *LoginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login unauthorized response has a 4xx status code
func (o *LoginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this login unauthorized response has a 5xx status code
func (o *LoginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this login unauthorized response a status code equal to that given
func (o *LoginUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the login unauthorized response
func (o *LoginUnauthorized) Code() int {
	return 401
}

func (o *LoginUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] loginUnauthorized %s", 401, payload)
}

func (o *LoginUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] loginUnauthorized %s", 401, payload)
}

func (o *LoginUnauthorized) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *LoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLoginUnprocessableEntity creates a LoginUnprocessableEntity with default headers values
func NewLoginUnprocessableEntity() *LoginUnprocessableEntity {
	return &LoginUnprocessableEntity{}
}

/*
LoginUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type LoginUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this login unprocessable entity response has a 2xx status code
func (o *LoginUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this login unprocessable entity response has a 3xx status code
func (o *LoginUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login unprocessable entity response has a 4xx status code
func (o *LoginUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this login unprocessable entity response has a 5xx status code
func (o *LoginUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this login unprocessable entity response a status code equal to that given
func (o *LoginUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the login unprocessable entity response
func (o *LoginUnprocessableEntity) Code() int {
	return 422
}

func (o *LoginUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] loginUnprocessableEntity %s", 422, payload)
}

func (o *LoginUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] loginUnprocessableEntity %s", 422, payload)
}

func (o *LoginUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *LoginUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLoginDefault creates a LoginDefault with default headers values
func NewLoginDefault(code int) *LoginDefault {
	return &LoginDefault{
		_statusCode: code,
	}
}

/*
LoginDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type LoginDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this login default response has a 2xx status code
func (o *LoginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this login default response has a 3xx status code
func (o *LoginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this login default response has a 4xx status code
func (o *LoginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this login default response has a 5xx status code
func (o *LoginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this login default response a status code equal to that given
func (o *LoginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the login default response
func (o *LoginDefault) Code() int {
	return o._statusCode
}

func (o *LoginDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] login default %s", o._statusCode, payload)
}

func (o *LoginDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/login][%d] login default %s", o._statusCode, payload)
}

func (o *LoginDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *LoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
LoginOKBody login o k body
swagger:model LoginOKBody
*/
type LoginOKBody struct {

	// data
	// Required: true
	Data *models.UserSession `json:"data"`
}

// Validate validates this login o k body
func (o *LoginOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("loginOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loginOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loginOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this login o k body based on the context it is used
func (o *LoginOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loginOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loginOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBody) UnmarshalBinary(b []byte) error {
	var res LoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

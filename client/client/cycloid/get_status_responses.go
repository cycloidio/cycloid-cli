// Code generated by go-swagger; DO NOT EDIT.

package cycloid

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

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetStatusUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/*
GetStatusOK describes a response with status code 200, with default header values.

General application status and services statuses.
*/
type GetStatusOK struct {
	Payload *GetStatusOKBody
}

// IsSuccess returns true when this get status o k response has a 2xx status code
func (o *GetStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get status o k response has a 3xx status code
func (o *GetStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get status o k response has a 4xx status code
func (o *GetStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get status o k response has a 5xx status code
func (o *GetStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get status o k response a status code equal to that given
func (o *GetStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get status o k response
func (o *GetStatusOK) Code() int {
	return 200
}

func (o *GetStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status][%d] getStatusOK %s", 200, payload)
}

func (o *GetStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status][%d] getStatusOK %s", 200, payload)
}

func (o *GetStatusOK) GetPayload() *GetStatusOKBody {
	return o.Payload
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusUnprocessableEntity creates a GetStatusUnprocessableEntity with default headers values
func NewGetStatusUnprocessableEntity() *GetStatusUnprocessableEntity {
	return &GetStatusUnprocessableEntity{}
}

/*
GetStatusUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetStatusUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get status unprocessable entity response has a 2xx status code
func (o *GetStatusUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get status unprocessable entity response has a 3xx status code
func (o *GetStatusUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get status unprocessable entity response has a 4xx status code
func (o *GetStatusUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get status unprocessable entity response has a 5xx status code
func (o *GetStatusUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get status unprocessable entity response a status code equal to that given
func (o *GetStatusUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get status unprocessable entity response
func (o *GetStatusUnprocessableEntity) Code() int {
	return 422
}

func (o *GetStatusUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status][%d] getStatusUnprocessableEntity %s", 422, payload)
}

func (o *GetStatusUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status][%d] getStatusUnprocessableEntity %s", 422, payload)
}

func (o *GetStatusUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStatusUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStatusDefault creates a GetStatusDefault with default headers values
func NewGetStatusDefault(code int) *GetStatusDefault {
	return &GetStatusDefault{
		_statusCode: code,
	}
}

/*
GetStatusDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetStatusDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get status default response has a 2xx status code
func (o *GetStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get status default response has a 3xx status code
func (o *GetStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get status default response has a 4xx status code
func (o *GetStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get status default response has a 5xx status code
func (o *GetStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get status default response a status code equal to that given
func (o *GetStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get status default response
func (o *GetStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetStatusDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status][%d] getStatus default %s", o._statusCode, payload)
}

func (o *GetStatusDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status][%d] getStatus default %s", o._statusCode, payload)
}

func (o *GetStatusDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetStatusOKBody get status o k body
swagger:model GetStatusOKBody
*/
type GetStatusOKBody struct {

	// data
	// Required: true
	Data *models.GeneralStatus `json:"data"`
}

// Validate validates this get status o k body
func (o *GetStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getStatusOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStatusOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStatusOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get status o k body based on the context it is used
func (o *GetStatusOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStatusOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStatusOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getStatusOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

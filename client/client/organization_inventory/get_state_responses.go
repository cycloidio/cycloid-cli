// Code generated by go-swagger; DO NOT EDIT.

package organization_inventory

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

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetStateReader is a Reader for the GetState structure.
type GetStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStateOK creates a GetStateOK with default headers values
func NewGetStateOK() *GetStateOK {
	return &GetStateOK{}
}

/*
GetStateOK describes a response with status code 200, with default header values.

Get the state of an organization
*/
type GetStateOK struct {
	Payload *GetStateOKBody
}

// IsSuccess returns true when this get state o k response has a 2xx status code
func (o *GetStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get state o k response has a 3xx status code
func (o *GetStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get state o k response has a 4xx status code
func (o *GetStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get state o k response has a 5xx status code
func (o *GetStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get state o k response a status code equal to that given
func (o *GetStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get state o k response
func (o *GetStateOK) Code() int {
	return 200
}

func (o *GetStateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getStateOK %s", 200, payload)
}

func (o *GetStateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getStateOK %s", 200, payload)
}

func (o *GetStateOK) GetPayload() *GetStateOKBody {
	return o.Payload
}

func (o *GetStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStateForbidden creates a GetStateForbidden with default headers values
func NewGetStateForbidden() *GetStateForbidden {
	return &GetStateForbidden{}
}

/*
GetStateForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetStateForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get state forbidden response has a 2xx status code
func (o *GetStateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get state forbidden response has a 3xx status code
func (o *GetStateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get state forbidden response has a 4xx status code
func (o *GetStateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get state forbidden response has a 5xx status code
func (o *GetStateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get state forbidden response a status code equal to that given
func (o *GetStateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get state forbidden response
func (o *GetStateForbidden) Code() int {
	return 403
}

func (o *GetStateForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getStateForbidden %s", 403, payload)
}

func (o *GetStateForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getStateForbidden %s", 403, payload)
}

func (o *GetStateForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStateNotFound creates a GetStateNotFound with default headers values
func NewGetStateNotFound() *GetStateNotFound {
	return &GetStateNotFound{}
}

/*
GetStateNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetStateNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get state not found response has a 2xx status code
func (o *GetStateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get state not found response has a 3xx status code
func (o *GetStateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get state not found response has a 4xx status code
func (o *GetStateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get state not found response has a 5xx status code
func (o *GetStateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get state not found response a status code equal to that given
func (o *GetStateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get state not found response
func (o *GetStateNotFound) Code() int {
	return 404
}

func (o *GetStateNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getStateNotFound %s", 404, payload)
}

func (o *GetStateNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getStateNotFound %s", 404, payload)
}

func (o *GetStateNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStateDefault creates a GetStateDefault with default headers values
func NewGetStateDefault(code int) *GetStateDefault {
	return &GetStateDefault{
		_statusCode: code,
	}
}

/*
GetStateDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetStateDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get state default response has a 2xx status code
func (o *GetStateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get state default response has a 3xx status code
func (o *GetStateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get state default response has a 4xx status code
func (o *GetStateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get state default response has a 5xx status code
func (o *GetStateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get state default response a status code equal to that given
func (o *GetStateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get state default response
func (o *GetStateDefault) Code() int {
	return o._statusCode
}

func (o *GetStateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getState default %s", o._statusCode, payload)
}

func (o *GetStateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /inventory][%d] getState default %s", o._statusCode, payload)
}

func (o *GetStateDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetStateOKBody The state of an organization
swagger:model GetStateOKBody
*/
type GetStateOKBody struct {

	// data
	// Required: true
	Data interface{} `json:"data"`
}

// Validate validates this get state o k body
func (o *GetStateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStateOKBody) validateData(formats strfmt.Registry) error {

	if o.Data == nil {
		return errors.Required("getStateOK"+"."+"data", "body", nil)
	}

	return nil
}

// ContextValidate validates this get state o k body based on context it is used
func (o *GetStateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStateOKBody) UnmarshalBinary(b []byte) error {
	var res GetStateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organization_forms

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

// ValuesRefFormsReader is a Reader for the ValuesRefForms structure.
type ValuesRefFormsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValuesRefFormsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValuesRefFormsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewValuesRefFormsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValuesRefFormsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewValuesRefFormsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewValuesRefFormsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValuesRefFormsOK creates a ValuesRefFormsOK with default headers values
func NewValuesRefFormsOK() *ValuesRefFormsOK {
	return &ValuesRefFormsOK{}
}

/*
ValuesRefFormsOK describes a response with status code 200, with default header values.

The result pulling the values from the url
*/
type ValuesRefFormsOK struct {
	Payload *ValuesRefFormsOKBody
}

// IsSuccess returns true when this values ref forms o k response has a 2xx status code
func (o *ValuesRefFormsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this values ref forms o k response has a 3xx status code
func (o *ValuesRefFormsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this values ref forms o k response has a 4xx status code
func (o *ValuesRefFormsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this values ref forms o k response has a 5xx status code
func (o *ValuesRefFormsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this values ref forms o k response a status code equal to that given
func (o *ValuesRefFormsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the values ref forms o k response
func (o *ValuesRefFormsOK) Code() int {
	return 200
}

func (o *ValuesRefFormsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsOK %s", 200, payload)
}

func (o *ValuesRefFormsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsOK %s", 200, payload)
}

func (o *ValuesRefFormsOK) GetPayload() *ValuesRefFormsOKBody {
	return o.Payload
}

func (o *ValuesRefFormsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ValuesRefFormsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValuesRefFormsForbidden creates a ValuesRefFormsForbidden with default headers values
func NewValuesRefFormsForbidden() *ValuesRefFormsForbidden {
	return &ValuesRefFormsForbidden{}
}

/*
ValuesRefFormsForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type ValuesRefFormsForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this values ref forms forbidden response has a 2xx status code
func (o *ValuesRefFormsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this values ref forms forbidden response has a 3xx status code
func (o *ValuesRefFormsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this values ref forms forbidden response has a 4xx status code
func (o *ValuesRefFormsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this values ref forms forbidden response has a 5xx status code
func (o *ValuesRefFormsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this values ref forms forbidden response a status code equal to that given
func (o *ValuesRefFormsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the values ref forms forbidden response
func (o *ValuesRefFormsForbidden) Code() int {
	return 403
}

func (o *ValuesRefFormsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsForbidden %s", 403, payload)
}

func (o *ValuesRefFormsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsForbidden %s", 403, payload)
}

func (o *ValuesRefFormsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValuesRefFormsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewValuesRefFormsNotFound creates a ValuesRefFormsNotFound with default headers values
func NewValuesRefFormsNotFound() *ValuesRefFormsNotFound {
	return &ValuesRefFormsNotFound{}
}

/*
ValuesRefFormsNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type ValuesRefFormsNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this values ref forms not found response has a 2xx status code
func (o *ValuesRefFormsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this values ref forms not found response has a 3xx status code
func (o *ValuesRefFormsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this values ref forms not found response has a 4xx status code
func (o *ValuesRefFormsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this values ref forms not found response has a 5xx status code
func (o *ValuesRefFormsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this values ref forms not found response a status code equal to that given
func (o *ValuesRefFormsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the values ref forms not found response
func (o *ValuesRefFormsNotFound) Code() int {
	return 404
}

func (o *ValuesRefFormsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsNotFound %s", 404, payload)
}

func (o *ValuesRefFormsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsNotFound %s", 404, payload)
}

func (o *ValuesRefFormsNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValuesRefFormsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewValuesRefFormsUnprocessableEntity creates a ValuesRefFormsUnprocessableEntity with default headers values
func NewValuesRefFormsUnprocessableEntity() *ValuesRefFormsUnprocessableEntity {
	return &ValuesRefFormsUnprocessableEntity{}
}

/*
ValuesRefFormsUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type ValuesRefFormsUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this values ref forms unprocessable entity response has a 2xx status code
func (o *ValuesRefFormsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this values ref forms unprocessable entity response has a 3xx status code
func (o *ValuesRefFormsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this values ref forms unprocessable entity response has a 4xx status code
func (o *ValuesRefFormsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this values ref forms unprocessable entity response has a 5xx status code
func (o *ValuesRefFormsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this values ref forms unprocessable entity response a status code equal to that given
func (o *ValuesRefFormsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the values ref forms unprocessable entity response
func (o *ValuesRefFormsUnprocessableEntity) Code() int {
	return 422
}

func (o *ValuesRefFormsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsUnprocessableEntity %s", 422, payload)
}

func (o *ValuesRefFormsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefFormsUnprocessableEntity %s", 422, payload)
}

func (o *ValuesRefFormsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValuesRefFormsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewValuesRefFormsDefault creates a ValuesRefFormsDefault with default headers values
func NewValuesRefFormsDefault(code int) *ValuesRefFormsDefault {
	return &ValuesRefFormsDefault{
		_statusCode: code,
	}
}

/*
ValuesRefFormsDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type ValuesRefFormsDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this values ref forms default response has a 2xx status code
func (o *ValuesRefFormsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this values ref forms default response has a 3xx status code
func (o *ValuesRefFormsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this values ref forms default response has a 4xx status code
func (o *ValuesRefFormsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this values ref forms default response has a 5xx status code
func (o *ValuesRefFormsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this values ref forms default response a status code equal to that given
func (o *ValuesRefFormsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the values ref forms default response
func (o *ValuesRefFormsDefault) Code() int {
	return o._statusCode
}

func (o *ValuesRefFormsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefForms default %s", o._statusCode, payload)
}

func (o *ValuesRefFormsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/forms/values_ref][%d] valuesRefForms default %s", o._statusCode, payload)
}

func (o *ValuesRefFormsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValuesRefFormsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
ValuesRefFormsOKBody values ref forms o k body
swagger:model ValuesRefFormsOKBody
*/
type ValuesRefFormsOKBody struct {

	// data
	// Required: true
	Data interface{} `json:"data"`
}

// Validate validates this values ref forms o k body
func (o *ValuesRefFormsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValuesRefFormsOKBody) validateData(formats strfmt.Registry) error {

	if o.Data == nil {
		return errors.Required("valuesRefFormsOK"+"."+"data", "body", nil)
	}

	return nil
}

// ContextValidate validates this values ref forms o k body based on context it is used
func (o *ValuesRefFormsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValuesRefFormsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValuesRefFormsOKBody) UnmarshalBinary(b []byte) error {
	var res ValuesRefFormsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

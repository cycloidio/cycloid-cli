// Code generated by go-swagger; DO NOT EDIT.

package organization_workers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetWorkersReader is a Reader for the GetWorkers structure.
type GetWorkersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetWorkersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetWorkersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWorkersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkersOK creates a GetWorkersOK with default headers values
func NewGetWorkersOK() *GetWorkersOK {
	return &GetWorkersOK{}
}

/*
GetWorkersOK describes a response with status code 200, with default header values.

List of the workers which authenticated user has access to.
*/
type GetWorkersOK struct {
	Payload *GetWorkersOKBody
}

// IsSuccess returns true when this get workers o k response has a 2xx status code
func (o *GetWorkersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workers o k response has a 3xx status code
func (o *GetWorkersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workers o k response has a 4xx status code
func (o *GetWorkersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workers o k response has a 5xx status code
func (o *GetWorkersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workers o k response a status code equal to that given
func (o *GetWorkersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workers o k response
func (o *GetWorkersOK) Code() int {
	return 200
}

func (o *GetWorkersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkersOK %s", 200, payload)
}

func (o *GetWorkersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkersOK %s", 200, payload)
}

func (o *GetWorkersOK) GetPayload() *GetWorkersOKBody {
	return o.Payload
}

func (o *GetWorkersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkersNotFound creates a GetWorkersNotFound with default headers values
func NewGetWorkersNotFound() *GetWorkersNotFound {
	return &GetWorkersNotFound{}
}

/*
GetWorkersNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetWorkersNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get workers not found response has a 2xx status code
func (o *GetWorkersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workers not found response has a 3xx status code
func (o *GetWorkersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workers not found response has a 4xx status code
func (o *GetWorkersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workers not found response has a 5xx status code
func (o *GetWorkersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workers not found response a status code equal to that given
func (o *GetWorkersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workers not found response
func (o *GetWorkersNotFound) Code() int {
	return 404
}

func (o *GetWorkersNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkersNotFound %s", 404, payload)
}

func (o *GetWorkersNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkersNotFound %s", 404, payload)
}

func (o *GetWorkersNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetWorkersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWorkersUnprocessableEntity creates a GetWorkersUnprocessableEntity with default headers values
func NewGetWorkersUnprocessableEntity() *GetWorkersUnprocessableEntity {
	return &GetWorkersUnprocessableEntity{}
}

/*
GetWorkersUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetWorkersUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get workers unprocessable entity response has a 2xx status code
func (o *GetWorkersUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workers unprocessable entity response has a 3xx status code
func (o *GetWorkersUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workers unprocessable entity response has a 4xx status code
func (o *GetWorkersUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workers unprocessable entity response has a 5xx status code
func (o *GetWorkersUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get workers unprocessable entity response a status code equal to that given
func (o *GetWorkersUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get workers unprocessable entity response
func (o *GetWorkersUnprocessableEntity) Code() int {
	return 422
}

func (o *GetWorkersUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkersUnprocessableEntity %s", 422, payload)
}

func (o *GetWorkersUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkersUnprocessableEntity %s", 422, payload)
}

func (o *GetWorkersUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetWorkersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWorkersDefault creates a GetWorkersDefault with default headers values
func NewGetWorkersDefault(code int) *GetWorkersDefault {
	return &GetWorkersDefault{
		_statusCode: code,
	}
}

/*
GetWorkersDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetWorkersDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get workers default response has a 2xx status code
func (o *GetWorkersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get workers default response has a 3xx status code
func (o *GetWorkersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get workers default response has a 4xx status code
func (o *GetWorkersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get workers default response has a 5xx status code
func (o *GetWorkersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get workers default response a status code equal to that given
func (o *GetWorkersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get workers default response
func (o *GetWorkersDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkers default %s", o._statusCode, payload)
}

func (o *GetWorkersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/workers][%d] getWorkers default %s", o._statusCode, payload)
}

func (o *GetWorkersDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetWorkersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetWorkersOKBody get workers o k body
swagger:model GetWorkersOKBody
*/
type GetWorkersOKBody struct {

	// data
	// Required: true
	Data []*models.Worker `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get workers o k body
func (o *GetWorkersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkersOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getWorkersOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkersOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkersOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetWorkersOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getWorkersOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkersOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkersOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get workers o k body based on the context it is used
func (o *GetWorkersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkersOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkersOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkersOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetWorkersOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkersOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getWorkersOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkersOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

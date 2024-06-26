// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines

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

// GetPipelinesReader is a Reader for the GetPipelines structure.
type GetPipelinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPipelinesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetPipelinesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPipelinesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPipelinesOK creates a GetPipelinesOK with default headers values
func NewGetPipelinesOK() *GetPipelinesOK {
	return &GetPipelinesOK{}
}

/*
GetPipelinesOK describes a response with status code 200, with default header values.

List of all the pipelines which authenticated user has access to.
*/
type GetPipelinesOK struct {
	Payload *GetPipelinesOKBody
}

// IsSuccess returns true when this get pipelines o k response has a 2xx status code
func (o *GetPipelinesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipelines o k response has a 3xx status code
func (o *GetPipelinesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipelines o k response has a 4xx status code
func (o *GetPipelinesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipelines o k response has a 5xx status code
func (o *GetPipelinesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipelines o k response a status code equal to that given
func (o *GetPipelinesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pipelines o k response
func (o *GetPipelinesOK) Code() int {
	return 200
}

func (o *GetPipelinesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesOK %s", 200, payload)
}

func (o *GetPipelinesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesOK %s", 200, payload)
}

func (o *GetPipelinesOK) GetPayload() *GetPipelinesOKBody {
	return o.Payload
}

func (o *GetPipelinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPipelinesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelinesNotFound creates a GetPipelinesNotFound with default headers values
func NewGetPipelinesNotFound() *GetPipelinesNotFound {
	return &GetPipelinesNotFound{}
}

/*
GetPipelinesNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetPipelinesNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get pipelines not found response has a 2xx status code
func (o *GetPipelinesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipelines not found response has a 3xx status code
func (o *GetPipelinesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipelines not found response has a 4xx status code
func (o *GetPipelinesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipelines not found response has a 5xx status code
func (o *GetPipelinesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipelines not found response a status code equal to that given
func (o *GetPipelinesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get pipelines not found response
func (o *GetPipelinesNotFound) Code() int {
	return 404
}

func (o *GetPipelinesNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesNotFound %s", 404, payload)
}

func (o *GetPipelinesNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesNotFound %s", 404, payload)
}

func (o *GetPipelinesNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelinesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelinesUnprocessableEntity creates a GetPipelinesUnprocessableEntity with default headers values
func NewGetPipelinesUnprocessableEntity() *GetPipelinesUnprocessableEntity {
	return &GetPipelinesUnprocessableEntity{}
}

/*
GetPipelinesUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetPipelinesUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get pipelines unprocessable entity response has a 2xx status code
func (o *GetPipelinesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipelines unprocessable entity response has a 3xx status code
func (o *GetPipelinesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipelines unprocessable entity response has a 4xx status code
func (o *GetPipelinesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipelines unprocessable entity response has a 5xx status code
func (o *GetPipelinesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipelines unprocessable entity response a status code equal to that given
func (o *GetPipelinesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get pipelines unprocessable entity response
func (o *GetPipelinesUnprocessableEntity) Code() int {
	return 422
}

func (o *GetPipelinesUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesUnprocessableEntity %s", 422, payload)
}

func (o *GetPipelinesUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesUnprocessableEntity %s", 422, payload)
}

func (o *GetPipelinesUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelinesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelinesDefault creates a GetPipelinesDefault with default headers values
func NewGetPipelinesDefault(code int) *GetPipelinesDefault {
	return &GetPipelinesDefault{
		_statusCode: code,
	}
}

/*
GetPipelinesDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetPipelinesDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get pipelines default response has a 2xx status code
func (o *GetPipelinesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get pipelines default response has a 3xx status code
func (o *GetPipelinesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get pipelines default response has a 4xx status code
func (o *GetPipelinesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get pipelines default response has a 5xx status code
func (o *GetPipelinesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get pipelines default response a status code equal to that given
func (o *GetPipelinesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get pipelines default response
func (o *GetPipelinesDefault) Code() int {
	return o._statusCode
}

func (o *GetPipelinesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelines default %s", o._statusCode, payload)
}

func (o *GetPipelinesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelines default %s", o._statusCode, payload)
}

func (o *GetPipelinesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelinesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetPipelinesOKBody get pipelines o k body
swagger:model GetPipelinesOKBody
*/
type GetPipelinesOKBody struct {

	// data
	// Required: true
	Data []*models.Pipeline `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get pipelines o k body
func (o *GetPipelinesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetPipelinesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPipelinesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPipelinesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPipelinesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPipelinesOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getPipelinesOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPipelinesOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPipelinesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get pipelines o k body based on the context it is used
func (o *GetPipelinesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetPipelinesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPipelinesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPipelinesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPipelinesOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPipelinesOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPipelinesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPipelinesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPipelinesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPipelinesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

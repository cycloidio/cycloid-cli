// Code generated by go-swagger; DO NOT EDIT.

package organization_resource_pools

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

// CreateResourcePoolReader is a Reader for the CreateResourcePool structure.
type CreateResourcePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourcePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateResourcePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateResourcePoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateResourcePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateResourcePoolUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateResourcePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateResourcePoolOK creates a CreateResourcePoolOK with default headers values
func NewCreateResourcePoolOK() *CreateResourcePoolOK {
	return &CreateResourcePoolOK{}
}

/*
CreateResourcePoolOK describes a response with status code 200, with default header values.

New resource pool created in the organization.
*/
type CreateResourcePoolOK struct {
	Payload *CreateResourcePoolOKBody
}

// IsSuccess returns true when this create resource pool o k response has a 2xx status code
func (o *CreateResourcePoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create resource pool o k response has a 3xx status code
func (o *CreateResourcePoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource pool o k response has a 4xx status code
func (o *CreateResourcePoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create resource pool o k response has a 5xx status code
func (o *CreateResourcePoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource pool o k response a status code equal to that given
func (o *CreateResourcePoolOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create resource pool o k response
func (o *CreateResourcePoolOK) Code() int {
	return 200
}

func (o *CreateResourcePoolOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolOK %s", 200, payload)
}

func (o *CreateResourcePoolOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolOK %s", 200, payload)
}

func (o *CreateResourcePoolOK) GetPayload() *CreateResourcePoolOKBody {
	return o.Payload
}

func (o *CreateResourcePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateResourcePoolOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourcePoolForbidden creates a CreateResourcePoolForbidden with default headers values
func NewCreateResourcePoolForbidden() *CreateResourcePoolForbidden {
	return &CreateResourcePoolForbidden{}
}

/*
CreateResourcePoolForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateResourcePoolForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create resource pool forbidden response has a 2xx status code
func (o *CreateResourcePoolForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create resource pool forbidden response has a 3xx status code
func (o *CreateResourcePoolForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource pool forbidden response has a 4xx status code
func (o *CreateResourcePoolForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create resource pool forbidden response has a 5xx status code
func (o *CreateResourcePoolForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource pool forbidden response a status code equal to that given
func (o *CreateResourcePoolForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create resource pool forbidden response
func (o *CreateResourcePoolForbidden) Code() int {
	return 403
}

func (o *CreateResourcePoolForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolForbidden %s", 403, payload)
}

func (o *CreateResourcePoolForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolForbidden %s", 403, payload)
}

func (o *CreateResourcePoolForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateResourcePoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateResourcePoolNotFound creates a CreateResourcePoolNotFound with default headers values
func NewCreateResourcePoolNotFound() *CreateResourcePoolNotFound {
	return &CreateResourcePoolNotFound{}
}

/*
CreateResourcePoolNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateResourcePoolNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create resource pool not found response has a 2xx status code
func (o *CreateResourcePoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create resource pool not found response has a 3xx status code
func (o *CreateResourcePoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource pool not found response has a 4xx status code
func (o *CreateResourcePoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create resource pool not found response has a 5xx status code
func (o *CreateResourcePoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource pool not found response a status code equal to that given
func (o *CreateResourcePoolNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create resource pool not found response
func (o *CreateResourcePoolNotFound) Code() int {
	return 404
}

func (o *CreateResourcePoolNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolNotFound %s", 404, payload)
}

func (o *CreateResourcePoolNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolNotFound %s", 404, payload)
}

func (o *CreateResourcePoolNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateResourcePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateResourcePoolUnprocessableEntity creates a CreateResourcePoolUnprocessableEntity with default headers values
func NewCreateResourcePoolUnprocessableEntity() *CreateResourcePoolUnprocessableEntity {
	return &CreateResourcePoolUnprocessableEntity{}
}

/*
CreateResourcePoolUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateResourcePoolUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create resource pool unprocessable entity response has a 2xx status code
func (o *CreateResourcePoolUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create resource pool unprocessable entity response has a 3xx status code
func (o *CreateResourcePoolUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create resource pool unprocessable entity response has a 4xx status code
func (o *CreateResourcePoolUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create resource pool unprocessable entity response has a 5xx status code
func (o *CreateResourcePoolUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create resource pool unprocessable entity response a status code equal to that given
func (o *CreateResourcePoolUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create resource pool unprocessable entity response
func (o *CreateResourcePoolUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateResourcePoolUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolUnprocessableEntity %s", 422, payload)
}

func (o *CreateResourcePoolUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePoolUnprocessableEntity %s", 422, payload)
}

func (o *CreateResourcePoolUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateResourcePoolUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateResourcePoolDefault creates a CreateResourcePoolDefault with default headers values
func NewCreateResourcePoolDefault(code int) *CreateResourcePoolDefault {
	return &CreateResourcePoolDefault{
		_statusCode: code,
	}
}

/*
CreateResourcePoolDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateResourcePoolDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create resource pool default response has a 2xx status code
func (o *CreateResourcePoolDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create resource pool default response has a 3xx status code
func (o *CreateResourcePoolDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create resource pool default response has a 4xx status code
func (o *CreateResourcePoolDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create resource pool default response has a 5xx status code
func (o *CreateResourcePoolDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create resource pool default response a status code equal to that given
func (o *CreateResourcePoolDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create resource pool default response
func (o *CreateResourcePoolDefault) Code() int {
	return o._statusCode
}

func (o *CreateResourcePoolDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePool default %s", o._statusCode, payload)
}

func (o *CreateResourcePoolDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/resource_pools][%d] createResourcePool default %s", o._statusCode, payload)
}

func (o *CreateResourcePoolDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateResourcePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
CreateResourcePoolOKBody create resource pool o k body
swagger:model CreateResourcePoolOKBody
*/
type CreateResourcePoolOKBody struct {

	// data
	// Required: true
	Data *models.ResourcePool `json:"data"`
}

// Validate validates this create resource pool o k body
func (o *CreateResourcePoolOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateResourcePoolOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createResourcePoolOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createResourcePoolOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createResourcePoolOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create resource pool o k body based on the context it is used
func (o *CreateResourcePoolOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateResourcePoolOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createResourcePoolOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createResourcePoolOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateResourcePoolOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateResourcePoolOKBody) UnmarshalBinary(b []byte) error {
	var res CreateResourcePoolOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

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

// CreateServiceCatalogReader is a Reader for the CreateServiceCatalog structure.
type CreateServiceCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateServiceCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateServiceCatalogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateServiceCatalogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateServiceCatalogUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServiceCatalogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServiceCatalogOK creates a CreateServiceCatalogOK with default headers values
func NewCreateServiceCatalogOK() *CreateServiceCatalogOK {
	return &CreateServiceCatalogOK{}
}

/*
CreateServiceCatalogOK describes a response with status code 200, with default header values.

The information of the service catalog.
*/
type CreateServiceCatalogOK struct {
	Payload *CreateServiceCatalogOKBody
}

// IsSuccess returns true when this create service catalog o k response has a 2xx status code
func (o *CreateServiceCatalogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create service catalog o k response has a 3xx status code
func (o *CreateServiceCatalogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service catalog o k response has a 4xx status code
func (o *CreateServiceCatalogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service catalog o k response has a 5xx status code
func (o *CreateServiceCatalogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create service catalog o k response a status code equal to that given
func (o *CreateServiceCatalogOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create service catalog o k response
func (o *CreateServiceCatalogOK) Code() int {
	return 200
}

func (o *CreateServiceCatalogOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogOK %s", 200, payload)
}

func (o *CreateServiceCatalogOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogOK %s", 200, payload)
}

func (o *CreateServiceCatalogOK) GetPayload() *CreateServiceCatalogOKBody {
	return o.Payload
}

func (o *CreateServiceCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateServiceCatalogOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceCatalogForbidden creates a CreateServiceCatalogForbidden with default headers values
func NewCreateServiceCatalogForbidden() *CreateServiceCatalogForbidden {
	return &CreateServiceCatalogForbidden{}
}

/*
CreateServiceCatalogForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateServiceCatalogForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create service catalog forbidden response has a 2xx status code
func (o *CreateServiceCatalogForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service catalog forbidden response has a 3xx status code
func (o *CreateServiceCatalogForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service catalog forbidden response has a 4xx status code
func (o *CreateServiceCatalogForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service catalog forbidden response has a 5xx status code
func (o *CreateServiceCatalogForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create service catalog forbidden response a status code equal to that given
func (o *CreateServiceCatalogForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create service catalog forbidden response
func (o *CreateServiceCatalogForbidden) Code() int {
	return 403
}

func (o *CreateServiceCatalogForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogForbidden %s", 403, payload)
}

func (o *CreateServiceCatalogForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogForbidden %s", 403, payload)
}

func (o *CreateServiceCatalogForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServiceCatalogNotFound creates a CreateServiceCatalogNotFound with default headers values
func NewCreateServiceCatalogNotFound() *CreateServiceCatalogNotFound {
	return &CreateServiceCatalogNotFound{}
}

/*
CreateServiceCatalogNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateServiceCatalogNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create service catalog not found response has a 2xx status code
func (o *CreateServiceCatalogNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service catalog not found response has a 3xx status code
func (o *CreateServiceCatalogNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service catalog not found response has a 4xx status code
func (o *CreateServiceCatalogNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service catalog not found response has a 5xx status code
func (o *CreateServiceCatalogNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create service catalog not found response a status code equal to that given
func (o *CreateServiceCatalogNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create service catalog not found response
func (o *CreateServiceCatalogNotFound) Code() int {
	return 404
}

func (o *CreateServiceCatalogNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogNotFound %s", 404, payload)
}

func (o *CreateServiceCatalogNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogNotFound %s", 404, payload)
}

func (o *CreateServiceCatalogNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServiceCatalogUnprocessableEntity creates a CreateServiceCatalogUnprocessableEntity with default headers values
func NewCreateServiceCatalogUnprocessableEntity() *CreateServiceCatalogUnprocessableEntity {
	return &CreateServiceCatalogUnprocessableEntity{}
}

/*
CreateServiceCatalogUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateServiceCatalogUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create service catalog unprocessable entity response has a 2xx status code
func (o *CreateServiceCatalogUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service catalog unprocessable entity response has a 3xx status code
func (o *CreateServiceCatalogUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service catalog unprocessable entity response has a 4xx status code
func (o *CreateServiceCatalogUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service catalog unprocessable entity response has a 5xx status code
func (o *CreateServiceCatalogUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create service catalog unprocessable entity response a status code equal to that given
func (o *CreateServiceCatalogUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create service catalog unprocessable entity response
func (o *CreateServiceCatalogUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateServiceCatalogUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogUnprocessableEntity %s", 422, payload)
}

func (o *CreateServiceCatalogUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalogUnprocessableEntity %s", 422, payload)
}

func (o *CreateServiceCatalogUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateServiceCatalogDefault creates a CreateServiceCatalogDefault with default headers values
func NewCreateServiceCatalogDefault(code int) *CreateServiceCatalogDefault {
	return &CreateServiceCatalogDefault{
		_statusCode: code,
	}
}

/*
CreateServiceCatalogDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateServiceCatalogDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create service catalog default response has a 2xx status code
func (o *CreateServiceCatalogDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create service catalog default response has a 3xx status code
func (o *CreateServiceCatalogDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create service catalog default response has a 4xx status code
func (o *CreateServiceCatalogDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create service catalog default response has a 5xx status code
func (o *CreateServiceCatalogDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create service catalog default response a status code equal to that given
func (o *CreateServiceCatalogDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create service catalog default response
func (o *CreateServiceCatalogDefault) Code() int {
	return o._statusCode
}

func (o *CreateServiceCatalogDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalog default %s", o._statusCode, payload)
}

func (o *CreateServiceCatalogDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalogs][%d] createServiceCatalog default %s", o._statusCode, payload)
}

func (o *CreateServiceCatalogDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateServiceCatalogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
CreateServiceCatalogOKBody create service catalog o k body
swagger:model CreateServiceCatalogOKBody
*/
type CreateServiceCatalogOKBody struct {

	// data
	// Required: true
	Data *models.ServiceCatalog `json:"data"`
}

// Validate validates this create service catalog o k body
func (o *CreateServiceCatalogOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServiceCatalogOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createServiceCatalogOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createServiceCatalogOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createServiceCatalogOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create service catalog o k body based on the context it is used
func (o *CreateServiceCatalogOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateServiceCatalogOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createServiceCatalogOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createServiceCatalogOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateServiceCatalogOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateServiceCatalogOKBody) UnmarshalBinary(b []byte) error {
	var res CreateServiceCatalogOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

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

// CreateKpiReader is a Reader for the CreateKpi structure.
type CreateKpiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateKpiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateKpiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateKpiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateKpiUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateKpiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateKpiOK creates a CreateKpiOK with default headers values
func NewCreateKpiOK() *CreateKpiOK {
	return &CreateKpiOK{}
}

/*
CreateKpiOK describes a response with status code 200, with default header values.

KPI has been configured
*/
type CreateKpiOK struct {
	Payload *CreateKpiOKBody
}

// IsSuccess returns true when this create kpi o k response has a 2xx status code
func (o *CreateKpiOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create kpi o k response has a 3xx status code
func (o *CreateKpiOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kpi o k response has a 4xx status code
func (o *CreateKpiOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create kpi o k response has a 5xx status code
func (o *CreateKpiOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create kpi o k response a status code equal to that given
func (o *CreateKpiOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create kpi o k response
func (o *CreateKpiOK) Code() int {
	return 200
}

func (o *CreateKpiOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpiOK %s", 200, payload)
}

func (o *CreateKpiOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpiOK %s", 200, payload)
}

func (o *CreateKpiOK) GetPayload() *CreateKpiOKBody {
	return o.Payload
}

func (o *CreateKpiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateKpiOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateKpiForbidden creates a CreateKpiForbidden with default headers values
func NewCreateKpiForbidden() *CreateKpiForbidden {
	return &CreateKpiForbidden{}
}

/*
CreateKpiForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateKpiForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create kpi forbidden response has a 2xx status code
func (o *CreateKpiForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create kpi forbidden response has a 3xx status code
func (o *CreateKpiForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kpi forbidden response has a 4xx status code
func (o *CreateKpiForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create kpi forbidden response has a 5xx status code
func (o *CreateKpiForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create kpi forbidden response a status code equal to that given
func (o *CreateKpiForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create kpi forbidden response
func (o *CreateKpiForbidden) Code() int {
	return 403
}

func (o *CreateKpiForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpiForbidden %s", 403, payload)
}

func (o *CreateKpiForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpiForbidden %s", 403, payload)
}

func (o *CreateKpiForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateKpiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateKpiUnprocessableEntity creates a CreateKpiUnprocessableEntity with default headers values
func NewCreateKpiUnprocessableEntity() *CreateKpiUnprocessableEntity {
	return &CreateKpiUnprocessableEntity{}
}

/*
CreateKpiUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateKpiUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create kpi unprocessable entity response has a 2xx status code
func (o *CreateKpiUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create kpi unprocessable entity response has a 3xx status code
func (o *CreateKpiUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create kpi unprocessable entity response has a 4xx status code
func (o *CreateKpiUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create kpi unprocessable entity response has a 5xx status code
func (o *CreateKpiUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create kpi unprocessable entity response a status code equal to that given
func (o *CreateKpiUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create kpi unprocessable entity response
func (o *CreateKpiUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateKpiUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpiUnprocessableEntity %s", 422, payload)
}

func (o *CreateKpiUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpiUnprocessableEntity %s", 422, payload)
}

func (o *CreateKpiUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateKpiUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateKpiDefault creates a CreateKpiDefault with default headers values
func NewCreateKpiDefault(code int) *CreateKpiDefault {
	return &CreateKpiDefault{
		_statusCode: code,
	}
}

/*
CreateKpiDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateKpiDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create kpi default response has a 2xx status code
func (o *CreateKpiDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create kpi default response has a 3xx status code
func (o *CreateKpiDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create kpi default response has a 4xx status code
func (o *CreateKpiDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create kpi default response has a 5xx status code
func (o *CreateKpiDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create kpi default response a status code equal to that given
func (o *CreateKpiDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create kpi default response
func (o *CreateKpiDefault) Code() int {
	return o._statusCode
}

func (o *CreateKpiDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpi default %s", o._statusCode, payload)
}

func (o *CreateKpiDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/kpis][%d] createKpi default %s", o._statusCode, payload)
}

func (o *CreateKpiDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateKpiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
CreateKpiOKBody create kpi o k body
swagger:model CreateKpiOKBody
*/
type CreateKpiOKBody struct {

	// data
	// Required: true
	Data *models.KPI `json:"data"`
}

// Validate validates this create kpi o k body
func (o *CreateKpiOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateKpiOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createKpiOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createKpiOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createKpiOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create kpi o k body based on the context it is used
func (o *CreateKpiOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateKpiOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createKpiOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createKpiOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateKpiOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateKpiOKBody) UnmarshalBinary(b []byte) error {
	var res CreateKpiOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

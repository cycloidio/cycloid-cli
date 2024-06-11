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

// UpdateKpiReader is a Reader for the UpdateKpi structure.
type UpdateKpiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateKpiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateKpiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateKpiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateKpiNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewUpdateKpiLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateKpiUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateKpiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateKpiOK creates a UpdateKpiOK with default headers values
func NewUpdateKpiOK() *UpdateKpiOK {
	return &UpdateKpiOK{}
}

/*
UpdateKpiOK describes a response with status code 200, with default header values.

Success update
*/
type UpdateKpiOK struct {
	Payload *UpdateKpiOKBody
}

// IsSuccess returns true when this update kpi o k response has a 2xx status code
func (o *UpdateKpiOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update kpi o k response has a 3xx status code
func (o *UpdateKpiOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kpi o k response has a 4xx status code
func (o *UpdateKpiOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update kpi o k response has a 5xx status code
func (o *UpdateKpiOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update kpi o k response a status code equal to that given
func (o *UpdateKpiOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update kpi o k response
func (o *UpdateKpiOK) Code() int {
	return 200
}

func (o *UpdateKpiOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiOK %s", 200, payload)
}

func (o *UpdateKpiOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiOK %s", 200, payload)
}

func (o *UpdateKpiOK) GetPayload() *UpdateKpiOKBody {
	return o.Payload
}

func (o *UpdateKpiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateKpiOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKpiForbidden creates a UpdateKpiForbidden with default headers values
func NewUpdateKpiForbidden() *UpdateKpiForbidden {
	return &UpdateKpiForbidden{}
}

/*
UpdateKpiForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateKpiForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update kpi forbidden response has a 2xx status code
func (o *UpdateKpiForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kpi forbidden response has a 3xx status code
func (o *UpdateKpiForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kpi forbidden response has a 4xx status code
func (o *UpdateKpiForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kpi forbidden response has a 5xx status code
func (o *UpdateKpiForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update kpi forbidden response a status code equal to that given
func (o *UpdateKpiForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update kpi forbidden response
func (o *UpdateKpiForbidden) Code() int {
	return 403
}

func (o *UpdateKpiForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiForbidden %s", 403, payload)
}

func (o *UpdateKpiForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiForbidden %s", 403, payload)
}

func (o *UpdateKpiForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateKpiNotFound creates a UpdateKpiNotFound with default headers values
func NewUpdateKpiNotFound() *UpdateKpiNotFound {
	return &UpdateKpiNotFound{}
}

/*
UpdateKpiNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateKpiNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update kpi not found response has a 2xx status code
func (o *UpdateKpiNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kpi not found response has a 3xx status code
func (o *UpdateKpiNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kpi not found response has a 4xx status code
func (o *UpdateKpiNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kpi not found response has a 5xx status code
func (o *UpdateKpiNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update kpi not found response a status code equal to that given
func (o *UpdateKpiNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update kpi not found response
func (o *UpdateKpiNotFound) Code() int {
	return 404
}

func (o *UpdateKpiNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiNotFound %s", 404, payload)
}

func (o *UpdateKpiNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiNotFound %s", 404, payload)
}

func (o *UpdateKpiNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateKpiLengthRequired creates a UpdateKpiLengthRequired with default headers values
func NewUpdateKpiLengthRequired() *UpdateKpiLengthRequired {
	return &UpdateKpiLengthRequired{}
}

/*
UpdateKpiLengthRequired describes a response with status code 411, with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type UpdateKpiLengthRequired struct {
}

// IsSuccess returns true when this update kpi length required response has a 2xx status code
func (o *UpdateKpiLengthRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kpi length required response has a 3xx status code
func (o *UpdateKpiLengthRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kpi length required response has a 4xx status code
func (o *UpdateKpiLengthRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kpi length required response has a 5xx status code
func (o *UpdateKpiLengthRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this update kpi length required response a status code equal to that given
func (o *UpdateKpiLengthRequired) IsCode(code int) bool {
	return code == 411
}

// Code gets the status code for the update kpi length required response
func (o *UpdateKpiLengthRequired) Code() int {
	return 411
}

func (o *UpdateKpiLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiLengthRequired", 411)
}

func (o *UpdateKpiLengthRequired) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiLengthRequired", 411)
}

func (o *UpdateKpiLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKpiUnprocessableEntity creates a UpdateKpiUnprocessableEntity with default headers values
func NewUpdateKpiUnprocessableEntity() *UpdateKpiUnprocessableEntity {
	return &UpdateKpiUnprocessableEntity{}
}

/*
UpdateKpiUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateKpiUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update kpi unprocessable entity response has a 2xx status code
func (o *UpdateKpiUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update kpi unprocessable entity response has a 3xx status code
func (o *UpdateKpiUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update kpi unprocessable entity response has a 4xx status code
func (o *UpdateKpiUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update kpi unprocessable entity response has a 5xx status code
func (o *UpdateKpiUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update kpi unprocessable entity response a status code equal to that given
func (o *UpdateKpiUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update kpi unprocessable entity response
func (o *UpdateKpiUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateKpiUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiUnprocessableEntity %s", 422, payload)
}

func (o *UpdateKpiUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiUnprocessableEntity %s", 422, payload)
}

func (o *UpdateKpiUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateKpiDefault creates a UpdateKpiDefault with default headers values
func NewUpdateKpiDefault(code int) *UpdateKpiDefault {
	return &UpdateKpiDefault{
		_statusCode: code,
	}
}

/*
UpdateKpiDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateKpiDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update kpi default response has a 2xx status code
func (o *UpdateKpiDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update kpi default response has a 3xx status code
func (o *UpdateKpiDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update kpi default response has a 4xx status code
func (o *UpdateKpiDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update kpi default response has a 5xx status code
func (o *UpdateKpiDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update kpi default response a status code equal to that given
func (o *UpdateKpiDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update kpi default response
func (o *UpdateKpiDefault) Code() int {
	return o._statusCode
}

func (o *UpdateKpiDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpi default %s", o._statusCode, payload)
}

func (o *UpdateKpiDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpi default %s", o._statusCode, payload)
}

func (o *UpdateKpiDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
UpdateKpiOKBody update kpi o k body
swagger:model UpdateKpiOKBody
*/
type UpdateKpiOKBody struct {

	// data
	// Required: true
	Data *models.KPI `json:"data"`
}

// Validate validates this update kpi o k body
func (o *UpdateKpiOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateKpiOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateKpiOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateKpiOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateKpiOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update kpi o k body based on the context it is used
func (o *UpdateKpiOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateKpiOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateKpiOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateKpiOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateKpiOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateKpiOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateKpiOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

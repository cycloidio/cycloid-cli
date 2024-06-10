// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

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

// GetKpisReader is a Reader for the GetKpis structure.
type GetKpisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKpisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKpisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetKpisForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetKpisUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetKpisDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKpisOK creates a GetKpisOK with default headers values
func NewGetKpisOK() *GetKpisOK {
	return &GetKpisOK{}
}

/*
GetKpisOK describes a response with status code 200, with default header values.

The list of the KPIs
*/
type GetKpisOK struct {
	Payload *GetKpisOKBody
}

// IsSuccess returns true when this get kpis o k response has a 2xx status code
func (o *GetKpisOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kpis o k response has a 3xx status code
func (o *GetKpisOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kpis o k response has a 4xx status code
func (o *GetKpisOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kpis o k response has a 5xx status code
func (o *GetKpisOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kpis o k response a status code equal to that given
func (o *GetKpisOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kpis o k response
func (o *GetKpisOK) Code() int {
	return 200
}

func (o *GetKpisOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpisOK %s", 200, payload)
}

func (o *GetKpisOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpisOK %s", 200, payload)
}

func (o *GetKpisOK) GetPayload() *GetKpisOKBody {
	return o.Payload
}

func (o *GetKpisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetKpisOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKpisForbidden creates a GetKpisForbidden with default headers values
func NewGetKpisForbidden() *GetKpisForbidden {
	return &GetKpisForbidden{}
}

/*
GetKpisForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetKpisForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get kpis forbidden response has a 2xx status code
func (o *GetKpisForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kpis forbidden response has a 3xx status code
func (o *GetKpisForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kpis forbidden response has a 4xx status code
func (o *GetKpisForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kpis forbidden response has a 5xx status code
func (o *GetKpisForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kpis forbidden response a status code equal to that given
func (o *GetKpisForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get kpis forbidden response
func (o *GetKpisForbidden) Code() int {
	return 403
}

func (o *GetKpisForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpisForbidden %s", 403, payload)
}

func (o *GetKpisForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpisForbidden %s", 403, payload)
}

func (o *GetKpisForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpisForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetKpisUnprocessableEntity creates a GetKpisUnprocessableEntity with default headers values
func NewGetKpisUnprocessableEntity() *GetKpisUnprocessableEntity {
	return &GetKpisUnprocessableEntity{}
}

/*
GetKpisUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetKpisUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get kpis unprocessable entity response has a 2xx status code
func (o *GetKpisUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kpis unprocessable entity response has a 3xx status code
func (o *GetKpisUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kpis unprocessable entity response has a 4xx status code
func (o *GetKpisUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kpis unprocessable entity response has a 5xx status code
func (o *GetKpisUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get kpis unprocessable entity response a status code equal to that given
func (o *GetKpisUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get kpis unprocessable entity response
func (o *GetKpisUnprocessableEntity) Code() int {
	return 422
}

func (o *GetKpisUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpisUnprocessableEntity %s", 422, payload)
}

func (o *GetKpisUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpisUnprocessableEntity %s", 422, payload)
}

func (o *GetKpisUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpisUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetKpisDefault creates a GetKpisDefault with default headers values
func NewGetKpisDefault(code int) *GetKpisDefault {
	return &GetKpisDefault{
		_statusCode: code,
	}
}

/*
GetKpisDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetKpisDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get kpis default response has a 2xx status code
func (o *GetKpisDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get kpis default response has a 3xx status code
func (o *GetKpisDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get kpis default response has a 4xx status code
func (o *GetKpisDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get kpis default response has a 5xx status code
func (o *GetKpisDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get kpis default response a status code equal to that given
func (o *GetKpisDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get kpis default response
func (o *GetKpisDefault) Code() int {
	return o._statusCode
}

func (o *GetKpisDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpis default %s", o._statusCode, payload)
}

func (o *GetKpisDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis][%d] getKpis default %s", o._statusCode, payload)
}

func (o *GetKpisDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpisDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetKpisOKBody get kpis o k body
swagger:model GetKpisOKBody
*/
type GetKpisOKBody struct {

	// data
	// Required: true
	Data []*models.KPI `json:"data"`

	// pagination
	Pagination *models.Pagination `json:"pagination,omitempty"`
}

// Validate validates this get kpis o k body
func (o *GetKpisOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetKpisOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getKpisOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getKpisOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getKpisOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetKpisOKBody) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getKpisOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getKpisOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get kpis o k body based on the context it is used
func (o *GetKpisOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetKpisOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getKpisOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getKpisOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetKpisOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if swag.IsZero(o.Pagination) { // not required
			return nil
		}

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getKpisOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getKpisOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetKpisOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKpisOKBody) UnmarshalBinary(b []byte) error {
	var res GetKpisOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

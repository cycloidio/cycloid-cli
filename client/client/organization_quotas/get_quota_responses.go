// Code generated by go-swagger; DO NOT EDIT.

package organization_quotas

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

// GetQuotaReader is a Reader for the GetQuota structure.
type GetQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetQuotaOK creates a GetQuotaOK with default headers values
func NewGetQuotaOK() *GetQuotaOK {
	return &GetQuotaOK{}
}

/*
GetQuotaOK describes a response with status code 200, with default header values.

Quota available in the organization with such canonical.
*/
type GetQuotaOK struct {
	Payload *GetQuotaOKBody
}

// IsSuccess returns true when this get quota o k response has a 2xx status code
func (o *GetQuotaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quota o k response has a 3xx status code
func (o *GetQuotaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quota o k response has a 4xx status code
func (o *GetQuotaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quota o k response has a 5xx status code
func (o *GetQuotaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quota o k response a status code equal to that given
func (o *GetQuotaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get quota o k response
func (o *GetQuotaOK) Code() int {
	return 200
}

func (o *GetQuotaOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuotaOK %s", 200, payload)
}

func (o *GetQuotaOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuotaOK %s", 200, payload)
}

func (o *GetQuotaOK) GetPayload() *GetQuotaOKBody {
	return o.Payload
}

func (o *GetQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetQuotaOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotaForbidden creates a GetQuotaForbidden with default headers values
func NewGetQuotaForbidden() *GetQuotaForbidden {
	return &GetQuotaForbidden{}
}

/*
GetQuotaForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetQuotaForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get quota forbidden response has a 2xx status code
func (o *GetQuotaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quota forbidden response has a 3xx status code
func (o *GetQuotaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quota forbidden response has a 4xx status code
func (o *GetQuotaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quota forbidden response has a 5xx status code
func (o *GetQuotaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quota forbidden response a status code equal to that given
func (o *GetQuotaForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get quota forbidden response
func (o *GetQuotaForbidden) Code() int {
	return 403
}

func (o *GetQuotaForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuotaForbidden %s", 403, payload)
}

func (o *GetQuotaForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuotaForbidden %s", 403, payload)
}

func (o *GetQuotaForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetQuotaNotFound creates a GetQuotaNotFound with default headers values
func NewGetQuotaNotFound() *GetQuotaNotFound {
	return &GetQuotaNotFound{}
}

/*
GetQuotaNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetQuotaNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get quota not found response has a 2xx status code
func (o *GetQuotaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quota not found response has a 3xx status code
func (o *GetQuotaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quota not found response has a 4xx status code
func (o *GetQuotaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quota not found response has a 5xx status code
func (o *GetQuotaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quota not found response a status code equal to that given
func (o *GetQuotaNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get quota not found response
func (o *GetQuotaNotFound) Code() int {
	return 404
}

func (o *GetQuotaNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuotaNotFound %s", 404, payload)
}

func (o *GetQuotaNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuotaNotFound %s", 404, payload)
}

func (o *GetQuotaNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetQuotaDefault creates a GetQuotaDefault with default headers values
func NewGetQuotaDefault(code int) *GetQuotaDefault {
	return &GetQuotaDefault{
		_statusCode: code,
	}
}

/*
GetQuotaDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetQuotaDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get quota default response has a 2xx status code
func (o *GetQuotaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get quota default response has a 3xx status code
func (o *GetQuotaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get quota default response has a 4xx status code
func (o *GetQuotaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get quota default response has a 5xx status code
func (o *GetQuotaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get quota default response a status code equal to that given
func (o *GetQuotaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get quota default response
func (o *GetQuotaDefault) Code() int {
	return o._statusCode
}

func (o *GetQuotaDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuota default %s", o._statusCode, payload)
}

func (o *GetQuotaDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/quotas/{quota_id}][%d] getQuota default %s", o._statusCode, payload)
}

func (o *GetQuotaDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetQuotaOKBody get quota o k body
swagger:model GetQuotaOKBody
*/
type GetQuotaOKBody struct {

	// data
	// Required: true
	Data *models.Quota `json:"data"`
}

// Validate validates this get quota o k body
func (o *GetQuotaOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetQuotaOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getQuotaOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getQuotaOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getQuotaOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get quota o k body based on the context it is used
func (o *GetQuotaOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetQuotaOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getQuotaOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getQuotaOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetQuotaOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetQuotaOKBody) UnmarshalBinary(b []byte) error {
	var res GetQuotaOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

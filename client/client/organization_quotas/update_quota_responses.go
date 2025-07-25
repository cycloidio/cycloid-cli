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

// UpdateQuotaReader is a Reader for the UpdateQuota structure.
type UpdateQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateQuotaUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateQuotaOK creates a UpdateQuotaOK with default headers values
func NewUpdateQuotaOK() *UpdateQuotaOK {
	return &UpdateQuotaOK{}
}

/*
UpdateQuotaOK describes a response with status code 200, with default header values.

Updated quota belonging to the organization.
*/
type UpdateQuotaOK struct {
	Payload *UpdateQuotaOKBody
}

// IsSuccess returns true when this update quota o k response has a 2xx status code
func (o *UpdateQuotaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update quota o k response has a 3xx status code
func (o *UpdateQuotaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update quota o k response has a 4xx status code
func (o *UpdateQuotaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update quota o k response has a 5xx status code
func (o *UpdateQuotaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update quota o k response a status code equal to that given
func (o *UpdateQuotaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update quota o k response
func (o *UpdateQuotaOK) Code() int {
	return 200
}

func (o *UpdateQuotaOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaOK %s", 200, payload)
}

func (o *UpdateQuotaOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaOK %s", 200, payload)
}

func (o *UpdateQuotaOK) GetPayload() *UpdateQuotaOKBody {
	return o.Payload
}

func (o *UpdateQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateQuotaOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateQuotaForbidden creates a UpdateQuotaForbidden with default headers values
func NewUpdateQuotaForbidden() *UpdateQuotaForbidden {
	return &UpdateQuotaForbidden{}
}

/*
UpdateQuotaForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateQuotaForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update quota forbidden response has a 2xx status code
func (o *UpdateQuotaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update quota forbidden response has a 3xx status code
func (o *UpdateQuotaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update quota forbidden response has a 4xx status code
func (o *UpdateQuotaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update quota forbidden response has a 5xx status code
func (o *UpdateQuotaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update quota forbidden response a status code equal to that given
func (o *UpdateQuotaForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update quota forbidden response
func (o *UpdateQuotaForbidden) Code() int {
	return 403
}

func (o *UpdateQuotaForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaForbidden %s", 403, payload)
}

func (o *UpdateQuotaForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaForbidden %s", 403, payload)
}

func (o *UpdateQuotaForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateQuotaNotFound creates a UpdateQuotaNotFound with default headers values
func NewUpdateQuotaNotFound() *UpdateQuotaNotFound {
	return &UpdateQuotaNotFound{}
}

/*
UpdateQuotaNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateQuotaNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update quota not found response has a 2xx status code
func (o *UpdateQuotaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update quota not found response has a 3xx status code
func (o *UpdateQuotaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update quota not found response has a 4xx status code
func (o *UpdateQuotaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update quota not found response has a 5xx status code
func (o *UpdateQuotaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update quota not found response a status code equal to that given
func (o *UpdateQuotaNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update quota not found response
func (o *UpdateQuotaNotFound) Code() int {
	return 404
}

func (o *UpdateQuotaNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaNotFound %s", 404, payload)
}

func (o *UpdateQuotaNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaNotFound %s", 404, payload)
}

func (o *UpdateQuotaNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateQuotaUnprocessableEntity creates a UpdateQuotaUnprocessableEntity with default headers values
func NewUpdateQuotaUnprocessableEntity() *UpdateQuotaUnprocessableEntity {
	return &UpdateQuotaUnprocessableEntity{}
}

/*
UpdateQuotaUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateQuotaUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update quota unprocessable entity response has a 2xx status code
func (o *UpdateQuotaUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update quota unprocessable entity response has a 3xx status code
func (o *UpdateQuotaUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update quota unprocessable entity response has a 4xx status code
func (o *UpdateQuotaUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update quota unprocessable entity response has a 5xx status code
func (o *UpdateQuotaUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update quota unprocessable entity response a status code equal to that given
func (o *UpdateQuotaUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update quota unprocessable entity response
func (o *UpdateQuotaUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateQuotaUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaUnprocessableEntity %s", 422, payload)
}

func (o *UpdateQuotaUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuotaUnprocessableEntity %s", 422, payload)
}

func (o *UpdateQuotaUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateQuotaUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateQuotaDefault creates a UpdateQuotaDefault with default headers values
func NewUpdateQuotaDefault(code int) *UpdateQuotaDefault {
	return &UpdateQuotaDefault{
		_statusCode: code,
	}
}

/*
UpdateQuotaDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateQuotaDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update quota default response has a 2xx status code
func (o *UpdateQuotaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update quota default response has a 3xx status code
func (o *UpdateQuotaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update quota default response has a 4xx status code
func (o *UpdateQuotaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update quota default response has a 5xx status code
func (o *UpdateQuotaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update quota default response a status code equal to that given
func (o *UpdateQuotaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update quota default response
func (o *UpdateQuotaDefault) Code() int {
	return o._statusCode
}

func (o *UpdateQuotaDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuota default %s", o._statusCode, payload)
}

func (o *UpdateQuotaDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/quotas/{quota_id}][%d] updateQuota default %s", o._statusCode, payload)
}

func (o *UpdateQuotaDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
UpdateQuotaOKBody update quota o k body
swagger:model UpdateQuotaOKBody
*/
type UpdateQuotaOKBody struct {

	// data
	// Required: true
	Data *models.Quota `json:"data"`
}

// Validate validates this update quota o k body
func (o *UpdateQuotaOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateQuotaOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateQuotaOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateQuotaOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateQuotaOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update quota o k body based on the context it is used
func (o *UpdateQuotaOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateQuotaOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateQuotaOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateQuotaOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateQuotaOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateQuotaOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateQuotaOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

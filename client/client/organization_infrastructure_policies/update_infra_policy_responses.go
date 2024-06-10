// Code generated by go-swagger; DO NOT EDIT.

package organization_infrastructure_policies

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

// UpdateInfraPolicyReader is a Reader for the UpdateInfraPolicy structure.
type UpdateInfraPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInfraPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInfraPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateInfraPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInfraPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateInfraPolicyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateInfraPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateInfraPolicyOK creates a UpdateInfraPolicyOK with default headers values
func NewUpdateInfraPolicyOK() *UpdateInfraPolicyOK {
	return &UpdateInfraPolicyOK{}
}

/*
UpdateInfraPolicyOK describes a response with status code 200, with default header values.

InfraPolicy updated.
*/
type UpdateInfraPolicyOK struct {
	Payload *UpdateInfraPolicyOKBody
}

// IsSuccess returns true when this update infra policy o k response has a 2xx status code
func (o *UpdateInfraPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update infra policy o k response has a 3xx status code
func (o *UpdateInfraPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update infra policy o k response has a 4xx status code
func (o *UpdateInfraPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update infra policy o k response has a 5xx status code
func (o *UpdateInfraPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update infra policy o k response a status code equal to that given
func (o *UpdateInfraPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update infra policy o k response
func (o *UpdateInfraPolicyOK) Code() int {
	return 200
}

func (o *UpdateInfraPolicyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyOK %s", 200, payload)
}

func (o *UpdateInfraPolicyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyOK %s", 200, payload)
}

func (o *UpdateInfraPolicyOK) GetPayload() *UpdateInfraPolicyOKBody {
	return o.Payload
}

func (o *UpdateInfraPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateInfraPolicyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInfraPolicyForbidden creates a UpdateInfraPolicyForbidden with default headers values
func NewUpdateInfraPolicyForbidden() *UpdateInfraPolicyForbidden {
	return &UpdateInfraPolicyForbidden{}
}

/*
UpdateInfraPolicyForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateInfraPolicyForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update infra policy forbidden response has a 2xx status code
func (o *UpdateInfraPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update infra policy forbidden response has a 3xx status code
func (o *UpdateInfraPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update infra policy forbidden response has a 4xx status code
func (o *UpdateInfraPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update infra policy forbidden response has a 5xx status code
func (o *UpdateInfraPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update infra policy forbidden response a status code equal to that given
func (o *UpdateInfraPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update infra policy forbidden response
func (o *UpdateInfraPolicyForbidden) Code() int {
	return 403
}

func (o *UpdateInfraPolicyForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyForbidden %s", 403, payload)
}

func (o *UpdateInfraPolicyForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyForbidden %s", 403, payload)
}

func (o *UpdateInfraPolicyForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateInfraPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateInfraPolicyNotFound creates a UpdateInfraPolicyNotFound with default headers values
func NewUpdateInfraPolicyNotFound() *UpdateInfraPolicyNotFound {
	return &UpdateInfraPolicyNotFound{}
}

/*
UpdateInfraPolicyNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateInfraPolicyNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update infra policy not found response has a 2xx status code
func (o *UpdateInfraPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update infra policy not found response has a 3xx status code
func (o *UpdateInfraPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update infra policy not found response has a 4xx status code
func (o *UpdateInfraPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update infra policy not found response has a 5xx status code
func (o *UpdateInfraPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update infra policy not found response a status code equal to that given
func (o *UpdateInfraPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update infra policy not found response
func (o *UpdateInfraPolicyNotFound) Code() int {
	return 404
}

func (o *UpdateInfraPolicyNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyNotFound %s", 404, payload)
}

func (o *UpdateInfraPolicyNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyNotFound %s", 404, payload)
}

func (o *UpdateInfraPolicyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateInfraPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateInfraPolicyUnprocessableEntity creates a UpdateInfraPolicyUnprocessableEntity with default headers values
func NewUpdateInfraPolicyUnprocessableEntity() *UpdateInfraPolicyUnprocessableEntity {
	return &UpdateInfraPolicyUnprocessableEntity{}
}

/*
UpdateInfraPolicyUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateInfraPolicyUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update infra policy unprocessable entity response has a 2xx status code
func (o *UpdateInfraPolicyUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update infra policy unprocessable entity response has a 3xx status code
func (o *UpdateInfraPolicyUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update infra policy unprocessable entity response has a 4xx status code
func (o *UpdateInfraPolicyUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update infra policy unprocessable entity response has a 5xx status code
func (o *UpdateInfraPolicyUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update infra policy unprocessable entity response a status code equal to that given
func (o *UpdateInfraPolicyUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update infra policy unprocessable entity response
func (o *UpdateInfraPolicyUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateInfraPolicyUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyUnprocessableEntity %s", 422, payload)
}

func (o *UpdateInfraPolicyUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicyUnprocessableEntity %s", 422, payload)
}

func (o *UpdateInfraPolicyUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateInfraPolicyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateInfraPolicyDefault creates a UpdateInfraPolicyDefault with default headers values
func NewUpdateInfraPolicyDefault(code int) *UpdateInfraPolicyDefault {
	return &UpdateInfraPolicyDefault{
		_statusCode: code,
	}
}

/*
UpdateInfraPolicyDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateInfraPolicyDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update infra policy default response has a 2xx status code
func (o *UpdateInfraPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update infra policy default response has a 3xx status code
func (o *UpdateInfraPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update infra policy default response has a 4xx status code
func (o *UpdateInfraPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update infra policy default response has a 5xx status code
func (o *UpdateInfraPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update infra policy default response a status code equal to that given
func (o *UpdateInfraPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update infra policy default response
func (o *UpdateInfraPolicyDefault) Code() int {
	return o._statusCode
}

func (o *UpdateInfraPolicyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicy default %s", o._statusCode, payload)
}

func (o *UpdateInfraPolicyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] updateInfraPolicy default %s", o._statusCode, payload)
}

func (o *UpdateInfraPolicyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateInfraPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
UpdateInfraPolicyOKBody update infra policy o k body
swagger:model UpdateInfraPolicyOKBody
*/
type UpdateInfraPolicyOKBody struct {

	// data
	// Required: true
	Data *models.InfraPolicy `json:"data"`
}

// Validate validates this update infra policy o k body
func (o *UpdateInfraPolicyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateInfraPolicyOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateInfraPolicyOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateInfraPolicyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateInfraPolicyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update infra policy o k body based on the context it is used
func (o *UpdateInfraPolicyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateInfraPolicyOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateInfraPolicyOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateInfraPolicyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateInfraPolicyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateInfraPolicyOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateInfraPolicyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

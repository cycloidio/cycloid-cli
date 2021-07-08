// Code generated by go-swagger; DO NOT EDIT.

package organization_external_backends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// UpdateExternalBackendReader is a Reader for the UpdateExternalBackend structure.
type UpdateExternalBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateExternalBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateExternalBackendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateExternalBackendForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewUpdateExternalBackendLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateExternalBackendUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateExternalBackendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateExternalBackendOK creates a UpdateExternalBackendOK with default headers values
func NewUpdateExternalBackendOK() *UpdateExternalBackendOK {
	return &UpdateExternalBackendOK{}
}

/*UpdateExternalBackendOK handles this case with default header values.

Success update
*/
type UpdateExternalBackendOK struct {
	Payload *UpdateExternalBackendOKBody
}

func (o *UpdateExternalBackendOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] updateExternalBackendOK  %+v", 200, o.Payload)
}

func (o *UpdateExternalBackendOK) GetPayload() *UpdateExternalBackendOKBody {
	return o.Payload
}

func (o *UpdateExternalBackendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateExternalBackendOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalBackendForbidden creates a UpdateExternalBackendForbidden with default headers values
func NewUpdateExternalBackendForbidden() *UpdateExternalBackendForbidden {
	return &UpdateExternalBackendForbidden{}
}

/*UpdateExternalBackendForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateExternalBackendForbidden struct {
	Payload *models.ErrorPayload
}

func (o *UpdateExternalBackendForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] updateExternalBackendForbidden  %+v", 403, o.Payload)
}

func (o *UpdateExternalBackendForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateExternalBackendForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalBackendLengthRequired creates a UpdateExternalBackendLengthRequired with default headers values
func NewUpdateExternalBackendLengthRequired() *UpdateExternalBackendLengthRequired {
	return &UpdateExternalBackendLengthRequired{}
}

/*UpdateExternalBackendLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type UpdateExternalBackendLengthRequired struct {
}

func (o *UpdateExternalBackendLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] updateExternalBackendLengthRequired ", 411)
}

func (o *UpdateExternalBackendLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateExternalBackendUnprocessableEntity creates a UpdateExternalBackendUnprocessableEntity with default headers values
func NewUpdateExternalBackendUnprocessableEntity() *UpdateExternalBackendUnprocessableEntity {
	return &UpdateExternalBackendUnprocessableEntity{}
}

/*UpdateExternalBackendUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateExternalBackendUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *UpdateExternalBackendUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] updateExternalBackendUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateExternalBackendUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateExternalBackendUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalBackendDefault creates a UpdateExternalBackendDefault with default headers values
func NewUpdateExternalBackendDefault(code int) *UpdateExternalBackendDefault {
	return &UpdateExternalBackendDefault{
		_statusCode: code,
	}
}

/*UpdateExternalBackendDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateExternalBackendDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the update external backend default response
func (o *UpdateExternalBackendDefault) Code() int {
	return o._statusCode
}

func (o *UpdateExternalBackendDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] updateExternalBackend default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateExternalBackendDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateExternalBackendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateExternalBackendOKBody update external backend o k body
swagger:model UpdateExternalBackendOKBody
*/
type UpdateExternalBackendOKBody struct {

	// data
	// Required: true
	Data *models.ExternalBackend `json:"data"`
}

// Validate validates this update external backend o k body
func (o *UpdateExternalBackendOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateExternalBackendOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateExternalBackendOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateExternalBackendOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateExternalBackendOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateExternalBackendOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateExternalBackendOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

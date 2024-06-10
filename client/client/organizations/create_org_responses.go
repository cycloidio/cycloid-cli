// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// CreateOrgReader is a Reader for the CreateOrg structure.
type CreateOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 411:
		result := NewCreateOrgLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateOrgUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOrgDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrgOK creates a CreateOrgOK with default headers values
func NewCreateOrgOK() *CreateOrgOK {
	return &CreateOrgOK{}
}

/*
CreateOrgOK describes a response with status code 200, with default header values.

Organization created. The body contains the information of the new created organization.
*/
type CreateOrgOK struct {
	Payload *CreateOrgOKBody
}

// IsSuccess returns true when this create org o k response has a 2xx status code
func (o *CreateOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create org o k response has a 3xx status code
func (o *CreateOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org o k response has a 4xx status code
func (o *CreateOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create org o k response has a 5xx status code
func (o *CreateOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create org o k response a status code equal to that given
func (o *CreateOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create org o k response
func (o *CreateOrgOK) Code() int {
	return 200
}

func (o *CreateOrgOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations][%d] createOrgOK %s", 200, payload)
}

func (o *CreateOrgOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations][%d] createOrgOK %s", 200, payload)
}

func (o *CreateOrgOK) GetPayload() *CreateOrgOKBody {
	return o.Payload
}

func (o *CreateOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateOrgOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgLengthRequired creates a CreateOrgLengthRequired with default headers values
func NewCreateOrgLengthRequired() *CreateOrgLengthRequired {
	return &CreateOrgLengthRequired{}
}

/*
CreateOrgLengthRequired describes a response with status code 411, with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreateOrgLengthRequired struct {
}

// IsSuccess returns true when this create org length required response has a 2xx status code
func (o *CreateOrgLengthRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create org length required response has a 3xx status code
func (o *CreateOrgLengthRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org length required response has a 4xx status code
func (o *CreateOrgLengthRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this create org length required response has a 5xx status code
func (o *CreateOrgLengthRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this create org length required response a status code equal to that given
func (o *CreateOrgLengthRequired) IsCode(code int) bool {
	return code == 411
}

// Code gets the status code for the create org length required response
func (o *CreateOrgLengthRequired) Code() int {
	return 411
}

func (o *CreateOrgLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrgLengthRequired", 411)
}

func (o *CreateOrgLengthRequired) String() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrgLengthRequired", 411)
}

func (o *CreateOrgLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrgUnprocessableEntity creates a CreateOrgUnprocessableEntity with default headers values
func NewCreateOrgUnprocessableEntity() *CreateOrgUnprocessableEntity {
	return &CreateOrgUnprocessableEntity{}
}

/*
CreateOrgUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateOrgUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create org unprocessable entity response has a 2xx status code
func (o *CreateOrgUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create org unprocessable entity response has a 3xx status code
func (o *CreateOrgUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org unprocessable entity response has a 4xx status code
func (o *CreateOrgUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create org unprocessable entity response has a 5xx status code
func (o *CreateOrgUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create org unprocessable entity response a status code equal to that given
func (o *CreateOrgUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create org unprocessable entity response
func (o *CreateOrgUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateOrgUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations][%d] createOrgUnprocessableEntity %s", 422, payload)
}

func (o *CreateOrgUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations][%d] createOrgUnprocessableEntity %s", 422, payload)
}

func (o *CreateOrgUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOrgUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateOrgDefault creates a CreateOrgDefault with default headers values
func NewCreateOrgDefault(code int) *CreateOrgDefault {
	return &CreateOrgDefault{
		_statusCode: code,
	}
}

/*
CreateOrgDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateOrgDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this create org default response has a 2xx status code
func (o *CreateOrgDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create org default response has a 3xx status code
func (o *CreateOrgDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create org default response has a 4xx status code
func (o *CreateOrgDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create org default response has a 5xx status code
func (o *CreateOrgDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create org default response a status code equal to that given
func (o *CreateOrgDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create org default response
func (o *CreateOrgDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrgDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations][%d] createOrg default %s", o._statusCode, payload)
}

func (o *CreateOrgDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations][%d] createOrg default %s", o._statusCode, payload)
}

func (o *CreateOrgDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOrgDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
CreateOrgOKBody create org o k body
swagger:model CreateOrgOKBody
*/
type CreateOrgOKBody struct {

	// data
	// Required: true
	Data *models.Organization `json:"data"`
}

// Validate validates this create org o k body
func (o *CreateOrgOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrgOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createOrgOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrgOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrgOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create org o k body based on the context it is used
func (o *CreateOrgOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrgOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrgOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrgOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrgOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrgOKBody) UnmarshalBinary(b []byte) error {
	var res CreateOrgOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

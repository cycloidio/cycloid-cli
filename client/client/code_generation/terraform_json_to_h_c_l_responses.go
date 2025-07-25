// Code generated by go-swagger; DO NOT EDIT.

package code_generation

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

// TerraformJSONToHCLReader is a Reader for the TerraformJSONToHCL structure.
type TerraformJSONToHCLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TerraformJSONToHCLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTerraformJSONToHCLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTerraformJSONToHCLForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTerraformJSONToHCLUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTerraformJSONToHCLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTerraformJSONToHCLOK creates a TerraformJSONToHCLOK with default headers values
func NewTerraformJSONToHCLOK() *TerraformJSONToHCLOK {
	return &TerraformJSONToHCLOK{}
}

/*
TerraformJSONToHCLOK describes a response with status code 200, with default header values.

The HCL translation of the config
*/
type TerraformJSONToHCLOK struct {
	Payload *TerraformJSONToHCLOKBody
}

// IsSuccess returns true when this terraform Json to h c l o k response has a 2xx status code
func (o *TerraformJSONToHCLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this terraform Json to h c l o k response has a 3xx status code
func (o *TerraformJSONToHCLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this terraform Json to h c l o k response has a 4xx status code
func (o *TerraformJSONToHCLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this terraform Json to h c l o k response has a 5xx status code
func (o *TerraformJSONToHCLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this terraform Json to h c l o k response a status code equal to that given
func (o *TerraformJSONToHCLOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the terraform Json to h c l o k response
func (o *TerraformJSONToHCLOK) Code() int {
	return 200
}

func (o *TerraformJSONToHCLOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJsonToHCLOK %s", 200, payload)
}

func (o *TerraformJSONToHCLOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJsonToHCLOK %s", 200, payload)
}

func (o *TerraformJSONToHCLOK) GetPayload() *TerraformJSONToHCLOKBody {
	return o.Payload
}

func (o *TerraformJSONToHCLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TerraformJSONToHCLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerraformJSONToHCLForbidden creates a TerraformJSONToHCLForbidden with default headers values
func NewTerraformJSONToHCLForbidden() *TerraformJSONToHCLForbidden {
	return &TerraformJSONToHCLForbidden{}
}

/*
TerraformJSONToHCLForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type TerraformJSONToHCLForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this terraform Json to h c l forbidden response has a 2xx status code
func (o *TerraformJSONToHCLForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this terraform Json to h c l forbidden response has a 3xx status code
func (o *TerraformJSONToHCLForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this terraform Json to h c l forbidden response has a 4xx status code
func (o *TerraformJSONToHCLForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this terraform Json to h c l forbidden response has a 5xx status code
func (o *TerraformJSONToHCLForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this terraform Json to h c l forbidden response a status code equal to that given
func (o *TerraformJSONToHCLForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the terraform Json to h c l forbidden response
func (o *TerraformJSONToHCLForbidden) Code() int {
	return 403
}

func (o *TerraformJSONToHCLForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJsonToHCLForbidden %s", 403, payload)
}

func (o *TerraformJSONToHCLForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJsonToHCLForbidden %s", 403, payload)
}

func (o *TerraformJSONToHCLForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *TerraformJSONToHCLForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTerraformJSONToHCLUnprocessableEntity creates a TerraformJSONToHCLUnprocessableEntity with default headers values
func NewTerraformJSONToHCLUnprocessableEntity() *TerraformJSONToHCLUnprocessableEntity {
	return &TerraformJSONToHCLUnprocessableEntity{}
}

/*
TerraformJSONToHCLUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type TerraformJSONToHCLUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this terraform Json to h c l unprocessable entity response has a 2xx status code
func (o *TerraformJSONToHCLUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this terraform Json to h c l unprocessable entity response has a 3xx status code
func (o *TerraformJSONToHCLUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this terraform Json to h c l unprocessable entity response has a 4xx status code
func (o *TerraformJSONToHCLUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this terraform Json to h c l unprocessable entity response has a 5xx status code
func (o *TerraformJSONToHCLUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this terraform Json to h c l unprocessable entity response a status code equal to that given
func (o *TerraformJSONToHCLUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the terraform Json to h c l unprocessable entity response
func (o *TerraformJSONToHCLUnprocessableEntity) Code() int {
	return 422
}

func (o *TerraformJSONToHCLUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJsonToHCLUnprocessableEntity %s", 422, payload)
}

func (o *TerraformJSONToHCLUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJsonToHCLUnprocessableEntity %s", 422, payload)
}

func (o *TerraformJSONToHCLUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *TerraformJSONToHCLUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewTerraformJSONToHCLDefault creates a TerraformJSONToHCLDefault with default headers values
func NewTerraformJSONToHCLDefault(code int) *TerraformJSONToHCLDefault {
	return &TerraformJSONToHCLDefault{
		_statusCode: code,
	}
}

/*
TerraformJSONToHCLDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type TerraformJSONToHCLDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this terraform JSON to h c l default response has a 2xx status code
func (o *TerraformJSONToHCLDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this terraform JSON to h c l default response has a 3xx status code
func (o *TerraformJSONToHCLDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this terraform JSON to h c l default response has a 4xx status code
func (o *TerraformJSONToHCLDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this terraform JSON to h c l default response has a 5xx status code
func (o *TerraformJSONToHCLDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this terraform JSON to h c l default response a status code equal to that given
func (o *TerraformJSONToHCLDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the terraform JSON to h c l default response
func (o *TerraformJSONToHCLDefault) Code() int {
	return o._statusCode
}

func (o *TerraformJSONToHCLDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJSONToHCL default %s", o._statusCode, payload)
}

func (o *TerraformJSONToHCLDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/code_generation/terraform/jsontohcl][%d] terraformJSONToHCL default %s", o._statusCode, payload)
}

func (o *TerraformJSONToHCLDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *TerraformJSONToHCLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
TerraformJSONToHCLOKBody terraform JSON to h c l o k body
swagger:model TerraformJSONToHCLOKBody
*/
type TerraformJSONToHCLOKBody struct {

	// data
	// Required: true
	Data *models.TerraformHCLConfig `json:"data"`
}

// Validate validates this terraform JSON to h c l o k body
func (o *TerraformJSONToHCLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TerraformJSONToHCLOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("terraformJsonToHCLOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraformJsonToHCLOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraformJsonToHCLOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this terraform JSON to h c l o k body based on the context it is used
func (o *TerraformJSONToHCLOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TerraformJSONToHCLOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraformJsonToHCLOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraformJsonToHCLOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TerraformJSONToHCLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TerraformJSONToHCLOKBody) UnmarshalBinary(b []byte) error {
	var res TerraformJSONToHCLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

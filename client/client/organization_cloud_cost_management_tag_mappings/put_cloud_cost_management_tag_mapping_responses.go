// Code generated by go-swagger; DO NOT EDIT.

package organization_cloud_cost_management_tag_mappings

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

// PutCloudCostManagementTagMappingReader is a Reader for the PutCloudCostManagementTagMapping structure.
type PutCloudCostManagementTagMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCloudCostManagementTagMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCloudCostManagementTagMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 411:
		result := NewPutCloudCostManagementTagMappingLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutCloudCostManagementTagMappingUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutCloudCostManagementTagMappingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCloudCostManagementTagMappingOK creates a PutCloudCostManagementTagMappingOK with default headers values
func NewPutCloudCostManagementTagMappingOK() *PutCloudCostManagementTagMappingOK {
	return &PutCloudCostManagementTagMappingOK{}
}

/*
PutCloudCostManagementTagMappingOK describes a response with status code 200, with default header values.

The new or updated Cloud Cost Management Tag Mapping
*/
type PutCloudCostManagementTagMappingOK struct {
	Payload *PutCloudCostManagementTagMappingOKBody
}

// IsSuccess returns true when this put cloud cost management tag mapping o k response has a 2xx status code
func (o *PutCloudCostManagementTagMappingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put cloud cost management tag mapping o k response has a 3xx status code
func (o *PutCloudCostManagementTagMappingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cloud cost management tag mapping o k response has a 4xx status code
func (o *PutCloudCostManagementTagMappingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put cloud cost management tag mapping o k response has a 5xx status code
func (o *PutCloudCostManagementTagMappingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put cloud cost management tag mapping o k response a status code equal to that given
func (o *PutCloudCostManagementTagMappingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put cloud cost management tag mapping o k response
func (o *PutCloudCostManagementTagMappingOK) Code() int {
	return 200
}

func (o *PutCloudCostManagementTagMappingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMappingOK %s", 200, payload)
}

func (o *PutCloudCostManagementTagMappingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMappingOK %s", 200, payload)
}

func (o *PutCloudCostManagementTagMappingOK) GetPayload() *PutCloudCostManagementTagMappingOKBody {
	return o.Payload
}

func (o *PutCloudCostManagementTagMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCloudCostManagementTagMappingOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCloudCostManagementTagMappingLengthRequired creates a PutCloudCostManagementTagMappingLengthRequired with default headers values
func NewPutCloudCostManagementTagMappingLengthRequired() *PutCloudCostManagementTagMappingLengthRequired {
	return &PutCloudCostManagementTagMappingLengthRequired{}
}

/*
PutCloudCostManagementTagMappingLengthRequired describes a response with status code 411, with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type PutCloudCostManagementTagMappingLengthRequired struct {
}

// IsSuccess returns true when this put cloud cost management tag mapping length required response has a 2xx status code
func (o *PutCloudCostManagementTagMappingLengthRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put cloud cost management tag mapping length required response has a 3xx status code
func (o *PutCloudCostManagementTagMappingLengthRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cloud cost management tag mapping length required response has a 4xx status code
func (o *PutCloudCostManagementTagMappingLengthRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this put cloud cost management tag mapping length required response has a 5xx status code
func (o *PutCloudCostManagementTagMappingLengthRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this put cloud cost management tag mapping length required response a status code equal to that given
func (o *PutCloudCostManagementTagMappingLengthRequired) IsCode(code int) bool {
	return code == 411
}

// Code gets the status code for the put cloud cost management tag mapping length required response
func (o *PutCloudCostManagementTagMappingLengthRequired) Code() int {
	return 411
}

func (o *PutCloudCostManagementTagMappingLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMappingLengthRequired", 411)
}

func (o *PutCloudCostManagementTagMappingLengthRequired) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMappingLengthRequired", 411)
}

func (o *PutCloudCostManagementTagMappingLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCloudCostManagementTagMappingUnprocessableEntity creates a PutCloudCostManagementTagMappingUnprocessableEntity with default headers values
func NewPutCloudCostManagementTagMappingUnprocessableEntity() *PutCloudCostManagementTagMappingUnprocessableEntity {
	return &PutCloudCostManagementTagMappingUnprocessableEntity{}
}

/*
PutCloudCostManagementTagMappingUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type PutCloudCostManagementTagMappingUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this put cloud cost management tag mapping unprocessable entity response has a 2xx status code
func (o *PutCloudCostManagementTagMappingUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put cloud cost management tag mapping unprocessable entity response has a 3xx status code
func (o *PutCloudCostManagementTagMappingUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put cloud cost management tag mapping unprocessable entity response has a 4xx status code
func (o *PutCloudCostManagementTagMappingUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this put cloud cost management tag mapping unprocessable entity response has a 5xx status code
func (o *PutCloudCostManagementTagMappingUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this put cloud cost management tag mapping unprocessable entity response a status code equal to that given
func (o *PutCloudCostManagementTagMappingUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the put cloud cost management tag mapping unprocessable entity response
func (o *PutCloudCostManagementTagMappingUnprocessableEntity) Code() int {
	return 422
}

func (o *PutCloudCostManagementTagMappingUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMappingUnprocessableEntity %s", 422, payload)
}

func (o *PutCloudCostManagementTagMappingUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMappingUnprocessableEntity %s", 422, payload)
}

func (o *PutCloudCostManagementTagMappingUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PutCloudCostManagementTagMappingUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPutCloudCostManagementTagMappingDefault creates a PutCloudCostManagementTagMappingDefault with default headers values
func NewPutCloudCostManagementTagMappingDefault(code int) *PutCloudCostManagementTagMappingDefault {
	return &PutCloudCostManagementTagMappingDefault{
		_statusCode: code,
	}
}

/*
PutCloudCostManagementTagMappingDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type PutCloudCostManagementTagMappingDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this put cloud cost management tag mapping default response has a 2xx status code
func (o *PutCloudCostManagementTagMappingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put cloud cost management tag mapping default response has a 3xx status code
func (o *PutCloudCostManagementTagMappingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put cloud cost management tag mapping default response has a 4xx status code
func (o *PutCloudCostManagementTagMappingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put cloud cost management tag mapping default response has a 5xx status code
func (o *PutCloudCostManagementTagMappingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put cloud cost management tag mapping default response a status code equal to that given
func (o *PutCloudCostManagementTagMappingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put cloud cost management tag mapping default response
func (o *PutCloudCostManagementTagMappingDefault) Code() int {
	return o._statusCode
}

func (o *PutCloudCostManagementTagMappingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMapping default %s", o._statusCode, payload)
}

func (o *PutCloudCostManagementTagMappingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] putCloudCostManagementTagMapping default %s", o._statusCode, payload)
}

func (o *PutCloudCostManagementTagMappingDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *PutCloudCostManagementTagMappingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
PutCloudCostManagementTagMappingOKBody put cloud cost management tag mapping o k body
swagger:model PutCloudCostManagementTagMappingOKBody
*/
type PutCloudCostManagementTagMappingOKBody struct {

	// data
	// Required: true
	Data *models.CloudCostManagementTagMapping `json:"data"`
}

// Validate validates this put cloud cost management tag mapping o k body
func (o *PutCloudCostManagementTagMappingOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutCloudCostManagementTagMappingOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("putCloudCostManagementTagMappingOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCloudCostManagementTagMappingOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putCloudCostManagementTagMappingOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put cloud cost management tag mapping o k body based on the context it is used
func (o *PutCloudCostManagementTagMappingOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutCloudCostManagementTagMappingOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putCloudCostManagementTagMappingOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putCloudCostManagementTagMappingOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCloudCostManagementTagMappingOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCloudCostManagementTagMappingOKBody) UnmarshalBinary(b []byte) error {
	var res PutCloudCostManagementTagMappingOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

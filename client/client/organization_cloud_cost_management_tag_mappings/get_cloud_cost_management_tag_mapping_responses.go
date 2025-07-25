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

// GetCloudCostManagementTagMappingReader is a Reader for the GetCloudCostManagementTagMapping structure.
type GetCloudCostManagementTagMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudCostManagementTagMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudCostManagementTagMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetCloudCostManagementTagMappingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudCostManagementTagMappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetCloudCostManagementTagMappingUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCloudCostManagementTagMappingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudCostManagementTagMappingOK creates a GetCloudCostManagementTagMappingOK with default headers values
func NewGetCloudCostManagementTagMappingOK() *GetCloudCostManagementTagMappingOK {
	return &GetCloudCostManagementTagMappingOK{}
}

/*
GetCloudCostManagementTagMappingOK describes a response with status code 200, with default header values.

The CloudCostManagementTagMapping for the organization
*/
type GetCloudCostManagementTagMappingOK struct {
	Payload *GetCloudCostManagementTagMappingOKBody
}

// IsSuccess returns true when this get cloud cost management tag mapping o k response has a 2xx status code
func (o *GetCloudCostManagementTagMappingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cloud cost management tag mapping o k response has a 3xx status code
func (o *GetCloudCostManagementTagMappingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tag mapping o k response has a 4xx status code
func (o *GetCloudCostManagementTagMappingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud cost management tag mapping o k response has a 5xx status code
func (o *GetCloudCostManagementTagMappingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tag mapping o k response a status code equal to that given
func (o *GetCloudCostManagementTagMappingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cloud cost management tag mapping o k response
func (o *GetCloudCostManagementTagMappingOK) Code() int {
	return 200
}

func (o *GetCloudCostManagementTagMappingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingOK %s", 200, payload)
}

func (o *GetCloudCostManagementTagMappingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingOK %s", 200, payload)
}

func (o *GetCloudCostManagementTagMappingOK) GetPayload() *GetCloudCostManagementTagMappingOKBody {
	return o.Payload
}

func (o *GetCloudCostManagementTagMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCloudCostManagementTagMappingOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudCostManagementTagMappingForbidden creates a GetCloudCostManagementTagMappingForbidden with default headers values
func NewGetCloudCostManagementTagMappingForbidden() *GetCloudCostManagementTagMappingForbidden {
	return &GetCloudCostManagementTagMappingForbidden{}
}

/*
GetCloudCostManagementTagMappingForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetCloudCostManagementTagMappingForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tag mapping forbidden response has a 2xx status code
func (o *GetCloudCostManagementTagMappingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud cost management tag mapping forbidden response has a 3xx status code
func (o *GetCloudCostManagementTagMappingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tag mapping forbidden response has a 4xx status code
func (o *GetCloudCostManagementTagMappingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud cost management tag mapping forbidden response has a 5xx status code
func (o *GetCloudCostManagementTagMappingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tag mapping forbidden response a status code equal to that given
func (o *GetCloudCostManagementTagMappingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cloud cost management tag mapping forbidden response
func (o *GetCloudCostManagementTagMappingForbidden) Code() int {
	return 403
}

func (o *GetCloudCostManagementTagMappingForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingForbidden %s", 403, payload)
}

func (o *GetCloudCostManagementTagMappingForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingForbidden %s", 403, payload)
}

func (o *GetCloudCostManagementTagMappingForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagMappingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudCostManagementTagMappingNotFound creates a GetCloudCostManagementTagMappingNotFound with default headers values
func NewGetCloudCostManagementTagMappingNotFound() *GetCloudCostManagementTagMappingNotFound {
	return &GetCloudCostManagementTagMappingNotFound{}
}

/*
GetCloudCostManagementTagMappingNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetCloudCostManagementTagMappingNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tag mapping not found response has a 2xx status code
func (o *GetCloudCostManagementTagMappingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud cost management tag mapping not found response has a 3xx status code
func (o *GetCloudCostManagementTagMappingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tag mapping not found response has a 4xx status code
func (o *GetCloudCostManagementTagMappingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud cost management tag mapping not found response has a 5xx status code
func (o *GetCloudCostManagementTagMappingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tag mapping not found response a status code equal to that given
func (o *GetCloudCostManagementTagMappingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cloud cost management tag mapping not found response
func (o *GetCloudCostManagementTagMappingNotFound) Code() int {
	return 404
}

func (o *GetCloudCostManagementTagMappingNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingNotFound %s", 404, payload)
}

func (o *GetCloudCostManagementTagMappingNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingNotFound %s", 404, payload)
}

func (o *GetCloudCostManagementTagMappingNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagMappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudCostManagementTagMappingUnprocessableEntity creates a GetCloudCostManagementTagMappingUnprocessableEntity with default headers values
func NewGetCloudCostManagementTagMappingUnprocessableEntity() *GetCloudCostManagementTagMappingUnprocessableEntity {
	return &GetCloudCostManagementTagMappingUnprocessableEntity{}
}

/*
GetCloudCostManagementTagMappingUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetCloudCostManagementTagMappingUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tag mapping unprocessable entity response has a 2xx status code
func (o *GetCloudCostManagementTagMappingUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud cost management tag mapping unprocessable entity response has a 3xx status code
func (o *GetCloudCostManagementTagMappingUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud cost management tag mapping unprocessable entity response has a 4xx status code
func (o *GetCloudCostManagementTagMappingUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud cost management tag mapping unprocessable entity response has a 5xx status code
func (o *GetCloudCostManagementTagMappingUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud cost management tag mapping unprocessable entity response a status code equal to that given
func (o *GetCloudCostManagementTagMappingUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get cloud cost management tag mapping unprocessable entity response
func (o *GetCloudCostManagementTagMappingUnprocessableEntity) Code() int {
	return 422
}

func (o *GetCloudCostManagementTagMappingUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingUnprocessableEntity %s", 422, payload)
}

func (o *GetCloudCostManagementTagMappingUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMappingUnprocessableEntity %s", 422, payload)
}

func (o *GetCloudCostManagementTagMappingUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagMappingUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudCostManagementTagMappingDefault creates a GetCloudCostManagementTagMappingDefault with default headers values
func NewGetCloudCostManagementTagMappingDefault(code int) *GetCloudCostManagementTagMappingDefault {
	return &GetCloudCostManagementTagMappingDefault{
		_statusCode: code,
	}
}

/*
GetCloudCostManagementTagMappingDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetCloudCostManagementTagMappingDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get cloud cost management tag mapping default response has a 2xx status code
func (o *GetCloudCostManagementTagMappingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cloud cost management tag mapping default response has a 3xx status code
func (o *GetCloudCostManagementTagMappingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cloud cost management tag mapping default response has a 4xx status code
func (o *GetCloudCostManagementTagMappingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cloud cost management tag mapping default response has a 5xx status code
func (o *GetCloudCostManagementTagMappingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cloud cost management tag mapping default response a status code equal to that given
func (o *GetCloudCostManagementTagMappingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cloud cost management tag mapping default response
func (o *GetCloudCostManagementTagMappingDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudCostManagementTagMappingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMapping default %s", o._statusCode, payload)
}

func (o *GetCloudCostManagementTagMappingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/cloud_cost_management/tag_mapping][%d] getCloudCostManagementTagMapping default %s", o._statusCode, payload)
}

func (o *GetCloudCostManagementTagMappingDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCloudCostManagementTagMappingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetCloudCostManagementTagMappingOKBody get cloud cost management tag mapping o k body
swagger:model GetCloudCostManagementTagMappingOKBody
*/
type GetCloudCostManagementTagMappingOKBody struct {

	// data
	// Required: true
	Data *models.CloudCostManagementTagMapping `json:"data"`
}

// Validate validates this get cloud cost management tag mapping o k body
func (o *GetCloudCostManagementTagMappingOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCloudCostManagementTagMappingOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getCloudCostManagementTagMappingOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCloudCostManagementTagMappingOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCloudCostManagementTagMappingOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get cloud cost management tag mapping o k body based on the context it is used
func (o *GetCloudCostManagementTagMappingOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCloudCostManagementTagMappingOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCloudCostManagementTagMappingOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getCloudCostManagementTagMappingOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCloudCostManagementTagMappingOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCloudCostManagementTagMappingOKBody) UnmarshalBinary(b []byte) error {
	var res GetCloudCostManagementTagMappingOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

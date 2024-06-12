// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

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

// GetServiceCatalogTerraformDiagramReader is a Reader for the GetServiceCatalogTerraformDiagram structure.
type GetServiceCatalogTerraformDiagramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogTerraformDiagramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogTerraformDiagramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetServiceCatalogTerraformDiagramForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceCatalogTerraformDiagramNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetServiceCatalogTerraformDiagramDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceCatalogTerraformDiagramOK creates a GetServiceCatalogTerraformDiagramOK with default headers values
func NewGetServiceCatalogTerraformDiagramOK() *GetServiceCatalogTerraformDiagramOK {
	return &GetServiceCatalogTerraformDiagramOK{}
}

/*
GetServiceCatalogTerraformDiagramOK describes a response with status code 200, with default header values.

The information of Terraform Diagram
*/
type GetServiceCatalogTerraformDiagramOK struct {
	Payload *GetServiceCatalogTerraformDiagramOKBody
}

// IsSuccess returns true when this get service catalog terraform diagram o k response has a 2xx status code
func (o *GetServiceCatalogTerraformDiagramOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service catalog terraform diagram o k response has a 3xx status code
func (o *GetServiceCatalogTerraformDiagramOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service catalog terraform diagram o k response has a 4xx status code
func (o *GetServiceCatalogTerraformDiagramOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service catalog terraform diagram o k response has a 5xx status code
func (o *GetServiceCatalogTerraformDiagramOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get service catalog terraform diagram o k response a status code equal to that given
func (o *GetServiceCatalogTerraformDiagramOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get service catalog terraform diagram o k response
func (o *GetServiceCatalogTerraformDiagramOK) Code() int {
	return 200
}

func (o *GetServiceCatalogTerraformDiagramOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramOK %s", 200, payload)
}

func (o *GetServiceCatalogTerraformDiagramOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramOK %s", 200, payload)
}

func (o *GetServiceCatalogTerraformDiagramOK) GetPayload() *GetServiceCatalogTerraformDiagramOKBody {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceCatalogTerraformDiagramOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogTerraformDiagramForbidden creates a GetServiceCatalogTerraformDiagramForbidden with default headers values
func NewGetServiceCatalogTerraformDiagramForbidden() *GetServiceCatalogTerraformDiagramForbidden {
	return &GetServiceCatalogTerraformDiagramForbidden{}
}

/*
GetServiceCatalogTerraformDiagramForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetServiceCatalogTerraformDiagramForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get service catalog terraform diagram forbidden response has a 2xx status code
func (o *GetServiceCatalogTerraformDiagramForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service catalog terraform diagram forbidden response has a 3xx status code
func (o *GetServiceCatalogTerraformDiagramForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service catalog terraform diagram forbidden response has a 4xx status code
func (o *GetServiceCatalogTerraformDiagramForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service catalog terraform diagram forbidden response has a 5xx status code
func (o *GetServiceCatalogTerraformDiagramForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get service catalog terraform diagram forbidden response a status code equal to that given
func (o *GetServiceCatalogTerraformDiagramForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get service catalog terraform diagram forbidden response
func (o *GetServiceCatalogTerraformDiagramForbidden) Code() int {
	return 403
}

func (o *GetServiceCatalogTerraformDiagramForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramForbidden %s", 403, payload)
}

func (o *GetServiceCatalogTerraformDiagramForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramForbidden %s", 403, payload)
}

func (o *GetServiceCatalogTerraformDiagramForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogTerraformDiagramNotFound creates a GetServiceCatalogTerraformDiagramNotFound with default headers values
func NewGetServiceCatalogTerraformDiagramNotFound() *GetServiceCatalogTerraformDiagramNotFound {
	return &GetServiceCatalogTerraformDiagramNotFound{}
}

/*
GetServiceCatalogTerraformDiagramNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetServiceCatalogTerraformDiagramNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get service catalog terraform diagram not found response has a 2xx status code
func (o *GetServiceCatalogTerraformDiagramNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get service catalog terraform diagram not found response has a 3xx status code
func (o *GetServiceCatalogTerraformDiagramNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service catalog terraform diagram not found response has a 4xx status code
func (o *GetServiceCatalogTerraformDiagramNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get service catalog terraform diagram not found response has a 5xx status code
func (o *GetServiceCatalogTerraformDiagramNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get service catalog terraform diagram not found response a status code equal to that given
func (o *GetServiceCatalogTerraformDiagramNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get service catalog terraform diagram not found response
func (o *GetServiceCatalogTerraformDiagramNotFound) Code() int {
	return 404
}

func (o *GetServiceCatalogTerraformDiagramNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramNotFound %s", 404, payload)
}

func (o *GetServiceCatalogTerraformDiagramNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagramNotFound %s", 404, payload)
}

func (o *GetServiceCatalogTerraformDiagramNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetServiceCatalogTerraformDiagramDefault creates a GetServiceCatalogTerraformDiagramDefault with default headers values
func NewGetServiceCatalogTerraformDiagramDefault(code int) *GetServiceCatalogTerraformDiagramDefault {
	return &GetServiceCatalogTerraformDiagramDefault{
		_statusCode: code,
	}
}

/*
GetServiceCatalogTerraformDiagramDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceCatalogTerraformDiagramDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get service catalog terraform diagram default response has a 2xx status code
func (o *GetServiceCatalogTerraformDiagramDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get service catalog terraform diagram default response has a 3xx status code
func (o *GetServiceCatalogTerraformDiagramDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get service catalog terraform diagram default response has a 4xx status code
func (o *GetServiceCatalogTerraformDiagramDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get service catalog terraform diagram default response has a 5xx status code
func (o *GetServiceCatalogTerraformDiagramDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get service catalog terraform diagram default response a status code equal to that given
func (o *GetServiceCatalogTerraformDiagramDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get service catalog terraform diagram default response
func (o *GetServiceCatalogTerraformDiagramDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceCatalogTerraformDiagramDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagram default %s", o._statusCode, payload)
}

func (o *GetServiceCatalogTerraformDiagramDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] getServiceCatalogTerraformDiagram default %s", o._statusCode, payload)
}

func (o *GetServiceCatalogTerraformDiagramDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceCatalogTerraformDiagramDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetServiceCatalogTerraformDiagramOKBody get service catalog terraform diagram o k body
swagger:model GetServiceCatalogTerraformDiagramOKBody
*/
type GetServiceCatalogTerraformDiagramOKBody struct {

	// created at
	// Minimum: 0
	CreatedAt *uint64 `json:"created_at,omitempty"`

	// data
	// Required: true
	Data models.TerraformJSONDiagram `json:"data"`

	// updated at
	// Minimum: 0
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// Validate validates this get service catalog terraform diagram o k body
func (o *GetServiceCatalogTerraformDiagramOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceCatalogTerraformDiagramOKBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("getServiceCatalogTerraformDiagramOK"+"."+"created_at", "body", *o.CreatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

func (o *GetServiceCatalogTerraformDiagramOKBody) validateData(formats strfmt.Registry) error {

	if o.Data == nil {
		return errors.Required("getServiceCatalogTerraformDiagramOK"+"."+"data", "body", nil)
	}

	return nil
}

func (o *GetServiceCatalogTerraformDiagramOKBody) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.MinimumUint("getServiceCatalogTerraformDiagramOK"+"."+"updated_at", "body", *o.UpdatedAt, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get service catalog terraform diagram o k body based on context it is used
func (o *GetServiceCatalogTerraformDiagramOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceCatalogTerraformDiagramOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceCatalogTerraformDiagramOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceCatalogTerraformDiagramOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// UpdateServiceCatalogTerraformDiagramReader is a Reader for the UpdateServiceCatalogTerraformDiagram structure.
type UpdateServiceCatalogTerraformDiagramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceCatalogTerraformDiagramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateServiceCatalogTerraformDiagramNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateServiceCatalogTerraformDiagramForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServiceCatalogTerraformDiagramNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateServiceCatalogTerraformDiagramUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateServiceCatalogTerraformDiagramDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateServiceCatalogTerraformDiagramNoContent creates a UpdateServiceCatalogTerraformDiagramNoContent with default headers values
func NewUpdateServiceCatalogTerraformDiagramNoContent() *UpdateServiceCatalogTerraformDiagramNoContent {
	return &UpdateServiceCatalogTerraformDiagramNoContent{}
}

/*
UpdateServiceCatalogTerraformDiagramNoContent describes a response with status code 204, with default header values.

Configuration has been updated
*/
type UpdateServiceCatalogTerraformDiagramNoContent struct {
}

// IsSuccess returns true when this update service catalog terraform diagram no content response has a 2xx status code
func (o *UpdateServiceCatalogTerraformDiagramNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update service catalog terraform diagram no content response has a 3xx status code
func (o *UpdateServiceCatalogTerraformDiagramNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service catalog terraform diagram no content response has a 4xx status code
func (o *UpdateServiceCatalogTerraformDiagramNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update service catalog terraform diagram no content response has a 5xx status code
func (o *UpdateServiceCatalogTerraformDiagramNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update service catalog terraform diagram no content response a status code equal to that given
func (o *UpdateServiceCatalogTerraformDiagramNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the update service catalog terraform diagram no content response
func (o *UpdateServiceCatalogTerraformDiagramNoContent) Code() int {
	return 204
}

func (o *UpdateServiceCatalogTerraformDiagramNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramNoContent", 204)
}

func (o *UpdateServiceCatalogTerraformDiagramNoContent) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramNoContent", 204)
}

func (o *UpdateServiceCatalogTerraformDiagramNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceCatalogTerraformDiagramForbidden creates a UpdateServiceCatalogTerraformDiagramForbidden with default headers values
func NewUpdateServiceCatalogTerraformDiagramForbidden() *UpdateServiceCatalogTerraformDiagramForbidden {
	return &UpdateServiceCatalogTerraformDiagramForbidden{}
}

/*
UpdateServiceCatalogTerraformDiagramForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateServiceCatalogTerraformDiagramForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update service catalog terraform diagram forbidden response has a 2xx status code
func (o *UpdateServiceCatalogTerraformDiagramForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service catalog terraform diagram forbidden response has a 3xx status code
func (o *UpdateServiceCatalogTerraformDiagramForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service catalog terraform diagram forbidden response has a 4xx status code
func (o *UpdateServiceCatalogTerraformDiagramForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service catalog terraform diagram forbidden response has a 5xx status code
func (o *UpdateServiceCatalogTerraformDiagramForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update service catalog terraform diagram forbidden response a status code equal to that given
func (o *UpdateServiceCatalogTerraformDiagramForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update service catalog terraform diagram forbidden response
func (o *UpdateServiceCatalogTerraformDiagramForbidden) Code() int {
	return 403
}

func (o *UpdateServiceCatalogTerraformDiagramForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramForbidden %s", 403, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramForbidden %s", 403, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateServiceCatalogTerraformDiagramNotFound creates a UpdateServiceCatalogTerraformDiagramNotFound with default headers values
func NewUpdateServiceCatalogTerraformDiagramNotFound() *UpdateServiceCatalogTerraformDiagramNotFound {
	return &UpdateServiceCatalogTerraformDiagramNotFound{}
}

/*
UpdateServiceCatalogTerraformDiagramNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateServiceCatalogTerraformDiagramNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update service catalog terraform diagram not found response has a 2xx status code
func (o *UpdateServiceCatalogTerraformDiagramNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service catalog terraform diagram not found response has a 3xx status code
func (o *UpdateServiceCatalogTerraformDiagramNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service catalog terraform diagram not found response has a 4xx status code
func (o *UpdateServiceCatalogTerraformDiagramNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service catalog terraform diagram not found response has a 5xx status code
func (o *UpdateServiceCatalogTerraformDiagramNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update service catalog terraform diagram not found response a status code equal to that given
func (o *UpdateServiceCatalogTerraformDiagramNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update service catalog terraform diagram not found response
func (o *UpdateServiceCatalogTerraformDiagramNotFound) Code() int {
	return 404
}

func (o *UpdateServiceCatalogTerraformDiagramNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramNotFound %s", 404, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramNotFound %s", 404, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateServiceCatalogTerraformDiagramUnprocessableEntity creates a UpdateServiceCatalogTerraformDiagramUnprocessableEntity with default headers values
func NewUpdateServiceCatalogTerraformDiagramUnprocessableEntity() *UpdateServiceCatalogTerraformDiagramUnprocessableEntity {
	return &UpdateServiceCatalogTerraformDiagramUnprocessableEntity{}
}

/*
UpdateServiceCatalogTerraformDiagramUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateServiceCatalogTerraformDiagramUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update service catalog terraform diagram unprocessable entity response has a 2xx status code
func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service catalog terraform diagram unprocessable entity response has a 3xx status code
func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service catalog terraform diagram unprocessable entity response has a 4xx status code
func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service catalog terraform diagram unprocessable entity response has a 5xx status code
func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update service catalog terraform diagram unprocessable entity response a status code equal to that given
func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update service catalog terraform diagram unprocessable entity response
func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramUnprocessableEntity %s", 422, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramUnprocessableEntity %s", 422, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateServiceCatalogTerraformDiagramDefault creates a UpdateServiceCatalogTerraformDiagramDefault with default headers values
func NewUpdateServiceCatalogTerraformDiagramDefault(code int) *UpdateServiceCatalogTerraformDiagramDefault {
	return &UpdateServiceCatalogTerraformDiagramDefault{
		_statusCode: code,
	}
}

/*
UpdateServiceCatalogTerraformDiagramDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateServiceCatalogTerraformDiagramDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this update service catalog terraform diagram default response has a 2xx status code
func (o *UpdateServiceCatalogTerraformDiagramDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update service catalog terraform diagram default response has a 3xx status code
func (o *UpdateServiceCatalogTerraformDiagramDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update service catalog terraform diagram default response has a 4xx status code
func (o *UpdateServiceCatalogTerraformDiagramDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update service catalog terraform diagram default response has a 5xx status code
func (o *UpdateServiceCatalogTerraformDiagramDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update service catalog terraform diagram default response a status code equal to that given
func (o *UpdateServiceCatalogTerraformDiagramDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update service catalog terraform diagram default response
func (o *UpdateServiceCatalogTerraformDiagramDefault) Code() int {
	return o._statusCode
}

func (o *UpdateServiceCatalogTerraformDiagramDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagram default %s", o._statusCode, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagram default %s", o._statusCode, payload)
}

func (o *UpdateServiceCatalogTerraformDiagramDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

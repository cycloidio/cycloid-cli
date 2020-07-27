// Code generated by go-swagger; DO NOT EDIT.

package service_catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
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

/*UpdateServiceCatalogTerraformDiagramNoContent handles this case with default header values.

Configuration has been updated
*/
type UpdateServiceCatalogTerraformDiagramNoContent struct {
}

func (o *UpdateServiceCatalogTerraformDiagramNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramNoContent ", 204)
}

func (o *UpdateServiceCatalogTerraformDiagramNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceCatalogTerraformDiagramForbidden creates a UpdateServiceCatalogTerraformDiagramForbidden with default headers values
func NewUpdateServiceCatalogTerraformDiagramForbidden() *UpdateServiceCatalogTerraformDiagramForbidden {
	return &UpdateServiceCatalogTerraformDiagramForbidden{}
}

/*UpdateServiceCatalogTerraformDiagramForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateServiceCatalogTerraformDiagramForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *UpdateServiceCatalogTerraformDiagramForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramForbidden  %+v", 403, o.Payload)
}

func (o *UpdateServiceCatalogTerraformDiagramForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

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

/*UpdateServiceCatalogTerraformDiagramNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateServiceCatalogTerraformDiagramNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *UpdateServiceCatalogTerraformDiagramNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramNotFound  %+v", 404, o.Payload)
}

func (o *UpdateServiceCatalogTerraformDiagramNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

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

/*UpdateServiceCatalogTerraformDiagramUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateServiceCatalogTerraformDiagramUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagramUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*UpdateServiceCatalogTerraformDiagramDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateServiceCatalogTerraformDiagramDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the update service catalog terraform diagram default response
func (o *UpdateServiceCatalogTerraformDiagramDefault) Code() int {
	return o._statusCode
}

func (o *UpdateServiceCatalogTerraformDiagramDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram][%d] updateServiceCatalogTerraformDiagram default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateServiceCatalogTerraformDiagramDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformDiagramDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

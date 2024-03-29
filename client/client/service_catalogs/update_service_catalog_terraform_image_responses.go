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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// UpdateServiceCatalogTerraformImageReader is a Reader for the UpdateServiceCatalogTerraformImage structure.
type UpdateServiceCatalogTerraformImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceCatalogTerraformImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateServiceCatalogTerraformImageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateServiceCatalogTerraformImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateServiceCatalogTerraformImageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateServiceCatalogTerraformImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateServiceCatalogTerraformImageNoContent creates a UpdateServiceCatalogTerraformImageNoContent with default headers values
func NewUpdateServiceCatalogTerraformImageNoContent() *UpdateServiceCatalogTerraformImageNoContent {
	return &UpdateServiceCatalogTerraformImageNoContent{}
}

/*UpdateServiceCatalogTerraformImageNoContent handles this case with default header values.

Configuration has been updated
*/
type UpdateServiceCatalogTerraformImageNoContent struct {
}

func (o *UpdateServiceCatalogTerraformImageNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] updateServiceCatalogTerraformImageNoContent ", 204)
}

func (o *UpdateServiceCatalogTerraformImageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateServiceCatalogTerraformImageForbidden creates a UpdateServiceCatalogTerraformImageForbidden with default headers values
func NewUpdateServiceCatalogTerraformImageForbidden() *UpdateServiceCatalogTerraformImageForbidden {
	return &UpdateServiceCatalogTerraformImageForbidden{}
}

/*UpdateServiceCatalogTerraformImageForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateServiceCatalogTerraformImageForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateServiceCatalogTerraformImageForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] updateServiceCatalogTerraformImageForbidden  %+v", 403, o.Payload)
}

func (o *UpdateServiceCatalogTerraformImageForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceCatalogTerraformImageUnprocessableEntity creates a UpdateServiceCatalogTerraformImageUnprocessableEntity with default headers values
func NewUpdateServiceCatalogTerraformImageUnprocessableEntity() *UpdateServiceCatalogTerraformImageUnprocessableEntity {
	return &UpdateServiceCatalogTerraformImageUnprocessableEntity{}
}

/*UpdateServiceCatalogTerraformImageUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateServiceCatalogTerraformImageUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateServiceCatalogTerraformImageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] updateServiceCatalogTerraformImageUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateServiceCatalogTerraformImageUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformImageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceCatalogTerraformImageDefault creates a UpdateServiceCatalogTerraformImageDefault with default headers values
func NewUpdateServiceCatalogTerraformImageDefault(code int) *UpdateServiceCatalogTerraformImageDefault {
	return &UpdateServiceCatalogTerraformImageDefault{
		_statusCode: code,
	}
}

/*UpdateServiceCatalogTerraformImageDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateServiceCatalogTerraformImageDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the update service catalog terraform image default response
func (o *UpdateServiceCatalogTerraformImageDefault) Code() int {
	return o._statusCode
}

func (o *UpdateServiceCatalogTerraformImageDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/terraform/diagram/image][%d] updateServiceCatalogTerraformImage default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateServiceCatalogTerraformImageDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateServiceCatalogTerraformImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

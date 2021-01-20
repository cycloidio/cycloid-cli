// Code generated by go-swagger; DO NOT EDIT.

package organization_service_catalog_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// ValidateServiceCatalogSourceReader is a Reader for the ValidateServiceCatalogSource structure.
type ValidateServiceCatalogSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateServiceCatalogSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateServiceCatalogSourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewValidateServiceCatalogSourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateServiceCatalogSourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewValidateServiceCatalogSourceLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewValidateServiceCatalogSourceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewValidateServiceCatalogSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateServiceCatalogSourceNoContent creates a ValidateServiceCatalogSourceNoContent with default headers values
func NewValidateServiceCatalogSourceNoContent() *ValidateServiceCatalogSourceNoContent {
	return &ValidateServiceCatalogSourceNoContent{}
}

/*ValidateServiceCatalogSourceNoContent handles this case with default header values.

The SCS has been validated
*/
type ValidateServiceCatalogSourceNoContent struct {
}

func (o *ValidateServiceCatalogSourceNoContent) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/validate][%d] validateServiceCatalogSourceNoContent ", 204)
}

func (o *ValidateServiceCatalogSourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateServiceCatalogSourceForbidden creates a ValidateServiceCatalogSourceForbidden with default headers values
func NewValidateServiceCatalogSourceForbidden() *ValidateServiceCatalogSourceForbidden {
	return &ValidateServiceCatalogSourceForbidden{}
}

/*ValidateServiceCatalogSourceForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type ValidateServiceCatalogSourceForbidden struct {
	Payload *models.ErrorPayload
}

func (o *ValidateServiceCatalogSourceForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/validate][%d] validateServiceCatalogSourceForbidden  %+v", 403, o.Payload)
}

func (o *ValidateServiceCatalogSourceForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateServiceCatalogSourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateServiceCatalogSourceNotFound creates a ValidateServiceCatalogSourceNotFound with default headers values
func NewValidateServiceCatalogSourceNotFound() *ValidateServiceCatalogSourceNotFound {
	return &ValidateServiceCatalogSourceNotFound{}
}

/*ValidateServiceCatalogSourceNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type ValidateServiceCatalogSourceNotFound struct {
	Payload *models.ErrorPayload
}

func (o *ValidateServiceCatalogSourceNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/validate][%d] validateServiceCatalogSourceNotFound  %+v", 404, o.Payload)
}

func (o *ValidateServiceCatalogSourceNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateServiceCatalogSourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateServiceCatalogSourceLengthRequired creates a ValidateServiceCatalogSourceLengthRequired with default headers values
func NewValidateServiceCatalogSourceLengthRequired() *ValidateServiceCatalogSourceLengthRequired {
	return &ValidateServiceCatalogSourceLengthRequired{}
}

/*ValidateServiceCatalogSourceLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type ValidateServiceCatalogSourceLengthRequired struct {
}

func (o *ValidateServiceCatalogSourceLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/validate][%d] validateServiceCatalogSourceLengthRequired ", 411)
}

func (o *ValidateServiceCatalogSourceLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateServiceCatalogSourceUnprocessableEntity creates a ValidateServiceCatalogSourceUnprocessableEntity with default headers values
func NewValidateServiceCatalogSourceUnprocessableEntity() *ValidateServiceCatalogSourceUnprocessableEntity {
	return &ValidateServiceCatalogSourceUnprocessableEntity{}
}

/*ValidateServiceCatalogSourceUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type ValidateServiceCatalogSourceUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *ValidateServiceCatalogSourceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/validate][%d] validateServiceCatalogSourceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ValidateServiceCatalogSourceUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateServiceCatalogSourceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateServiceCatalogSourceDefault creates a ValidateServiceCatalogSourceDefault with default headers values
func NewValidateServiceCatalogSourceDefault(code int) *ValidateServiceCatalogSourceDefault {
	return &ValidateServiceCatalogSourceDefault{
		_statusCode: code,
	}
}

/*ValidateServiceCatalogSourceDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type ValidateServiceCatalogSourceDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the validate service catalog source default response
func (o *ValidateServiceCatalogSourceDefault) Code() int {
	return o._statusCode
}

func (o *ValidateServiceCatalogSourceDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/service_catalog_sources/{service_catalog_source_canonical}/validate][%d] validateServiceCatalogSource default  %+v", o._statusCode, o.Payload)
}

func (o *ValidateServiceCatalogSourceDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ValidateServiceCatalogSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

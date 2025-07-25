// Code generated by go-swagger; DO NOT EDIT.

package organization_infra_imports

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

// DeleteServiceCatalogImportReader is a Reader for the DeleteServiceCatalogImport structure.
type DeleteServiceCatalogImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceCatalogImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServiceCatalogImportNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteServiceCatalogImportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServiceCatalogImportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServiceCatalogImportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServiceCatalogImportNoContent creates a DeleteServiceCatalogImportNoContent with default headers values
func NewDeleteServiceCatalogImportNoContent() *DeleteServiceCatalogImportNoContent {
	return &DeleteServiceCatalogImportNoContent{}
}

/*
DeleteServiceCatalogImportNoContent describes a response with status code 204, with default header values.

The Stack import has been deleted.
*/
type DeleteServiceCatalogImportNoContent struct {
}

// IsSuccess returns true when this delete service catalog import no content response has a 2xx status code
func (o *DeleteServiceCatalogImportNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service catalog import no content response has a 3xx status code
func (o *DeleteServiceCatalogImportNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service catalog import no content response has a 4xx status code
func (o *DeleteServiceCatalogImportNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service catalog import no content response has a 5xx status code
func (o *DeleteServiceCatalogImportNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service catalog import no content response a status code equal to that given
func (o *DeleteServiceCatalogImportNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete service catalog import no content response
func (o *DeleteServiceCatalogImportNoContent) Code() int {
	return 204
}

func (o *DeleteServiceCatalogImportNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImportNoContent", 204)
}

func (o *DeleteServiceCatalogImportNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImportNoContent", 204)
}

func (o *DeleteServiceCatalogImportNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCatalogImportForbidden creates a DeleteServiceCatalogImportForbidden with default headers values
func NewDeleteServiceCatalogImportForbidden() *DeleteServiceCatalogImportForbidden {
	return &DeleteServiceCatalogImportForbidden{}
}

/*
DeleteServiceCatalogImportForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteServiceCatalogImportForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete service catalog import forbidden response has a 2xx status code
func (o *DeleteServiceCatalogImportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service catalog import forbidden response has a 3xx status code
func (o *DeleteServiceCatalogImportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service catalog import forbidden response has a 4xx status code
func (o *DeleteServiceCatalogImportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service catalog import forbidden response has a 5xx status code
func (o *DeleteServiceCatalogImportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service catalog import forbidden response a status code equal to that given
func (o *DeleteServiceCatalogImportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete service catalog import forbidden response
func (o *DeleteServiceCatalogImportForbidden) Code() int {
	return 403
}

func (o *DeleteServiceCatalogImportForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImportForbidden %s", 403, payload)
}

func (o *DeleteServiceCatalogImportForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImportForbidden %s", 403, payload)
}

func (o *DeleteServiceCatalogImportForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteServiceCatalogImportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteServiceCatalogImportNotFound creates a DeleteServiceCatalogImportNotFound with default headers values
func NewDeleteServiceCatalogImportNotFound() *DeleteServiceCatalogImportNotFound {
	return &DeleteServiceCatalogImportNotFound{}
}

/*
DeleteServiceCatalogImportNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteServiceCatalogImportNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete service catalog import not found response has a 2xx status code
func (o *DeleteServiceCatalogImportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service catalog import not found response has a 3xx status code
func (o *DeleteServiceCatalogImportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service catalog import not found response has a 4xx status code
func (o *DeleteServiceCatalogImportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service catalog import not found response has a 5xx status code
func (o *DeleteServiceCatalogImportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service catalog import not found response a status code equal to that given
func (o *DeleteServiceCatalogImportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete service catalog import not found response
func (o *DeleteServiceCatalogImportNotFound) Code() int {
	return 404
}

func (o *DeleteServiceCatalogImportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImportNotFound %s", 404, payload)
}

func (o *DeleteServiceCatalogImportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImportNotFound %s", 404, payload)
}

func (o *DeleteServiceCatalogImportNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteServiceCatalogImportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteServiceCatalogImportDefault creates a DeleteServiceCatalogImportDefault with default headers values
func NewDeleteServiceCatalogImportDefault(code int) *DeleteServiceCatalogImportDefault {
	return &DeleteServiceCatalogImportDefault{
		_statusCode: code,
	}
}

/*
DeleteServiceCatalogImportDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteServiceCatalogImportDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete service catalog import default response has a 2xx status code
func (o *DeleteServiceCatalogImportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete service catalog import default response has a 3xx status code
func (o *DeleteServiceCatalogImportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete service catalog import default response has a 4xx status code
func (o *DeleteServiceCatalogImportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete service catalog import default response has a 5xx status code
func (o *DeleteServiceCatalogImportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete service catalog import default response a status code equal to that given
func (o *DeleteServiceCatalogImportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete service catalog import default response
func (o *DeleteServiceCatalogImportDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServiceCatalogImportDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImport default %s", o._statusCode, payload)
}

func (o *DeleteServiceCatalogImportDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/service_catalogs/{service_catalog_ref}/import][%d] deleteServiceCatalogImport default %s", o._statusCode, payload)
}

func (o *DeleteServiceCatalogImportDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteServiceCatalogImportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

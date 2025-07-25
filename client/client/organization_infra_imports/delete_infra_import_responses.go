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

// DeleteInfraImportReader is a Reader for the DeleteInfraImport structure.
type DeleteInfraImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInfraImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteInfraImportNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteInfraImportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInfraImportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteInfraImportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteInfraImportNoContent creates a DeleteInfraImportNoContent with default headers values
func NewDeleteInfraImportNoContent() *DeleteInfraImportNoContent {
	return &DeleteInfraImportNoContent{}
}

/*
DeleteInfraImportNoContent describes a response with status code 204, with default header values.

The Import has been deleted.
*/
type DeleteInfraImportNoContent struct {
}

// IsSuccess returns true when this delete infra import no content response has a 2xx status code
func (o *DeleteInfraImportNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete infra import no content response has a 3xx status code
func (o *DeleteInfraImportNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete infra import no content response has a 4xx status code
func (o *DeleteInfraImportNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete infra import no content response has a 5xx status code
func (o *DeleteInfraImportNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete infra import no content response a status code equal to that given
func (o *DeleteInfraImportNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete infra import no content response
func (o *DeleteInfraImportNoContent) Code() int {
	return 204
}

func (o *DeleteInfraImportNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImportNoContent", 204)
}

func (o *DeleteInfraImportNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImportNoContent", 204)
}

func (o *DeleteInfraImportNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInfraImportForbidden creates a DeleteInfraImportForbidden with default headers values
func NewDeleteInfraImportForbidden() *DeleteInfraImportForbidden {
	return &DeleteInfraImportForbidden{}
}

/*
DeleteInfraImportForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteInfraImportForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete infra import forbidden response has a 2xx status code
func (o *DeleteInfraImportForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete infra import forbidden response has a 3xx status code
func (o *DeleteInfraImportForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete infra import forbidden response has a 4xx status code
func (o *DeleteInfraImportForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete infra import forbidden response has a 5xx status code
func (o *DeleteInfraImportForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete infra import forbidden response a status code equal to that given
func (o *DeleteInfraImportForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete infra import forbidden response
func (o *DeleteInfraImportForbidden) Code() int {
	return 403
}

func (o *DeleteInfraImportForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImportForbidden %s", 403, payload)
}

func (o *DeleteInfraImportForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImportForbidden %s", 403, payload)
}

func (o *DeleteInfraImportForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInfraImportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInfraImportNotFound creates a DeleteInfraImportNotFound with default headers values
func NewDeleteInfraImportNotFound() *DeleteInfraImportNotFound {
	return &DeleteInfraImportNotFound{}
}

/*
DeleteInfraImportNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteInfraImportNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete infra import not found response has a 2xx status code
func (o *DeleteInfraImportNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete infra import not found response has a 3xx status code
func (o *DeleteInfraImportNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete infra import not found response has a 4xx status code
func (o *DeleteInfraImportNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete infra import not found response has a 5xx status code
func (o *DeleteInfraImportNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete infra import not found response a status code equal to that given
func (o *DeleteInfraImportNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete infra import not found response
func (o *DeleteInfraImportNotFound) Code() int {
	return 404
}

func (o *DeleteInfraImportNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImportNotFound %s", 404, payload)
}

func (o *DeleteInfraImportNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImportNotFound %s", 404, payload)
}

func (o *DeleteInfraImportNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInfraImportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInfraImportDefault creates a DeleteInfraImportDefault with default headers values
func NewDeleteInfraImportDefault(code int) *DeleteInfraImportDefault {
	return &DeleteInfraImportDefault{
		_statusCode: code,
	}
}

/*
DeleteInfraImportDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteInfraImportDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete infra import default response has a 2xx status code
func (o *DeleteInfraImportDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete infra import default response has a 3xx status code
func (o *DeleteInfraImportDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete infra import default response has a 4xx status code
func (o *DeleteInfraImportDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete infra import default response has a 5xx status code
func (o *DeleteInfraImportDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete infra import default response a status code equal to that given
func (o *DeleteInfraImportDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete infra import default response
func (o *DeleteInfraImportDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInfraImportDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImport default %s", o._statusCode, payload)
}

func (o *DeleteInfraImportDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/import][%d] deleteInfraImport default %s", o._statusCode, payload)
}

func (o *DeleteInfraImportDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInfraImportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

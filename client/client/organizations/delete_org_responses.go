// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// DeleteOrgReader is a Reader for the DeleteOrg structure.
type DeleteOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrgNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrgNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteOrgDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrgNoContent creates a DeleteOrgNoContent with default headers values
func NewDeleteOrgNoContent() *DeleteOrgNoContent {
	return &DeleteOrgNoContent{}
}

/*
DeleteOrgNoContent describes a response with status code 204, with default header values.

Organization has been deleted.
*/
type DeleteOrgNoContent struct {
}

// IsSuccess returns true when this delete org no content response has a 2xx status code
func (o *DeleteOrgNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete org no content response has a 3xx status code
func (o *DeleteOrgNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org no content response has a 4xx status code
func (o *DeleteOrgNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete org no content response has a 5xx status code
func (o *DeleteOrgNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org no content response a status code equal to that given
func (o *DeleteOrgNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete org no content response
func (o *DeleteOrgNoContent) Code() int {
	return 204
}

func (o *DeleteOrgNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrgNoContent", 204)
}

func (o *DeleteOrgNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrgNoContent", 204)
}

func (o *DeleteOrgNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgForbidden creates a DeleteOrgForbidden with default headers values
func NewDeleteOrgForbidden() *DeleteOrgForbidden {
	return &DeleteOrgForbidden{}
}

/*
DeleteOrgForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteOrgForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete org forbidden response has a 2xx status code
func (o *DeleteOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete org forbidden response has a 3xx status code
func (o *DeleteOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org forbidden response has a 4xx status code
func (o *DeleteOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete org forbidden response has a 5xx status code
func (o *DeleteOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org forbidden response a status code equal to that given
func (o *DeleteOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete org forbidden response
func (o *DeleteOrgForbidden) Code() int {
	return 403
}

func (o *DeleteOrgForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrgForbidden %s", 403, payload)
}

func (o *DeleteOrgForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrgForbidden %s", 403, payload)
}

func (o *DeleteOrgForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteOrgNotFound creates a DeleteOrgNotFound with default headers values
func NewDeleteOrgNotFound() *DeleteOrgNotFound {
	return &DeleteOrgNotFound{}
}

/*
DeleteOrgNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteOrgNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete org not found response has a 2xx status code
func (o *DeleteOrgNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete org not found response has a 3xx status code
func (o *DeleteOrgNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org not found response has a 4xx status code
func (o *DeleteOrgNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete org not found response has a 5xx status code
func (o *DeleteOrgNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org not found response a status code equal to that given
func (o *DeleteOrgNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete org not found response
func (o *DeleteOrgNotFound) Code() int {
	return 404
}

func (o *DeleteOrgNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrgNotFound %s", 404, payload)
}

func (o *DeleteOrgNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrgNotFound %s", 404, payload)
}

func (o *DeleteOrgNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteOrgNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteOrgDefault creates a DeleteOrgDefault with default headers values
func NewDeleteOrgDefault(code int) *DeleteOrgDefault {
	return &DeleteOrgDefault{
		_statusCode: code,
	}
}

/*
DeleteOrgDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteOrgDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete org default response has a 2xx status code
func (o *DeleteOrgDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete org default response has a 3xx status code
func (o *DeleteOrgDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete org default response has a 4xx status code
func (o *DeleteOrgDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete org default response has a 5xx status code
func (o *DeleteOrgDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete org default response a status code equal to that given
func (o *DeleteOrgDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete org default response
func (o *DeleteOrgDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrgDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrg default %s", o._statusCode, payload)
}

func (o *DeleteOrgDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}][%d] deleteOrg default %s", o._statusCode, payload)
}

func (o *DeleteOrgDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteOrgDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

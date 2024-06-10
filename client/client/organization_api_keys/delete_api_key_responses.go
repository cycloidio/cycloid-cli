// Code generated by go-swagger; DO NOT EDIT.

package organization_api_keys

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

// DeleteAPIKeyReader is a Reader for the DeleteAPIKey structure.
type DeleteAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPIKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteAPIKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAPIKeyNoContent creates a DeleteAPIKeyNoContent with default headers values
func NewDeleteAPIKeyNoContent() *DeleteAPIKeyNoContent {
	return &DeleteAPIKeyNoContent{}
}

/*
DeleteAPIKeyNoContent describes a response with status code 204, with default header values.

API key has been deleted.
*/
type DeleteAPIKeyNoContent struct {
}

// IsSuccess returns true when this delete Api key no content response has a 2xx status code
func (o *DeleteAPIKeyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api key no content response has a 3xx status code
func (o *DeleteAPIKeyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key no content response has a 4xx status code
func (o *DeleteAPIKeyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api key no content response has a 5xx status code
func (o *DeleteAPIKeyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api key no content response a status code equal to that given
func (o *DeleteAPIKeyNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Api key no content response
func (o *DeleteAPIKeyNoContent) Code() int {
	return 204
}

func (o *DeleteAPIKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyNoContent", 204)
}

func (o *DeleteAPIKeyNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyNoContent", 204)
}

func (o *DeleteAPIKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIKeyForbidden creates a DeleteAPIKeyForbidden with default headers values
func NewDeleteAPIKeyForbidden() *DeleteAPIKeyForbidden {
	return &DeleteAPIKeyForbidden{}
}

/*
DeleteAPIKeyForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteAPIKeyForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete Api key forbidden response has a 2xx status code
func (o *DeleteAPIKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api key forbidden response has a 3xx status code
func (o *DeleteAPIKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key forbidden response has a 4xx status code
func (o *DeleteAPIKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api key forbidden response has a 5xx status code
func (o *DeleteAPIKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api key forbidden response a status code equal to that given
func (o *DeleteAPIKeyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete Api key forbidden response
func (o *DeleteAPIKeyForbidden) Code() int {
	return 403
}

func (o *DeleteAPIKeyForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyForbidden %s", 403, payload)
}

func (o *DeleteAPIKeyForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyForbidden %s", 403, payload)
}

func (o *DeleteAPIKeyForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAPIKeyNotFound creates a DeleteAPIKeyNotFound with default headers values
func NewDeleteAPIKeyNotFound() *DeleteAPIKeyNotFound {
	return &DeleteAPIKeyNotFound{}
}

/*
DeleteAPIKeyNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteAPIKeyNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete Api key not found response has a 2xx status code
func (o *DeleteAPIKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete Api key not found response has a 3xx status code
func (o *DeleteAPIKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api key not found response has a 4xx status code
func (o *DeleteAPIKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete Api key not found response has a 5xx status code
func (o *DeleteAPIKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api key not found response a status code equal to that given
func (o *DeleteAPIKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete Api key not found response
func (o *DeleteAPIKeyNotFound) Code() int {
	return 404
}

func (o *DeleteAPIKeyNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyNotFound %s", 404, payload)
}

func (o *DeleteAPIKeyNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyNotFound %s", 404, payload)
}

func (o *DeleteAPIKeyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAPIKeyDefault creates a DeleteAPIKeyDefault with default headers values
func NewDeleteAPIKeyDefault(code int) *DeleteAPIKeyDefault {
	return &DeleteAPIKeyDefault{
		_statusCode: code,
	}
}

/*
DeleteAPIKeyDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteAPIKeyDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete API key default response has a 2xx status code
func (o *DeleteAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete API key default response has a 3xx status code
func (o *DeleteAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete API key default response has a 4xx status code
func (o *DeleteAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete API key default response has a 5xx status code
func (o *DeleteAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete API key default response a status code equal to that given
func (o *DeleteAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete API key default response
func (o *DeleteAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAPIKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteAPIKey default %s", o._statusCode, payload)
}

func (o *DeleteAPIKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteAPIKey default %s", o._statusCode, payload)
}

func (o *DeleteAPIKeyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

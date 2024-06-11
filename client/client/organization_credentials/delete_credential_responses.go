// Code generated by go-swagger; DO NOT EDIT.

package organization_credentials

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

// DeleteCredentialReader is a Reader for the DeleteCredential structure.
type DeleteCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCredentialNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteCredentialConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCredentialNoContent creates a DeleteCredentialNoContent with default headers values
func NewDeleteCredentialNoContent() *DeleteCredentialNoContent {
	return &DeleteCredentialNoContent{}
}

/*
DeleteCredentialNoContent describes a response with status code 204, with default header values.

Credential has been deleted.
*/
type DeleteCredentialNoContent struct {
}

// IsSuccess returns true when this delete credential no content response has a 2xx status code
func (o *DeleteCredentialNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete credential no content response has a 3xx status code
func (o *DeleteCredentialNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete credential no content response has a 4xx status code
func (o *DeleteCredentialNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete credential no content response has a 5xx status code
func (o *DeleteCredentialNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete credential no content response a status code equal to that given
func (o *DeleteCredentialNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete credential no content response
func (o *DeleteCredentialNoContent) Code() int {
	return 204
}

func (o *DeleteCredentialNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialNoContent", 204)
}

func (o *DeleteCredentialNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialNoContent", 204)
}

func (o *DeleteCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCredentialForbidden creates a DeleteCredentialForbidden with default headers values
func NewDeleteCredentialForbidden() *DeleteCredentialForbidden {
	return &DeleteCredentialForbidden{}
}

/*
DeleteCredentialForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteCredentialForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete credential forbidden response has a 2xx status code
func (o *DeleteCredentialForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete credential forbidden response has a 3xx status code
func (o *DeleteCredentialForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete credential forbidden response has a 4xx status code
func (o *DeleteCredentialForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete credential forbidden response has a 5xx status code
func (o *DeleteCredentialForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete credential forbidden response a status code equal to that given
func (o *DeleteCredentialForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete credential forbidden response
func (o *DeleteCredentialForbidden) Code() int {
	return 403
}

func (o *DeleteCredentialForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialForbidden %s", 403, payload)
}

func (o *DeleteCredentialForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialForbidden %s", 403, payload)
}

func (o *DeleteCredentialForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCredentialNotFound creates a DeleteCredentialNotFound with default headers values
func NewDeleteCredentialNotFound() *DeleteCredentialNotFound {
	return &DeleteCredentialNotFound{}
}

/*
DeleteCredentialNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteCredentialNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete credential not found response has a 2xx status code
func (o *DeleteCredentialNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete credential not found response has a 3xx status code
func (o *DeleteCredentialNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete credential not found response has a 4xx status code
func (o *DeleteCredentialNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete credential not found response has a 5xx status code
func (o *DeleteCredentialNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete credential not found response a status code equal to that given
func (o *DeleteCredentialNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete credential not found response
func (o *DeleteCredentialNotFound) Code() int {
	return 404
}

func (o *DeleteCredentialNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialNotFound %s", 404, payload)
}

func (o *DeleteCredentialNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialNotFound %s", 404, payload)
}

func (o *DeleteCredentialNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCredentialConflict creates a DeleteCredentialConflict with default headers values
func NewDeleteCredentialConflict() *DeleteCredentialConflict {
	return &DeleteCredentialConflict{}
}

/*
DeleteCredentialConflict describes a response with status code 409, with default header values.

Credential deletion has internal conflict
*/
type DeleteCredentialConflict struct {
}

// IsSuccess returns true when this delete credential conflict response has a 2xx status code
func (o *DeleteCredentialConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete credential conflict response has a 3xx status code
func (o *DeleteCredentialConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete credential conflict response has a 4xx status code
func (o *DeleteCredentialConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete credential conflict response has a 5xx status code
func (o *DeleteCredentialConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete credential conflict response a status code equal to that given
func (o *DeleteCredentialConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete credential conflict response
func (o *DeleteCredentialConflict) Code() int {
	return 409
}

func (o *DeleteCredentialConflict) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialConflict", 409)
}

func (o *DeleteCredentialConflict) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredentialConflict", 409)
}

func (o *DeleteCredentialConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCredentialDefault creates a DeleteCredentialDefault with default headers values
func NewDeleteCredentialDefault(code int) *DeleteCredentialDefault {
	return &DeleteCredentialDefault{
		_statusCode: code,
	}
}

/*
DeleteCredentialDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteCredentialDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this delete credential default response has a 2xx status code
func (o *DeleteCredentialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete credential default response has a 3xx status code
func (o *DeleteCredentialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete credential default response has a 4xx status code
func (o *DeleteCredentialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete credential default response has a 5xx status code
func (o *DeleteCredentialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete credential default response a status code equal to that given
func (o *DeleteCredentialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete credential default response
func (o *DeleteCredentialDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCredentialDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredential default %s", o._statusCode, payload)
}

func (o *DeleteCredentialDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_canonical}][%d] deleteCredential default %s", o._statusCode, payload)
}

func (o *DeleteCredentialDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

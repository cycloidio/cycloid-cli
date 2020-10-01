// Code generated by go-swagger; DO NOT EDIT.

package organization_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
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

/*DeleteCredentialNoContent handles this case with default header values.

Credential has been deleted.
*/
type DeleteCredentialNoContent struct {
}

func (o *DeleteCredentialNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_id}][%d] deleteCredentialNoContent ", 204)
}

func (o *DeleteCredentialNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCredentialForbidden creates a DeleteCredentialForbidden with default headers values
func NewDeleteCredentialForbidden() *DeleteCredentialForbidden {
	return &DeleteCredentialForbidden{}
}

/*DeleteCredentialForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteCredentialForbidden struct {
	Payload *models.ErrorPayload
}

func (o *DeleteCredentialForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_id}][%d] deleteCredentialForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCredentialForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*DeleteCredentialNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteCredentialNotFound struct {
	Payload *models.ErrorPayload
}

func (o *DeleteCredentialNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_id}][%d] deleteCredentialNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCredentialNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*DeleteCredentialConflict handles this case with default header values.

Credential deletion has internal conflict
*/
type DeleteCredentialConflict struct {
}

func (o *DeleteCredentialConflict) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_id}][%d] deleteCredentialConflict ", 409)
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

/*DeleteCredentialDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteCredentialDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the delete credential default response
func (o *DeleteCredentialDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCredentialDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/credentials/{credential_id}][%d] deleteCredential default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCredentialDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

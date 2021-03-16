// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// DeleteStackReader is a Reader for the DeleteStack structure.
type DeleteStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteStackNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteStackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteStackConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteStackNoContent creates a DeleteStackNoContent with default headers values
func NewDeleteStackNoContent() *DeleteStackNoContent {
	return &DeleteStackNoContent{}
}

/*DeleteStackNoContent handles this case with default header values.

Stack has been deleted.
*/
type DeleteStackNoContent struct {
}

func (o *DeleteStackNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/stacks/{stack_ref}][%d] deleteStackNoContent ", 204)
}

func (o *DeleteStackNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStackForbidden creates a DeleteStackForbidden with default headers values
func NewDeleteStackForbidden() *DeleteStackForbidden {
	return &DeleteStackForbidden{}
}

/*DeleteStackForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteStackForbidden struct {
	Payload *models.ErrorPayload
}

func (o *DeleteStackForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/stacks/{stack_ref}][%d] deleteStackForbidden  %+v", 403, o.Payload)
}

func (o *DeleteStackForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteStackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStackNotFound creates a DeleteStackNotFound with default headers values
func NewDeleteStackNotFound() *DeleteStackNotFound {
	return &DeleteStackNotFound{}
}

/*DeleteStackNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteStackNotFound struct {
	Payload *models.ErrorPayload
}

func (o *DeleteStackNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/stacks/{stack_ref}][%d] deleteStackNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStackNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStackConflict creates a DeleteStackConflict with default headers values
func NewDeleteStackConflict() *DeleteStackConflict {
	return &DeleteStackConflict{}
}

/*DeleteStackConflict handles this case with default header values.

Stack deletion has internal conflict
*/
type DeleteStackConflict struct {
}

func (o *DeleteStackConflict) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/stacks/{stack_ref}][%d] deleteStackConflict ", 409)
}

func (o *DeleteStackConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

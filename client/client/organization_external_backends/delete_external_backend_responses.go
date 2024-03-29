// Code generated by go-swagger; DO NOT EDIT.

package organization_external_backends

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

// DeleteExternalBackendReader is a Reader for the DeleteExternalBackend structure.
type DeleteExternalBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteExternalBackendNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteExternalBackendForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalBackendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteExternalBackendUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteExternalBackendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteExternalBackendNoContent creates a DeleteExternalBackendNoContent with default headers values
func NewDeleteExternalBackendNoContent() *DeleteExternalBackendNoContent {
	return &DeleteExternalBackendNoContent{}
}

/*DeleteExternalBackendNoContent handles this case with default header values.

Organization Service Catalog Sources has been deleted
*/
type DeleteExternalBackendNoContent struct {
}

func (o *DeleteExternalBackendNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] deleteExternalBackendNoContent ", 204)
}

func (o *DeleteExternalBackendNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalBackendForbidden creates a DeleteExternalBackendForbidden with default headers values
func NewDeleteExternalBackendForbidden() *DeleteExternalBackendForbidden {
	return &DeleteExternalBackendForbidden{}
}

/*DeleteExternalBackendForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteExternalBackendForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteExternalBackendForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] deleteExternalBackendForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalBackendForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteExternalBackendForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteExternalBackendNotFound creates a DeleteExternalBackendNotFound with default headers values
func NewDeleteExternalBackendNotFound() *DeleteExternalBackendNotFound {
	return &DeleteExternalBackendNotFound{}
}

/*DeleteExternalBackendNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteExternalBackendNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteExternalBackendNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] deleteExternalBackendNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalBackendNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteExternalBackendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteExternalBackendUnprocessableEntity creates a DeleteExternalBackendUnprocessableEntity with default headers values
func NewDeleteExternalBackendUnprocessableEntity() *DeleteExternalBackendUnprocessableEntity {
	return &DeleteExternalBackendUnprocessableEntity{}
}

/*DeleteExternalBackendUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type DeleteExternalBackendUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteExternalBackendUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] deleteExternalBackendUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteExternalBackendUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteExternalBackendUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteExternalBackendDefault creates a DeleteExternalBackendDefault with default headers values
func NewDeleteExternalBackendDefault(code int) *DeleteExternalBackendDefault {
	return &DeleteExternalBackendDefault{
		_statusCode: code,
	}
}

/*DeleteExternalBackendDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteExternalBackendDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the delete external backend default response
func (o *DeleteExternalBackendDefault) Code() int {
	return o._statusCode
}

func (o *DeleteExternalBackendDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] deleteExternalBackend default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteExternalBackendDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteExternalBackendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

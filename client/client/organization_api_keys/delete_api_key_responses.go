// Code generated by go-swagger; DO NOT EDIT.

package organization_api_keys

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

/*DeleteAPIKeyNoContent handles this case with default header values.

API key has been deleted.
*/
type DeleteAPIKeyNoContent struct {
}

func (o *DeleteAPIKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyNoContent ", 204)
}

func (o *DeleteAPIKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIKeyForbidden creates a DeleteAPIKeyForbidden with default headers values
func NewDeleteAPIKeyForbidden() *DeleteAPIKeyForbidden {
	return &DeleteAPIKeyForbidden{}
}

/*DeleteAPIKeyForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteAPIKeyForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteAPIKeyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIKeyForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAPIKeyNotFound creates a DeleteAPIKeyNotFound with default headers values
func NewDeleteAPIKeyNotFound() *DeleteAPIKeyNotFound {
	return &DeleteAPIKeyNotFound{}
}

/*DeleteAPIKeyNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteAPIKeyNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteApiKeyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPIKeyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAPIKeyDefault creates a DeleteAPIKeyDefault with default headers values
func NewDeleteAPIKeyDefault(code int) *DeleteAPIKeyDefault {
	return &DeleteAPIKeyDefault{
		_statusCode: code,
	}
}

/*DeleteAPIKeyDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteAPIKeyDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the delete API key default response
func (o *DeleteAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAPIKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] deleteAPIKey default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAPIKeyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

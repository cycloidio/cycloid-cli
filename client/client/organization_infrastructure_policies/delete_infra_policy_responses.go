// Code generated by go-swagger; DO NOT EDIT.

package organization_infrastructure_policies

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

// DeleteInfraPolicyReader is a Reader for the DeleteInfraPolicy structure.
type DeleteInfraPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInfraPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteInfraPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteInfraPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInfraPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteInfraPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteInfraPolicyNoContent creates a DeleteInfraPolicyNoContent with default headers values
func NewDeleteInfraPolicyNoContent() *DeleteInfraPolicyNoContent {
	return &DeleteInfraPolicyNoContent{}
}

/*DeleteInfraPolicyNoContent handles this case with default header values.

InfraPolicy has been deleted.
*/
type DeleteInfraPolicyNoContent struct {
}

func (o *DeleteInfraPolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] deleteInfraPolicyNoContent ", 204)
}

func (o *DeleteInfraPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInfraPolicyForbidden creates a DeleteInfraPolicyForbidden with default headers values
func NewDeleteInfraPolicyForbidden() *DeleteInfraPolicyForbidden {
	return &DeleteInfraPolicyForbidden{}
}

/*DeleteInfraPolicyForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteInfraPolicyForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteInfraPolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] deleteInfraPolicyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInfraPolicyForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInfraPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInfraPolicyNotFound creates a DeleteInfraPolicyNotFound with default headers values
func NewDeleteInfraPolicyNotFound() *DeleteInfraPolicyNotFound {
	return &DeleteInfraPolicyNotFound{}
}

/*DeleteInfraPolicyNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteInfraPolicyNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteInfraPolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] deleteInfraPolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInfraPolicyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInfraPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInfraPolicyDefault creates a DeleteInfraPolicyDefault with default headers values
func NewDeleteInfraPolicyDefault(code int) *DeleteInfraPolicyDefault {
	return &DeleteInfraPolicyDefault{
		_statusCode: code,
	}
}

/*DeleteInfraPolicyDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteInfraPolicyDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the delete infra policy default response
func (o *DeleteInfraPolicyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInfraPolicyDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/infra_policies/{infra_policy_canonical}][%d] deleteInfraPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteInfraPolicyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteInfraPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

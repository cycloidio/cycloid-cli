// Code generated by go-swagger; DO NOT EDIT.

package organization_members

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

// RemoveOrgMemberReader is a Reader for the RemoveOrgMember structure.
type RemoveOrgMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveOrgMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveOrgMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRemoveOrgMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveOrgMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRemoveOrgMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveOrgMemberNoContent creates a RemoveOrgMemberNoContent with default headers values
func NewRemoveOrgMemberNoContent() *RemoveOrgMemberNoContent {
	return &RemoveOrgMemberNoContent{}
}

/*RemoveOrgMemberNoContent handles this case with default header values.

Member has been removed.
*/
type RemoveOrgMemberNoContent struct {
}

func (o *RemoveOrgMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/members/{username}][%d] removeOrgMemberNoContent ", 204)
}

func (o *RemoveOrgMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveOrgMemberForbidden creates a RemoveOrgMemberForbidden with default headers values
func NewRemoveOrgMemberForbidden() *RemoveOrgMemberForbidden {
	return &RemoveOrgMemberForbidden{}
}

/*RemoveOrgMemberForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type RemoveOrgMemberForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *RemoveOrgMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/members/{username}][%d] removeOrgMemberForbidden  %+v", 403, o.Payload)
}

func (o *RemoveOrgMemberForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RemoveOrgMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveOrgMemberNotFound creates a RemoveOrgMemberNotFound with default headers values
func NewRemoveOrgMemberNotFound() *RemoveOrgMemberNotFound {
	return &RemoveOrgMemberNotFound{}
}

/*RemoveOrgMemberNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type RemoveOrgMemberNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *RemoveOrgMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/members/{username}][%d] removeOrgMemberNotFound  %+v", 404, o.Payload)
}

func (o *RemoveOrgMemberNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RemoveOrgMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveOrgMemberDefault creates a RemoveOrgMemberDefault with default headers values
func NewRemoveOrgMemberDefault(code int) *RemoveOrgMemberDefault {
	return &RemoveOrgMemberDefault{
		_statusCode: code,
	}
}

/*RemoveOrgMemberDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type RemoveOrgMemberDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the remove org member default response
func (o *RemoveOrgMemberDefault) Code() int {
	return o._statusCode
}

func (o *RemoveOrgMemberDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/members/{username}][%d] removeOrgMember default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveOrgMemberDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RemoveOrgMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

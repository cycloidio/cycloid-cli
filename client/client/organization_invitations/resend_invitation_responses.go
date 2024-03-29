// Code generated by go-swagger; DO NOT EDIT.

package organization_invitations

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

// ResendInvitationReader is a Reader for the ResendInvitation structure.
type ResendInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResendInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewResendInvitationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewResendInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResendInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResendInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResendInvitationNoContent creates a ResendInvitationNoContent with default headers values
func NewResendInvitationNoContent() *ResendInvitationNoContent {
	return &ResendInvitationNoContent{}
}

/*ResendInvitationNoContent handles this case with default header values.

The Invitation has been resent.
*/
type ResendInvitationNoContent struct {
}

func (o *ResendInvitationNoContent) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/invitations/{invitation_id}/resend][%d] resendInvitationNoContent ", 204)
}

func (o *ResendInvitationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendInvitationForbidden creates a ResendInvitationForbidden with default headers values
func NewResendInvitationForbidden() *ResendInvitationForbidden {
	return &ResendInvitationForbidden{}
}

/*ResendInvitationForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type ResendInvitationForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *ResendInvitationForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/invitations/{invitation_id}/resend][%d] resendInvitationForbidden  %+v", 403, o.Payload)
}

func (o *ResendInvitationForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ResendInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewResendInvitationNotFound creates a ResendInvitationNotFound with default headers values
func NewResendInvitationNotFound() *ResendInvitationNotFound {
	return &ResendInvitationNotFound{}
}

/*ResendInvitationNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type ResendInvitationNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *ResendInvitationNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/invitations/{invitation_id}/resend][%d] resendInvitationNotFound  %+v", 404, o.Payload)
}

func (o *ResendInvitationNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ResendInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewResendInvitationDefault creates a ResendInvitationDefault with default headers values
func NewResendInvitationDefault(code int) *ResendInvitationDefault {
	return &ResendInvitationDefault{
		_statusCode: code,
	}
}

/*ResendInvitationDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type ResendInvitationDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the resend invitation default response
func (o *ResendInvitationDefault) Code() int {
	return o._statusCode
}

func (o *ResendInvitationDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/invitations/{invitation_id}/resend][%d] resendInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *ResendInvitationDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ResendInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

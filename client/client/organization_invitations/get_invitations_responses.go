// Code generated by go-swagger; DO NOT EDIT.

package organization_invitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// GetInvitationsReader is a Reader for the GetInvitations structure.
type GetInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetInvitationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInvitationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetInvitationsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetInvitationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInvitationsOK creates a GetInvitationsOK with default headers values
func NewGetInvitationsOK() *GetInvitationsOK {
	return &GetInvitationsOK{}
}

/*GetInvitationsOK handles this case with default header values.

List of the Organization's Invitations.
*/
type GetInvitationsOK struct {
	Payload *GetInvitationsOKBody
}

func (o *GetInvitationsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/invitations][%d] getInvitationsOK  %+v", 200, o.Payload)
}

func (o *GetInvitationsOK) GetPayload() *GetInvitationsOKBody {
	return o.Payload
}

func (o *GetInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetInvitationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvitationsForbidden creates a GetInvitationsForbidden with default headers values
func NewGetInvitationsForbidden() *GetInvitationsForbidden {
	return &GetInvitationsForbidden{}
}

/*GetInvitationsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetInvitationsForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetInvitationsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/invitations][%d] getInvitationsForbidden  %+v", 403, o.Payload)
}

func (o *GetInvitationsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetInvitationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInvitationsNotFound creates a GetInvitationsNotFound with default headers values
func NewGetInvitationsNotFound() *GetInvitationsNotFound {
	return &GetInvitationsNotFound{}
}

/*GetInvitationsNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetInvitationsNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetInvitationsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/invitations][%d] getInvitationsNotFound  %+v", 404, o.Payload)
}

func (o *GetInvitationsNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetInvitationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInvitationsUnprocessableEntity creates a GetInvitationsUnprocessableEntity with default headers values
func NewGetInvitationsUnprocessableEntity() *GetInvitationsUnprocessableEntity {
	return &GetInvitationsUnprocessableEntity{}
}

/*GetInvitationsUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetInvitationsUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetInvitationsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/invitations][%d] getInvitationsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetInvitationsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetInvitationsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInvitationsDefault creates a GetInvitationsDefault with default headers values
func NewGetInvitationsDefault(code int) *GetInvitationsDefault {
	return &GetInvitationsDefault{
		_statusCode: code,
	}
}

/*GetInvitationsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetInvitationsDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get invitations default response
func (o *GetInvitationsDefault) Code() int {
	return o._statusCode
}

func (o *GetInvitationsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/invitations][%d] getInvitations default  %+v", o._statusCode, o.Payload)
}

func (o *GetInvitationsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetInvitationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetInvitationsOKBody get invitations o k body
swagger:model GetInvitationsOKBody
*/
type GetInvitationsOKBody struct {

	// data
	// Required: true
	Data []*models.Invitation `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get invitations o k body
func (o *GetInvitationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInvitationsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getInvitationsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getInvitationsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetInvitationsOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getInvitationsOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInvitationsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInvitationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInvitationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetInvitationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

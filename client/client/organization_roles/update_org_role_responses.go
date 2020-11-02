// Code generated by go-swagger; DO NOT EDIT.

package organization_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// UpdateOrgRoleReader is a Reader for the UpdateOrgRole structure.
type UpdateOrgRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrgRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateOrgRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrgRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateOrgRoleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateOrgRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOrgRoleOK creates a UpdateOrgRoleOK with default headers values
func NewUpdateOrgRoleOK() *UpdateOrgRoleOK {
	return &UpdateOrgRoleOK{}
}

/*UpdateOrgRoleOK handles this case with default header values.

Updated role belonging to the organization.
*/
type UpdateOrgRoleOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *UpdateOrgRoleOKBody
}

func (o *UpdateOrgRoleOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_id}][%d] updateOrgRoleOK  %+v", 200, o.Payload)
}

func (o *UpdateOrgRoleOK) GetPayload() *UpdateOrgRoleOKBody {
	return o.Payload
}

func (o *UpdateOrgRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertUint64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "uint64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(UpdateOrgRoleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgRoleForbidden creates a UpdateOrgRoleForbidden with default headers values
func NewUpdateOrgRoleForbidden() *UpdateOrgRoleForbidden {
	return &UpdateOrgRoleForbidden{}
}

/*UpdateOrgRoleForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateOrgRoleForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateOrgRoleForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_id}][%d] updateOrgRoleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrgRoleForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateOrgRoleNotFound creates a UpdateOrgRoleNotFound with default headers values
func NewUpdateOrgRoleNotFound() *UpdateOrgRoleNotFound {
	return &UpdateOrgRoleNotFound{}
}

/*UpdateOrgRoleNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateOrgRoleNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateOrgRoleNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_id}][%d] updateOrgRoleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrgRoleNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateOrgRoleUnprocessableEntity creates a UpdateOrgRoleUnprocessableEntity with default headers values
func NewUpdateOrgRoleUnprocessableEntity() *UpdateOrgRoleUnprocessableEntity {
	return &UpdateOrgRoleUnprocessableEntity{}
}

/*UpdateOrgRoleUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateOrgRoleUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateOrgRoleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_id}][%d] updateOrgRoleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateOrgRoleUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgRoleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateOrgRoleDefault creates a UpdateOrgRoleDefault with default headers values
func NewUpdateOrgRoleDefault(code int) *UpdateOrgRoleDefault {
	return &UpdateOrgRoleDefault{
		_statusCode: code,
	}
}

/*UpdateOrgRoleDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateOrgRoleDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the update org role default response
func (o *UpdateOrgRoleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateOrgRoleDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_id}][%d] updateOrgRole default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateOrgRoleDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*UpdateOrgRoleOKBody update org role o k body
swagger:model UpdateOrgRoleOKBody
*/
type UpdateOrgRoleOKBody struct {

	// data
	// Required: true
	Data *models.Role `json:"data"`
}

// Validate validates this update org role o k body
func (o *UpdateOrgRoleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrgRoleOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateOrgRoleOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrgRoleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrgRoleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrgRoleOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrgRoleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

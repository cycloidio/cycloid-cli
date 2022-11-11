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

// UpdateRoleReader is a Reader for the UpdateRole structure.
type UpdateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateRoleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRoleOK creates a UpdateRoleOK with default headers values
func NewUpdateRoleOK() *UpdateRoleOK {
	return &UpdateRoleOK{}
}

/*UpdateRoleOK handles this case with default header values.

Updated role belonging to the organization.
*/
type UpdateRoleOK struct {
	Payload *UpdateRoleOKBody
}

func (o *UpdateRoleOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_canonical}][%d] updateRoleOK  %+v", 200, o.Payload)
}

func (o *UpdateRoleOK) GetPayload() *UpdateRoleOKBody {
	return o.Payload
}

func (o *UpdateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateRoleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleForbidden creates a UpdateRoleForbidden with default headers values
func NewUpdateRoleForbidden() *UpdateRoleForbidden {
	return &UpdateRoleForbidden{}
}

/*UpdateRoleForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateRoleForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateRoleForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_canonical}][%d] updateRoleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRoleForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRoleNotFound creates a UpdateRoleNotFound with default headers values
func NewUpdateRoleNotFound() *UpdateRoleNotFound {
	return &UpdateRoleNotFound{}
}

/*UpdateRoleNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateRoleNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateRoleNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_canonical}][%d] updateRoleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRoleNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRoleUnprocessableEntity creates a UpdateRoleUnprocessableEntity with default headers values
func NewUpdateRoleUnprocessableEntity() *UpdateRoleUnprocessableEntity {
	return &UpdateRoleUnprocessableEntity{}
}

/*UpdateRoleUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateRoleUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateRoleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_canonical}][%d] updateRoleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateRoleUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateRoleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRoleDefault creates a UpdateRoleDefault with default headers values
func NewUpdateRoleDefault(code int) *UpdateRoleDefault {
	return &UpdateRoleDefault{
		_statusCode: code,
	}
}

/*UpdateRoleDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateRoleDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the update role default response
func (o *UpdateRoleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRoleDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/roles/{role_canonical}][%d] updateRole default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRoleDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*UpdateRoleOKBody update role o k body
swagger:model UpdateRoleOKBody
*/
type UpdateRoleOKBody struct {

	// data
	// Required: true
	Data *models.Role `json:"data"`
}

// Validate validates this update role o k body
func (o *UpdateRoleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateRoleOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateRoleOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateRoleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateRoleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateRoleOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateRoleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
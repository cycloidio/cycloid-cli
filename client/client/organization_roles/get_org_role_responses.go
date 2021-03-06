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

// GetOrgRoleReader is a Reader for the GetOrgRole structure.
type GetOrgRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrgRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgRoleOK creates a GetOrgRoleOK with default headers values
func NewGetOrgRoleOK() *GetOrgRoleOK {
	return &GetOrgRoleOK{}
}

/*GetOrgRoleOK handles this case with default header values.

Role available in the organization with such canonical.
*/
type GetOrgRoleOK struct {
	Payload *GetOrgRoleOKBody
}

func (o *GetOrgRoleOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles/{role_canonical}][%d] getOrgRoleOK  %+v", 200, o.Payload)
}

func (o *GetOrgRoleOK) GetPayload() *GetOrgRoleOKBody {
	return o.Payload
}

func (o *GetOrgRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrgRoleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRoleForbidden creates a GetOrgRoleForbidden with default headers values
func NewGetOrgRoleForbidden() *GetOrgRoleForbidden {
	return &GetOrgRoleForbidden{}
}

/*GetOrgRoleForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetOrgRoleForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles/{role_canonical}][%d] getOrgRoleForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgRoleForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRoleNotFound creates a GetOrgRoleNotFound with default headers values
func NewGetOrgRoleNotFound() *GetOrgRoleNotFound {
	return &GetOrgRoleNotFound{}
}

/*GetOrgRoleNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetOrgRoleNotFound struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgRoleNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles/{role_canonical}][%d] getOrgRoleNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgRoleNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRoleDefault creates a GetOrgRoleDefault with default headers values
func NewGetOrgRoleDefault(code int) *GetOrgRoleDefault {
	return &GetOrgRoleDefault{
		_statusCode: code,
	}
}

/*GetOrgRoleDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgRoleDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get org role default response
func (o *GetOrgRoleDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgRoleDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/roles/{role_canonical}][%d] getOrgRole default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrgRoleDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrgRoleOKBody get org role o k body
swagger:model GetOrgRoleOKBody
*/
type GetOrgRoleOKBody struct {

	// data
	// Required: true
	Data *models.Role `json:"data"`
}

// Validate validates this get org role o k body
func (o *GetOrgRoleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgRoleOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgRoleOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrgRoleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgRoleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgRoleOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgRoleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

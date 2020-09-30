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
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// UpdateOrgMemberReader is a Reader for the UpdateOrgMember structure.
type UpdateOrgMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrgMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateOrgMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrgMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateOrgMemberUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateOrgMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOrgMemberOK creates a UpdateOrgMemberOK with default headers values
func NewUpdateOrgMemberOK() *UpdateOrgMemberOK {
	return &UpdateOrgMemberOK{}
}

/*UpdateOrgMemberOK handles this case with default header values.

The information of the member of the organization.
*/
type UpdateOrgMemberOK struct {
	Payload *UpdateOrgMemberOKBody
}

func (o *UpdateOrgMemberOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/members/{username}][%d] updateOrgMemberOK  %+v", 200, o.Payload)
}

func (o *UpdateOrgMemberOK) GetPayload() *UpdateOrgMemberOKBody {
	return o.Payload
}

func (o *UpdateOrgMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateOrgMemberOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgMemberForbidden creates a UpdateOrgMemberForbidden with default headers values
func NewUpdateOrgMemberForbidden() *UpdateOrgMemberForbidden {
	return &UpdateOrgMemberForbidden{}
}

/*UpdateOrgMemberForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateOrgMemberForbidden struct {
	Payload *models.ErrorPayload
}

func (o *UpdateOrgMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/members/{username}][%d] updateOrgMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrgMemberForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgMemberNotFound creates a UpdateOrgMemberNotFound with default headers values
func NewUpdateOrgMemberNotFound() *UpdateOrgMemberNotFound {
	return &UpdateOrgMemberNotFound{}
}

/*UpdateOrgMemberNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateOrgMemberNotFound struct {
	Payload *models.ErrorPayload
}

func (o *UpdateOrgMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/members/{username}][%d] updateOrgMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrgMemberNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgMemberUnprocessableEntity creates a UpdateOrgMemberUnprocessableEntity with default headers values
func NewUpdateOrgMemberUnprocessableEntity() *UpdateOrgMemberUnprocessableEntity {
	return &UpdateOrgMemberUnprocessableEntity{}
}

/*UpdateOrgMemberUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateOrgMemberUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *UpdateOrgMemberUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/members/{username}][%d] updateOrgMemberUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateOrgMemberUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgMemberUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrgMemberDefault creates a UpdateOrgMemberDefault with default headers values
func NewUpdateOrgMemberDefault(code int) *UpdateOrgMemberDefault {
	return &UpdateOrgMemberDefault{
		_statusCode: code,
	}
}

/*UpdateOrgMemberDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateOrgMemberDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the update org member default response
func (o *UpdateOrgMemberDefault) Code() int {
	return o._statusCode
}

func (o *UpdateOrgMemberDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/members/{username}][%d] updateOrgMember default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateOrgMemberDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateOrgMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateOrgMemberOKBody update org member o k body
swagger:model UpdateOrgMemberOKBody
*/
type UpdateOrgMemberOKBody struct {

	// data
	// Required: true
	Data *models.MemberOrg `json:"data"`
}

// Validate validates this update org member o k body
func (o *UpdateOrgMemberOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrgMemberOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateOrgMemberOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrgMemberOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrgMemberOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrgMemberOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrgMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organization_members

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

// GetOrgMembersReader is a Reader for the GetOrgMembers structure.
type GetOrgMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrgMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetOrgMembersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgMembersOK creates a GetOrgMembersOK with default headers values
func NewGetOrgMembersOK() *GetOrgMembersOK {
	return &GetOrgMembersOK{}
}

/*GetOrgMembersOK handles this case with default header values.

List of the members of the organization.
*/
type GetOrgMembersOK struct {
	Payload *GetOrgMembersOKBody
}

func (o *GetOrgMembersOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members][%d] getOrgMembersOK  %+v", 200, o.Payload)
}

func (o *GetOrgMembersOK) GetPayload() *GetOrgMembersOKBody {
	return o.Payload
}

func (o *GetOrgMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrgMembersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMembersForbidden creates a GetOrgMembersForbidden with default headers values
func NewGetOrgMembersForbidden() *GetOrgMembersForbidden {
	return &GetOrgMembersForbidden{}
}

/*GetOrgMembersForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetOrgMembersForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members][%d] getOrgMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgMembersForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMembersNotFound creates a GetOrgMembersNotFound with default headers values
func NewGetOrgMembersNotFound() *GetOrgMembersNotFound {
	return &GetOrgMembersNotFound{}
}

/*GetOrgMembersNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetOrgMembersNotFound struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members][%d] getOrgMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgMembersNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMembersUnprocessableEntity creates a GetOrgMembersUnprocessableEntity with default headers values
func NewGetOrgMembersUnprocessableEntity() *GetOrgMembersUnprocessableEntity {
	return &GetOrgMembersUnprocessableEntity{}
}

/*GetOrgMembersUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetOrgMembersUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgMembersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members][%d] getOrgMembersUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetOrgMembersUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMembersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMembersDefault creates a GetOrgMembersDefault with default headers values
func NewGetOrgMembersDefault(code int) *GetOrgMembersDefault {
	return &GetOrgMembersDefault{
		_statusCode: code,
	}
}

/*GetOrgMembersDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgMembersDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get org members default response
func (o *GetOrgMembersDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgMembersDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members][%d] getOrgMembers default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrgMembersDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrgMembersOKBody get org members o k body
swagger:model GetOrgMembersOKBody
*/
type GetOrgMembersOKBody struct {

	// data
	// Required: true
	Data []*models.MemberOrg `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get org members o k body
func (o *GetOrgMembersOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOrgMembersOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgMembersOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrgMembersOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrgMembersOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getOrgMembersOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrgMembersOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgMembersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgMembersOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgMembersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

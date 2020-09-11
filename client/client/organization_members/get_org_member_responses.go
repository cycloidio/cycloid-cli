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

// GetOrgMemberReader is a Reader for the GetOrgMember structure.
type GetOrgMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrgMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgMemberOK creates a GetOrgMemberOK with default headers values
func NewGetOrgMemberOK() *GetOrgMemberOK {
	return &GetOrgMemberOK{}
}

/*GetOrgMemberOK handles this case with default header values.

The information of the member of the organization.
*/
type GetOrgMemberOK struct {
	Payload *GetOrgMemberOKBody
}

func (o *GetOrgMemberOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberOK  %+v", 200, o.Payload)
}

func (o *GetOrgMemberOK) GetPayload() *GetOrgMemberOKBody {
	return o.Payload
}

func (o *GetOrgMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrgMemberOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMemberForbidden creates a GetOrgMemberForbidden with default headers values
func NewGetOrgMemberForbidden() *GetOrgMemberForbidden {
	return &GetOrgMemberForbidden{}
}

/*GetOrgMemberForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetOrgMemberForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetOrgMemberForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgMemberForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMemberNotFound creates a GetOrgMemberNotFound with default headers values
func NewGetOrgMemberNotFound() *GetOrgMemberNotFound {
	return &GetOrgMemberNotFound{}
}

/*GetOrgMemberNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetOrgMemberNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetOrgMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgMemberNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgMemberDefault creates a GetOrgMemberDefault with default headers values
func NewGetOrgMemberDefault(code int) *GetOrgMemberDefault {
	return &GetOrgMemberDefault{
		_statusCode: code,
	}
}

/*GetOrgMemberDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgMemberDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get org member default response
func (o *GetOrgMemberDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgMemberDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/members/{username}][%d] getOrgMember default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrgMemberDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrgMemberOKBody get org member o k body
swagger:model GetOrgMemberOKBody
*/
type GetOrgMemberOKBody struct {

	// data
	// Required: true
	Data *models.MemberOrg `json:"data"`
}

// Validate validates this get org member o k body
func (o *GetOrgMemberOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgMemberOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgMemberOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrgMemberOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgMemberOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgMemberOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

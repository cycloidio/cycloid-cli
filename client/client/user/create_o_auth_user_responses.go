// Code generated by go-swagger; DO NOT EDIT.

package user

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

// CreateOAuthUserReader is a Reader for the CreateOAuthUser structure.
type CreateOAuthUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOAuthUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOAuthUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateOAuthUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateOAuthUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOAuthUserOK creates a CreateOAuthUserOK with default headers values
func NewCreateOAuthUserOK() *CreateOAuthUserOK {
	return &CreateOAuthUserOK{}
}

/*CreateOAuthUserOK handles this case with default header values.

Create a user from the OAuth 'social_type'
*/
type CreateOAuthUserOK struct {
	Payload *CreateOAuthUserOKBody
}

func (o *CreateOAuthUserOK) Error() string {
	return fmt.Sprintf("[POST /user/{social_type}/oauth][%d] createOAuthUserOK  %+v", 200, o.Payload)
}

func (o *CreateOAuthUserOK) GetPayload() *CreateOAuthUserOKBody {
	return o.Payload
}

func (o *CreateOAuthUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateOAuthUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOAuthUserUnauthorized creates a CreateOAuthUserUnauthorized with default headers values
func NewCreateOAuthUserUnauthorized() *CreateOAuthUserUnauthorized {
	return &CreateOAuthUserUnauthorized{}
}

/*CreateOAuthUserUnauthorized handles this case with default header values.

The user cannot be authenticated with the credentials which she/he has used.
*/
type CreateOAuthUserUnauthorized struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateOAuthUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/{social_type}/oauth][%d] createOAuthUserUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOAuthUserUnauthorized) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOAuthUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateOAuthUserDefault creates a CreateOAuthUserDefault with default headers values
func NewCreateOAuthUserDefault(code int) *CreateOAuthUserDefault {
	return &CreateOAuthUserDefault{
		_statusCode: code,
	}
}

/*CreateOAuthUserDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateOAuthUserDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the create o auth user default response
func (o *CreateOAuthUserDefault) Code() int {
	return o._statusCode
}

func (o *CreateOAuthUserDefault) Error() string {
	return fmt.Sprintf("[POST /user/{social_type}/oauth][%d] createOAuthUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOAuthUserDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateOAuthUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateOAuthUserOKBody create o auth user o k body
swagger:model CreateOAuthUserOKBody
*/
type CreateOAuthUserOKBody struct {

	// data
	// Required: true
	Data *models.UserSession `json:"data"`
}

// Validate validates this create o auth user o k body
func (o *CreateOAuthUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOAuthUserOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createOAuthUserOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOAuthUserOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOAuthUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOAuthUserOKBody) UnmarshalBinary(b []byte) error {
	var res CreateOAuthUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

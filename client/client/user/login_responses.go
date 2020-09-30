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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewLoginUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*LoginOK handles this case with default header values.

The token which represents the session of the user.
*/
type LoginOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *LoginOKBody
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[POST /user/login][%d] loginOK  %+v", 200, o.Payload)
}

func (o *LoginOK) GetPayload() *LoginOKBody {
	return o.Payload
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(LoginOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUnauthorized creates a LoginUnauthorized with default headers values
func NewLoginUnauthorized() *LoginUnauthorized {
	return &LoginUnauthorized{}
}

/*LoginUnauthorized handles this case with default header values.

The user cannot be authenticated with the credentials which she/he has used.
*/
type LoginUnauthorized struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *LoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/login][%d] loginUnauthorized  %+v", 401, o.Payload)
}

func (o *LoginUnauthorized) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *LoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLoginUnprocessableEntity creates a LoginUnprocessableEntity with default headers values
func NewLoginUnprocessableEntity() *LoginUnprocessableEntity {
	return &LoginUnprocessableEntity{}
}

/*LoginUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type LoginUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *LoginUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /user/login][%d] loginUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *LoginUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *LoginUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLoginDefault creates a LoginDefault with default headers values
func NewLoginDefault(code int) *LoginDefault {
	return &LoginDefault{
		_statusCode: code,
	}
}

/*LoginDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type LoginDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the login default response
func (o *LoginDefault) Code() int {
	return o._statusCode
}

func (o *LoginDefault) Error() string {
	return fmt.Sprintf("[POST /user/login][%d] login default  %+v", o._statusCode, o.Payload)
}

func (o *LoginDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *LoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*LoginOKBody login o k body
swagger:model LoginOKBody
*/
type LoginOKBody struct {

	// data
	// Required: true
	Data *models.UserSession `json:"data"`
}

// Validate validates this login o k body
func (o *LoginOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("loginOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loginOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *LoginOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LoginOKBody) UnmarshalBinary(b []byte) error {
	var res LoginOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

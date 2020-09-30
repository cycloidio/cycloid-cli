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

// RefreshTokenReader is a Reader for the RefreshToken structure.
type RefreshTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRefreshTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRefreshTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRefreshTokenOK creates a RefreshTokenOK with default headers values
func NewRefreshTokenOK() *RefreshTokenOK {
	return &RefreshTokenOK{}
}

/*RefreshTokenOK handles this case with default header values.

The token which represents the session of the user.
*/
type RefreshTokenOK struct {
	Payload *RefreshTokenOKBody
}

func (o *RefreshTokenOK) Error() string {
	return fmt.Sprintf("[GET /user/refresh_token][%d] refreshTokenOK  %+v", 200, o.Payload)
}

func (o *RefreshTokenOK) GetPayload() *RefreshTokenOKBody {
	return o.Payload
}

func (o *RefreshTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RefreshTokenOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshTokenUnauthorized creates a RefreshTokenUnauthorized with default headers values
func NewRefreshTokenUnauthorized() *RefreshTokenUnauthorized {
	return &RefreshTokenUnauthorized{}
}

/*RefreshTokenUnauthorized handles this case with default header values.

The user cannot be authenticated with the credentials which she/he has used.
*/
type RefreshTokenUnauthorized struct {
	Payload *models.ErrorPayload
}

func (o *RefreshTokenUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/refresh_token][%d] refreshTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *RefreshTokenUnauthorized) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RefreshTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshTokenDefault creates a RefreshTokenDefault with default headers values
func NewRefreshTokenDefault(code int) *RefreshTokenDefault {
	return &RefreshTokenDefault{
		_statusCode: code,
	}
}

/*RefreshTokenDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type RefreshTokenDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the refresh token default response
func (o *RefreshTokenDefault) Code() int {
	return o._statusCode
}

func (o *RefreshTokenDefault) Error() string {
	return fmt.Sprintf("[GET /user/refresh_token][%d] refreshToken default  %+v", o._statusCode, o.Payload)
}

func (o *RefreshTokenDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RefreshTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RefreshTokenOKBody refresh token o k body
swagger:model RefreshTokenOKBody
*/
type RefreshTokenOKBody struct {

	// data
	// Required: true
	Data *models.UserSession `json:"data"`
}

// Validate validates this refresh token o k body
func (o *RefreshTokenOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RefreshTokenOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("refreshTokenOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("refreshTokenOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RefreshTokenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RefreshTokenOKBody) UnmarshalBinary(b []byte) error {
	var res RefreshTokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

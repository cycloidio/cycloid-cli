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

// GetUserAccountReader is a Reader for the GetUserAccount structure.
type GetUserAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUserAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserAccountOK creates a GetUserAccountOK with default headers values
func NewGetUserAccountOK() *GetUserAccountOK {
	return &GetUserAccountOK{}
}

/*GetUserAccountOK handles this case with default header values.

The user account information.
*/
type GetUserAccountOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *GetUserAccountOKBody
}

func (o *GetUserAccountOK) Error() string {
	return fmt.Sprintf("[GET /user][%d] getUserAccountOK  %+v", 200, o.Payload)
}

func (o *GetUserAccountOK) GetPayload() *GetUserAccountOKBody {
	return o.Payload
}

func (o *GetUserAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetUserAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserAccountDefault creates a GetUserAccountDefault with default headers values
func NewGetUserAccountDefault(code int) *GetUserAccountDefault {
	return &GetUserAccountDefault{
		_statusCode: code,
	}
}

/*GetUserAccountDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetUserAccountDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get user account default response
func (o *GetUserAccountDefault) Code() int {
	return o._statusCode
}

func (o *GetUserAccountDefault) Error() string {
	return fmt.Sprintf("[GET /user][%d] getUserAccount default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserAccountDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetUserAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUserAccountOKBody get user account o k body
swagger:model GetUserAccountOKBody
*/
type GetUserAccountOKBody struct {

	// data
	// Required: true
	Data *models.UserAccount `json:"data"`
}

// Validate validates this get user account o k body
func (o *GetUserAccountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserAccountOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getUserAccountOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUserAccountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserAccountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserAccountOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserAccountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
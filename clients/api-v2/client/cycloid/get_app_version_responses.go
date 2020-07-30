// Code generated by go-swagger; DO NOT EDIT.

package cycloid

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

	models "github.com/cycloidio/youdeploy-cli/clients/api-v2/models"
)

// GetAppVersionReader is a Reader for the GetAppVersion structure.
type GetAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewGetAppVersionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAppVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppVersionOK creates a GetAppVersionOK with default headers values
func NewGetAppVersionOK() *GetAppVersionOK {
	return &GetAppVersionOK{}
}

/*GetAppVersionOK handles this case with default header values.

Application version.
*/
type GetAppVersionOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *GetAppVersionOKBody
}

func (o *GetAppVersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] getAppVersionOK  %+v", 200, o.Payload)
}

func (o *GetAppVersionOK) GetPayload() *GetAppVersionOKBody {
	return o.Payload
}

func (o *GetAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetAppVersionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppVersionUnprocessableEntity creates a GetAppVersionUnprocessableEntity with default headers values
func NewGetAppVersionUnprocessableEntity() *GetAppVersionUnprocessableEntity {
	return &GetAppVersionUnprocessableEntity{}
}

/*GetAppVersionUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetAppVersionUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetAppVersionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /version][%d] getAppVersionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetAppVersionUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAppVersionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAppVersionDefault creates a GetAppVersionDefault with default headers values
func NewGetAppVersionDefault(code int) *GetAppVersionDefault {
	return &GetAppVersionDefault{
		_statusCode: code,
	}
}

/*GetAppVersionDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetAppVersionDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get app version default response
func (o *GetAppVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetAppVersionDefault) Error() string {
	return fmt.Sprintf("[GET /version][%d] getAppVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppVersionDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAppVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetAppVersionOKBody get app version o k body
swagger:model GetAppVersionOKBody
*/
type GetAppVersionOKBody struct {

	// data
	// Required: true
	Data *models.AppVersion `json:"data"`
}

// Validate validates this get app version o k body
func (o *GetAppVersionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAppVersionOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getAppVersionOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAppVersionOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAppVersionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAppVersionOKBody) UnmarshalBinary(b []byte) error {
	var res GetAppVersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

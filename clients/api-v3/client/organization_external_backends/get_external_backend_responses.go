// Code generated by go-swagger; DO NOT EDIT.

package organization_external_backends

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

	models "github.com/cycloidio/youdeploy-cli/clients/api-v3/models"
)

// GetExternalBackendReader is a Reader for the GetExternalBackend structure.
type GetExternalBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalBackendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetExternalBackendForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetExternalBackendUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetExternalBackendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalBackendOK creates a GetExternalBackendOK with default headers values
func NewGetExternalBackendOK() *GetExternalBackendOK {
	return &GetExternalBackendOK{}
}

/*GetExternalBackendOK handles this case with default header values.

The external backend
*/
type GetExternalBackendOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *GetExternalBackendOKBody
}

func (o *GetExternalBackendOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] getExternalBackendOK  %+v", 200, o.Payload)
}

func (o *GetExternalBackendOK) GetPayload() *GetExternalBackendOKBody {
	return o.Payload
}

func (o *GetExternalBackendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetExternalBackendOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalBackendForbidden creates a GetExternalBackendForbidden with default headers values
func NewGetExternalBackendForbidden() *GetExternalBackendForbidden {
	return &GetExternalBackendForbidden{}
}

/*GetExternalBackendForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetExternalBackendForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetExternalBackendForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] getExternalBackendForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalBackendForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetExternalBackendForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetExternalBackendUnprocessableEntity creates a GetExternalBackendUnprocessableEntity with default headers values
func NewGetExternalBackendUnprocessableEntity() *GetExternalBackendUnprocessableEntity {
	return &GetExternalBackendUnprocessableEntity{}
}

/*GetExternalBackendUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetExternalBackendUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetExternalBackendUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] getExternalBackendUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetExternalBackendUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetExternalBackendUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetExternalBackendDefault creates a GetExternalBackendDefault with default headers values
func NewGetExternalBackendDefault(code int) *GetExternalBackendDefault {
	return &GetExternalBackendDefault{
		_statusCode: code,
	}
}

/*GetExternalBackendDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetExternalBackendDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get external backend default response
func (o *GetExternalBackendDefault) Code() int {
	return o._statusCode
}

func (o *GetExternalBackendDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends/{external_backend_id}][%d] getExternalBackend default  %+v", o._statusCode, o.Payload)
}

func (o *GetExternalBackendDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetExternalBackendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetExternalBackendOKBody get external backend o k body
swagger:model GetExternalBackendOKBody
*/
type GetExternalBackendOKBody struct {

	// data
	// Required: true
	Data *models.ExternalBackend `json:"data"`
}

// Validate validates this get external backend o k body
func (o *GetExternalBackendOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetExternalBackendOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getExternalBackendOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getExternalBackendOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetExternalBackendOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetExternalBackendOKBody) UnmarshalBinary(b []byte) error {
	var res GetExternalBackendOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

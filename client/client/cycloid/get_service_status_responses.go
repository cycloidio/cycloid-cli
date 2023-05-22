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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// GetServiceStatusReader is a Reader for the GetServiceStatus structure.
type GetServiceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetServiceStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceStatusOK creates a GetServiceStatusOK with default headers values
func NewGetServiceStatusOK() *GetServiceStatusOK {
	return &GetServiceStatusOK{}
}

/*GetServiceStatusOK handles this case with default header values.

General application status and services statuses.
*/
type GetServiceStatusOK struct {
	Payload *GetServiceStatusOKBody
}

func (o *GetServiceStatusOK) Error() string {
	return fmt.Sprintf("[GET /status/{service_status_canonical}][%d] getServiceStatusOK  %+v", 200, o.Payload)
}

func (o *GetServiceStatusOK) GetPayload() *GetServiceStatusOKBody {
	return o.Payload
}

func (o *GetServiceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceStatusInternalServerError creates a GetServiceStatusInternalServerError with default headers values
func NewGetServiceStatusInternalServerError() *GetServiceStatusInternalServerError {
	return &GetServiceStatusInternalServerError{}
}

/*GetServiceStatusInternalServerError handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetServiceStatusInternalServerError struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetServiceStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /status/{service_status_canonical}][%d] getServiceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServiceStatusInternalServerError) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetServiceStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetServiceStatusOKBody get service status o k body
swagger:model GetServiceStatusOKBody
*/
type GetServiceStatusOKBody struct {

	// data
	// Required: true
	Data *models.CheckReport `json:"data"`
}

// Validate validates this get service status o k body
func (o *GetServiceStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceStatusOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getServiceStatusOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceStatusOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
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

// CreateExternalBackendReader is a Reader for the CreateExternalBackend structure.
type CreateExternalBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExternalBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateExternalBackendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateExternalBackendOK creates a CreateExternalBackendOK with default headers values
func NewCreateExternalBackendOK() *CreateExternalBackendOK {
	return &CreateExternalBackendOK{}
}

/*CreateExternalBackendOK handles this case with default header values.

external backend has been registered
*/
type CreateExternalBackendOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *CreateExternalBackendOKBody
}

func (o *CreateExternalBackendOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/external_backends][%d] createExternalBackendOK  %+v", 200, o.Payload)
}

func (o *CreateExternalBackendOK) GetPayload() *CreateExternalBackendOKBody {
	return o.Payload
}

func (o *CreateExternalBackendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(CreateExternalBackendOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateExternalBackendOKBody create external backend o k body
swagger:model CreateExternalBackendOKBody
*/
type CreateExternalBackendOKBody struct {

	// data
	// Required: true
	Data *models.ExternalBackend `json:"data"`
}

// Validate validates this create external backend o k body
func (o *CreateExternalBackendOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateExternalBackendOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createExternalBackendOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createExternalBackendOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateExternalBackendOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateExternalBackendOKBody) UnmarshalBinary(b []byte) error {
	var res CreateExternalBackendOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
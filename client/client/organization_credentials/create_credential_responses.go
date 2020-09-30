// Code generated by go-swagger; DO NOT EDIT.

package organization_credentials

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

// CreateCredentialReader is a Reader for the CreateCredential structure.
type CreateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 411:
		result := NewCreateCredentialLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCredentialUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCredentialOK creates a CreateCredentialOK with default headers values
func NewCreateCredentialOK() *CreateCredentialOK {
	return &CreateCredentialOK{}
}

/*CreateCredentialOK handles this case with default header values.

Credential created. The body contains the information of the new created Credential.
*/
type CreateCredentialOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *CreateCredentialOKBody
}

func (o *CreateCredentialOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/credentials][%d] createCredentialOK  %+v", 200, o.Payload)
}

func (o *CreateCredentialOK) GetPayload() *CreateCredentialOKBody {
	return o.Payload
}

func (o *CreateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(CreateCredentialOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCredentialLengthRequired creates a CreateCredentialLengthRequired with default headers values
func NewCreateCredentialLengthRequired() *CreateCredentialLengthRequired {
	return &CreateCredentialLengthRequired{}
}

/*CreateCredentialLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type CreateCredentialLengthRequired struct {
}

func (o *CreateCredentialLengthRequired) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/credentials][%d] createCredentialLengthRequired ", 411)
}

func (o *CreateCredentialLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCredentialUnprocessableEntity creates a CreateCredentialUnprocessableEntity with default headers values
func NewCreateCredentialUnprocessableEntity() *CreateCredentialUnprocessableEntity {
	return &CreateCredentialUnprocessableEntity{}
}

/*CreateCredentialUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type CreateCredentialUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *CreateCredentialUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/credentials][%d] createCredentialUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateCredentialUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateCredentialUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCredentialDefault creates a CreateCredentialDefault with default headers values
func NewCreateCredentialDefault(code int) *CreateCredentialDefault {
	return &CreateCredentialDefault{
		_statusCode: code,
	}
}

/*CreateCredentialDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateCredentialDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the create credential default response
func (o *CreateCredentialDefault) Code() int {
	return o._statusCode
}

func (o *CreateCredentialDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/credentials][%d] createCredential default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCredentialDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateCredentialOKBody create credential o k body
swagger:model CreateCredentialOKBody
*/
type CreateCredentialOKBody struct {

	// data
	// Required: true
	Data *models.Credential `json:"data"`
}

// Validate validates this create credential o k body
func (o *CreateCredentialOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateCredentialOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createCredentialOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createCredentialOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateCredentialOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateCredentialOKBody) UnmarshalBinary(b []byte) error {
	var res CreateCredentialOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

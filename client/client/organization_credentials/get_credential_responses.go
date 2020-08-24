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

// GetCredentialReader is a Reader for the GetCredential structure.
type GetCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCredentialOK creates a GetCredentialOK with default headers values
func NewGetCredentialOK() *GetCredentialOK {
	return &GetCredentialOK{}
}

/*GetCredentialOK handles this case with default header values.

The information of the Credential which has the specified ID.
*/
type GetCredentialOK struct {
	Payload *GetCredentialOKBody
}

func (o *GetCredentialOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/credentials/{credential_id}][%d] getCredentialOK  %+v", 200, o.Payload)
}

func (o *GetCredentialOK) GetPayload() *GetCredentialOKBody {
	return o.Payload
}

func (o *GetCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCredentialOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialForbidden creates a GetCredentialForbidden with default headers values
func NewGetCredentialForbidden() *GetCredentialForbidden {
	return &GetCredentialForbidden{}
}

/*GetCredentialForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetCredentialForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetCredentialForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/credentials/{credential_id}][%d] getCredentialForbidden  %+v", 403, o.Payload)
}

func (o *GetCredentialForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCredentialNotFound creates a GetCredentialNotFound with default headers values
func NewGetCredentialNotFound() *GetCredentialNotFound {
	return &GetCredentialNotFound{}
}

/*GetCredentialNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetCredentialNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetCredentialNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/credentials/{credential_id}][%d] getCredentialNotFound  %+v", 404, o.Payload)
}

func (o *GetCredentialNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCredentialDefault creates a GetCredentialDefault with default headers values
func NewGetCredentialDefault(code int) *GetCredentialDefault {
	return &GetCredentialDefault{
		_statusCode: code,
	}
}

/*GetCredentialDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetCredentialDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get credential default response
func (o *GetCredentialDefault) Code() int {
	return o._statusCode
}

func (o *GetCredentialDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/credentials/{credential_id}][%d] getCredential default  %+v", o._statusCode, o.Payload)
}

func (o *GetCredentialDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCredentialOKBody get credential o k body
swagger:model GetCredentialOKBody
*/
type GetCredentialOKBody struct {

	// data
	// Required: true
	Data *models.Credential `json:"data"`
}

// Validate validates this get credential o k body
func (o *GetCredentialOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCredentialOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getCredentialOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCredentialOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCredentialOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCredentialOKBody) UnmarshalBinary(b []byte) error {
	var res GetCredentialOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

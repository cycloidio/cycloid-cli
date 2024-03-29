// Code generated by go-swagger; DO NOT EDIT.

package organization_api_keys

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

// GetAPIKeyReader is a Reader for the GetAPIKey structure.
type GetAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAPIKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIKeyOK creates a GetAPIKeyOK with default headers values
func NewGetAPIKeyOK() *GetAPIKeyOK {
	return &GetAPIKeyOK{}
}

/*GetAPIKeyOK handles this case with default header values.

The information of the API key of the organization which has the specified canonical.
*/
type GetAPIKeyOK struct {
	Payload *GetAPIKeyOKBody
}

func (o *GetAPIKeyOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyOK  %+v", 200, o.Payload)
}

func (o *GetAPIKeyOK) GetPayload() *GetAPIKeyOKBody {
	return o.Payload
}

func (o *GetAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAPIKeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIKeyForbidden creates a GetAPIKeyForbidden with default headers values
func NewGetAPIKeyForbidden() *GetAPIKeyForbidden {
	return &GetAPIKeyForbidden{}
}

/*GetAPIKeyForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetAPIKeyForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetAPIKeyForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIKeyForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAPIKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAPIKeyNotFound creates a GetAPIKeyNotFound with default headers values
func NewGetAPIKeyNotFound() *GetAPIKeyNotFound {
	return &GetAPIKeyNotFound{}
}

/*GetAPIKeyNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetAPIKeyNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getApiKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetAPIKeyNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAPIKeyDefault creates a GetAPIKeyDefault with default headers values
func NewGetAPIKeyDefault(code int) *GetAPIKeyDefault {
	return &GetAPIKeyDefault{
		_statusCode: code,
	}
}

/*GetAPIKeyDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetAPIKeyDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get API key default response
func (o *GetAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIKeyDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/api_keys/{api_key_canonical}][%d] getAPIKey default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPIKeyDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetAPIKeyOKBody get API key o k body
swagger:model GetAPIKeyOKBody
*/
type GetAPIKeyOKBody struct {

	// data
	// Required: true
	Data *models.APIKey `json:"data"`
}

// Validate validates this get API key o k body
func (o *GetAPIKeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIKeyOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getApiKeyOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getApiKeyOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIKeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIKeyOKBody) UnmarshalBinary(b []byte) error {
	var res GetAPIKeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

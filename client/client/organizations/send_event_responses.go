// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// SendEventReader is a Reader for the SendEvent structure.
type SendEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSendEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSendEventNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSendEventUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendEventOK creates a SendEventOK with default headers values
func NewSendEventOK() *SendEventOK {
	return &SendEventOK{}
}

/*SendEventOK handles this case with default header values.

Event has been registered
*/
type SendEventOK struct {
	Payload *SendEventOKBody
}

func (o *SendEventOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendEventOK  %+v", 200, o.Payload)
}

func (o *SendEventOK) GetPayload() *SendEventOKBody {
	return o.Payload
}

func (o *SendEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SendEventOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendEventForbidden creates a SendEventForbidden with default headers values
func NewSendEventForbidden() *SendEventForbidden {
	return &SendEventForbidden{}
}

/*SendEventForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type SendEventForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *SendEventForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendEventForbidden  %+v", 403, o.Payload)
}

func (o *SendEventForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *SendEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSendEventNotFound creates a SendEventNotFound with default headers values
func NewSendEventNotFound() *SendEventNotFound {
	return &SendEventNotFound{}
}

/*SendEventNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type SendEventNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *SendEventNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendEventNotFound  %+v", 404, o.Payload)
}

func (o *SendEventNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *SendEventNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSendEventUnprocessableEntity creates a SendEventUnprocessableEntity with default headers values
func NewSendEventUnprocessableEntity() *SendEventUnprocessableEntity {
	return &SendEventUnprocessableEntity{}
}

/*SendEventUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type SendEventUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *SendEventUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/events][%d] sendEventUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SendEventUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *SendEventUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*SendEventOKBody The newly created event
swagger:model SendEventOKBody
*/
type SendEventOKBody struct {

	// data
	// Required: true
	Data *models.Event `json:"data"`
}

// Validate validates this send event o k body
func (o *SendEventOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendEventOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("sendEventOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sendEventOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendEventOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendEventOKBody) UnmarshalBinary(b []byte) error {
	var res SendEventOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

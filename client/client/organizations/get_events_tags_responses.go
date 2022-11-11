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

// GetEventsTagsReader is a Reader for the GetEventsTags structure.
type GetEventsTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetEventsTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetEventsTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetEventsTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventsTagsOK creates a GetEventsTagsOK with default headers values
func NewGetEventsTagsOK() *GetEventsTagsOK {
	return &GetEventsTagsOK{}
}

/*GetEventsTagsOK handles this case with default header values.

The list of tags and set of values for all the events of the organization.
*/
type GetEventsTagsOK struct {
	Payload *GetEventsTagsOKBody
}

func (o *GetEventsTagsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getEventsTagsOK  %+v", 200, o.Payload)
}

func (o *GetEventsTagsOK) GetPayload() *GetEventsTagsOKBody {
	return o.Payload
}

func (o *GetEventsTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEventsTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsTagsForbidden creates a GetEventsTagsForbidden with default headers values
func NewGetEventsTagsForbidden() *GetEventsTagsForbidden {
	return &GetEventsTagsForbidden{}
}

/*GetEventsTagsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetEventsTagsForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetEventsTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getEventsTagsForbidden  %+v", 403, o.Payload)
}

func (o *GetEventsTagsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetEventsTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEventsTagsUnprocessableEntity creates a GetEventsTagsUnprocessableEntity with default headers values
func NewGetEventsTagsUnprocessableEntity() *GetEventsTagsUnprocessableEntity {
	return &GetEventsTagsUnprocessableEntity{}
}

/*GetEventsTagsUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetEventsTagsUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetEventsTagsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getEventsTagsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetEventsTagsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetEventsTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEventsTagsDefault creates a GetEventsTagsDefault with default headers values
func NewGetEventsTagsDefault(code int) *GetEventsTagsDefault {
	return &GetEventsTagsDefault{
		_statusCode: code,
	}
}

/*GetEventsTagsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetEventsTagsDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get events tags default response
func (o *GetEventsTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetEventsTagsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getEventsTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventsTagsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetEventsTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetEventsTagsOKBody The list of tags with associated set of values
swagger:model GetEventsTagsOKBody
*/
type GetEventsTagsOKBody struct {

	// data
	// Required: true
	Data interface{} `json:"data"`
}

// Validate validates this get events tags o k body
func (o *GetEventsTagsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventsTagsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getEventsTagsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEventsTagsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEventsTagsOKBody) UnmarshalBinary(b []byte) error {
	var res GetEventsTagsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
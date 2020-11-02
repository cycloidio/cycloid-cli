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

// GetOrgEventsTagsReader is a Reader for the GetOrgEventsTags structure.
type GetOrgEventsTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgEventsTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgEventsTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrgEventsTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetOrgEventsTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgEventsTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgEventsTagsOK creates a GetOrgEventsTagsOK with default headers values
func NewGetOrgEventsTagsOK() *GetOrgEventsTagsOK {
	return &GetOrgEventsTagsOK{}
}

/*GetOrgEventsTagsOK handles this case with default header values.

The list of tags and set of values for all the events of the organization.
  format: int64
*/
type GetOrgEventsTagsOK struct {
	Payload *GetOrgEventsTagsOKBody
}

func (o *GetOrgEventsTagsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getOrgEventsTagsOK  %+v", 200, o.Payload)
}

func (o *GetOrgEventsTagsOK) GetPayload() *GetOrgEventsTagsOKBody {
	return o.Payload
}

func (o *GetOrgEventsTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrgEventsTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgEventsTagsForbidden creates a GetOrgEventsTagsForbidden with default headers values
func NewGetOrgEventsTagsForbidden() *GetOrgEventsTagsForbidden {
	return &GetOrgEventsTagsForbidden{}
}

/*GetOrgEventsTagsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetOrgEventsTagsForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgEventsTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getOrgEventsTagsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgEventsTagsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgEventsTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgEventsTagsUnprocessableEntity creates a GetOrgEventsTagsUnprocessableEntity with default headers values
func NewGetOrgEventsTagsUnprocessableEntity() *GetOrgEventsTagsUnprocessableEntity {
	return &GetOrgEventsTagsUnprocessableEntity{}
}

/*GetOrgEventsTagsUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetOrgEventsTagsUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetOrgEventsTagsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getOrgEventsTagsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetOrgEventsTagsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgEventsTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgEventsTagsDefault creates a GetOrgEventsTagsDefault with default headers values
func NewGetOrgEventsTagsDefault(code int) *GetOrgEventsTagsDefault {
	return &GetOrgEventsTagsDefault{
		_statusCode: code,
	}
}

/*GetOrgEventsTagsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgEventsTagsDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get org events tags default response
func (o *GetOrgEventsTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgEventsTagsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/events/tags][%d] getOrgEventsTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrgEventsTagsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgEventsTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrgEventsTagsOKBody The list of tags with associated set of values
swagger:model GetOrgEventsTagsOKBody
*/
type GetOrgEventsTagsOKBody struct {

	// data
	// Required: true
	Data interface{} `json:"data"`
}

// Validate validates this get org events tags o k body
func (o *GetOrgEventsTagsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrgEventsTagsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgEventsTagsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgEventsTagsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgEventsTagsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgEventsTagsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

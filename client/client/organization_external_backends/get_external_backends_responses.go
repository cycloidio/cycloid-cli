// Code generated by go-swagger; DO NOT EDIT.

package organization_external_backends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetExternalBackendsReader is a Reader for the GetExternalBackends structure.
type GetExternalBackendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalBackendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalBackendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetExternalBackendsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetExternalBackendsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetExternalBackendsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalBackendsOK creates a GetExternalBackendsOK with default headers values
func NewGetExternalBackendsOK() *GetExternalBackendsOK {
	return &GetExternalBackendsOK{}
}

/*GetExternalBackendsOK handles this case with default header values.

The list of the external backends
*/
type GetExternalBackendsOK struct {
	Payload *GetExternalBackendsOKBody
}

func (o *GetExternalBackendsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends][%d] getExternalBackendsOK  %+v", 200, o.Payload)
}

func (o *GetExternalBackendsOK) GetPayload() *GetExternalBackendsOKBody {
	return o.Payload
}

func (o *GetExternalBackendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetExternalBackendsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalBackendsForbidden creates a GetExternalBackendsForbidden with default headers values
func NewGetExternalBackendsForbidden() *GetExternalBackendsForbidden {
	return &GetExternalBackendsForbidden{}
}

/*GetExternalBackendsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetExternalBackendsForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetExternalBackendsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends][%d] getExternalBackendsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalBackendsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetExternalBackendsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetExternalBackendsUnprocessableEntity creates a GetExternalBackendsUnprocessableEntity with default headers values
func NewGetExternalBackendsUnprocessableEntity() *GetExternalBackendsUnprocessableEntity {
	return &GetExternalBackendsUnprocessableEntity{}
}

/*GetExternalBackendsUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetExternalBackendsUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetExternalBackendsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends][%d] getExternalBackendsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetExternalBackendsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetExternalBackendsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalBackendsDefault creates a GetExternalBackendsDefault with default headers values
func NewGetExternalBackendsDefault(code int) *GetExternalBackendsDefault {
	return &GetExternalBackendsDefault{
		_statusCode: code,
	}
}

/*GetExternalBackendsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetExternalBackendsDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get external backends default response
func (o *GetExternalBackendsDefault) Code() int {
	return o._statusCode
}

func (o *GetExternalBackendsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/external_backends][%d] getExternalBackends default  %+v", o._statusCode, o.Payload)
}

func (o *GetExternalBackendsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetExternalBackendsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetExternalBackendsOKBody get external backends o k body
swagger:model GetExternalBackendsOKBody
*/
type GetExternalBackendsOKBody struct {

	// data
	// Required: true
	Data []*models.ExternalBackend `json:"data"`

	// pagination
	Pagination *models.Pagination `json:"pagination,omitempty"`
}

// Validate validates this get external backends o k body
func (o *GetExternalBackendsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetExternalBackendsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getExternalBackendsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getExternalBackendsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetExternalBackendsOKBody) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getExternalBackendsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetExternalBackendsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetExternalBackendsOKBody) UnmarshalBinary(b []byte) error {
	var res GetExternalBackendsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

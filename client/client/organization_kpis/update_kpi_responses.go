// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

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

// UpdateKpiReader is a Reader for the UpdateKpi structure.
type UpdateKpiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateKpiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateKpiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateKpiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateKpiNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 411:
		result := NewUpdateKpiLengthRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateKpiUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateKpiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateKpiOK creates a UpdateKpiOK with default headers values
func NewUpdateKpiOK() *UpdateKpiOK {
	return &UpdateKpiOK{}
}

/*UpdateKpiOK handles this case with default header values.

Success update
*/
type UpdateKpiOK struct {
	Payload *UpdateKpiOKBody
}

func (o *UpdateKpiOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiOK  %+v", 200, o.Payload)
}

func (o *UpdateKpiOK) GetPayload() *UpdateKpiOKBody {
	return o.Payload
}

func (o *UpdateKpiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateKpiOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateKpiForbidden creates a UpdateKpiForbidden with default headers values
func NewUpdateKpiForbidden() *UpdateKpiForbidden {
	return &UpdateKpiForbidden{}
}

/*UpdateKpiForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type UpdateKpiForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateKpiForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiForbidden  %+v", 403, o.Payload)
}

func (o *UpdateKpiForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateKpiNotFound creates a UpdateKpiNotFound with default headers values
func NewUpdateKpiNotFound() *UpdateKpiNotFound {
	return &UpdateKpiNotFound{}
}

/*UpdateKpiNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type UpdateKpiNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateKpiNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiNotFound  %+v", 404, o.Payload)
}

func (o *UpdateKpiNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateKpiLengthRequired creates a UpdateKpiLengthRequired with default headers values
func NewUpdateKpiLengthRequired() *UpdateKpiLengthRequired {
	return &UpdateKpiLengthRequired{}
}

/*UpdateKpiLengthRequired handles this case with default header values.

The request has a body but it doesn't have a Content-Length header.
*/
type UpdateKpiLengthRequired struct {
}

func (o *UpdateKpiLengthRequired) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiLengthRequired ", 411)
}

func (o *UpdateKpiLengthRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateKpiUnprocessableEntity creates a UpdateKpiUnprocessableEntity with default headers values
func NewUpdateKpiUnprocessableEntity() *UpdateKpiUnprocessableEntity {
	return &UpdateKpiUnprocessableEntity{}
}

/*UpdateKpiUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type UpdateKpiUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *UpdateKpiUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpiUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateKpiUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateKpiDefault creates a UpdateKpiDefault with default headers values
func NewUpdateKpiDefault(code int) *UpdateKpiDefault {
	return &UpdateKpiDefault{
		_statusCode: code,
	}
}

/*UpdateKpiDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type UpdateKpiDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the update kpi default response
func (o *UpdateKpiDefault) Code() int {
	return o._statusCode
}

func (o *UpdateKpiDefault) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] updateKpi default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateKpiDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *UpdateKpiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*UpdateKpiOKBody update kpi o k body
swagger:model UpdateKpiOKBody
*/
type UpdateKpiOKBody struct {

	// data
	// Required: true
	Data *models.KPI `json:"data"`
}

// Validate validates this update kpi o k body
func (o *UpdateKpiOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateKpiOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("updateKpiOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateKpiOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateKpiOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateKpiOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateKpiOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

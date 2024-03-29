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

// GetKpiReader is a Reader for the GetKpi structure.
type GetKpiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKpiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKpiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetKpiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetKpiUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetKpiDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKpiOK creates a GetKpiOK with default headers values
func NewGetKpiOK() *GetKpiOK {
	return &GetKpiOK{}
}

/*GetKpiOK handles this case with default header values.

The KPI
*/
type GetKpiOK struct {
	Payload *GetKpiOKBody
}

func (o *GetKpiOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] getKpiOK  %+v", 200, o.Payload)
}

func (o *GetKpiOK) GetPayload() *GetKpiOKBody {
	return o.Payload
}

func (o *GetKpiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetKpiOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKpiForbidden creates a GetKpiForbidden with default headers values
func NewGetKpiForbidden() *GetKpiForbidden {
	return &GetKpiForbidden{}
}

/*GetKpiForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetKpiForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetKpiForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] getKpiForbidden  %+v", 403, o.Payload)
}

func (o *GetKpiForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetKpiUnprocessableEntity creates a GetKpiUnprocessableEntity with default headers values
func NewGetKpiUnprocessableEntity() *GetKpiUnprocessableEntity {
	return &GetKpiUnprocessableEntity{}
}

/*GetKpiUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetKpiUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetKpiUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] getKpiUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetKpiUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpiUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetKpiDefault creates a GetKpiDefault with default headers values
func NewGetKpiDefault(code int) *GetKpiDefault {
	return &GetKpiDefault{
		_statusCode: code,
	}
}

/*GetKpiDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetKpiDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get kpi default response
func (o *GetKpiDefault) Code() int {
	return o._statusCode
}

func (o *GetKpiDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis/{kpi_canonical}][%d] getKpi default  %+v", o._statusCode, o.Payload)
}

func (o *GetKpiDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpiDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetKpiOKBody get kpi o k body
swagger:model GetKpiOKBody
*/
type GetKpiOKBody struct {

	// data
	// Required: true
	Data *models.KPI `json:"data"`
}

// Validate validates this get kpi o k body
func (o *GetKpiOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetKpiOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getKpiOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getKpiOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetKpiOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKpiOKBody) UnmarshalBinary(b []byte) error {
	var res GetKpiOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

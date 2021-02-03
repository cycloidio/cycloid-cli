// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

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

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// GetKpisAvailableReader is a Reader for the GetKpisAvailable structure.
type GetKpisAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKpisAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKpisAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetKpisAvailableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetKpisAvailableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetKpisAvailableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKpisAvailableOK creates a GetKpisAvailableOK with default headers values
func NewGetKpisAvailableOK() *GetKpisAvailableOK {
	return &GetKpisAvailableOK{}
}

/*GetKpisAvailableOK handles this case with default header values.

The list of available KPIs
*/
type GetKpisAvailableOK struct {
	Payload *GetKpisAvailableOKBody
}

func (o *GetKpisAvailableOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis_available][%d] getKpisAvailableOK  %+v", 200, o.Payload)
}

func (o *GetKpisAvailableOK) GetPayload() *GetKpisAvailableOKBody {
	return o.Payload
}

func (o *GetKpisAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetKpisAvailableOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKpisAvailableForbidden creates a GetKpisAvailableForbidden with default headers values
func NewGetKpisAvailableForbidden() *GetKpisAvailableForbidden {
	return &GetKpisAvailableForbidden{}
}

/*GetKpisAvailableForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetKpisAvailableForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetKpisAvailableForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis_available][%d] getKpisAvailableForbidden  %+v", 403, o.Payload)
}

func (o *GetKpisAvailableForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpisAvailableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKpisAvailableUnprocessableEntity creates a GetKpisAvailableUnprocessableEntity with default headers values
func NewGetKpisAvailableUnprocessableEntity() *GetKpisAvailableUnprocessableEntity {
	return &GetKpisAvailableUnprocessableEntity{}
}

/*GetKpisAvailableUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetKpisAvailableUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetKpisAvailableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis_available][%d] getKpisAvailableUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetKpisAvailableUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpisAvailableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKpisAvailableDefault creates a GetKpisAvailableDefault with default headers values
func NewGetKpisAvailableDefault(code int) *GetKpisAvailableDefault {
	return &GetKpisAvailableDefault{
		_statusCode: code,
	}
}

/*GetKpisAvailableDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetKpisAvailableDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get kpis available default response
func (o *GetKpisAvailableDefault) Code() int {
	return o._statusCode
}

func (o *GetKpisAvailableDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/kpis_available][%d] getKpisAvailable default  %+v", o._statusCode, o.Payload)
}

func (o *GetKpisAvailableDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetKpisAvailableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetKpisAvailableOKBody get kpis available o k body
swagger:model GetKpisAvailableOKBody
*/
type GetKpisAvailableOKBody struct {

	// data
	// Required: true
	Data []*models.KPI `json:"data"`

	// pagination
	Pagination *models.Pagination `json:"pagination,omitempty"`
}

// Validate validates this get kpis available o k body
func (o *GetKpisAvailableOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetKpisAvailableOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getKpisAvailableOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getKpisAvailableOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetKpisAvailableOKBody) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getKpisAvailableOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetKpisAvailableOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetKpisAvailableOKBody) UnmarshalBinary(b []byte) error {
	var res GetKpisAvailableOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

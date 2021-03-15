// Code generated by go-swagger; DO NOT EDIT.

package cycloid

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

// GetCountriesReader is a Reader for the GetCountries structure.
type GetCountriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCountriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCountriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCountriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCountriesOK creates a GetCountriesOK with default headers values
func NewGetCountriesOK() *GetCountriesOK {
	return &GetCountriesOK{}
}

/*GetCountriesOK handles this case with default header values.

Cycloid supported countries
*/
type GetCountriesOK struct {
	Payload *GetCountriesOKBody
}

func (o *GetCountriesOK) Error() string {
	return fmt.Sprintf("[GET /countries][%d] getCountriesOK  %+v", 200, o.Payload)
}

func (o *GetCountriesOK) GetPayload() *GetCountriesOKBody {
	return o.Payload
}

func (o *GetCountriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCountriesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCountriesDefault creates a GetCountriesDefault with default headers values
func NewGetCountriesDefault(code int) *GetCountriesDefault {
	return &GetCountriesDefault{
		_statusCode: code,
	}
}

/*GetCountriesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetCountriesDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get countries default response
func (o *GetCountriesDefault) Code() int {
	return o._statusCode
}

func (o *GetCountriesDefault) Error() string {
	return fmt.Sprintf("[GET /countries][%d] getCountries default  %+v", o._statusCode, o.Payload)
}

func (o *GetCountriesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetCountriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCountriesOKBody get countries o k body
swagger:model GetCountriesOKBody
*/
type GetCountriesOKBody struct {

	// data
	// Required: true
	Data []*models.Country `json:"data"`
}

// Validate validates this get countries o k body
func (o *GetCountriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCountriesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getCountriesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getCountriesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCountriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCountriesOKBody) UnmarshalBinary(b []byte) error {
	var res GetCountriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
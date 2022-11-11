// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// GetAncestorsReader is a Reader for the GetAncestors structure.
type GetAncestorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAncestorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAncestorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAncestorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAncestorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAncestorsOK creates a GetAncestorsOK with default headers values
func NewGetAncestorsOK() *GetAncestorsOK {
	return &GetAncestorsOK{}
}

/*GetAncestorsOK handles this case with default header values.

Get all the ancestors between the Organization and the User with the shortest path. 0 index is the parent and n is the searched child
*/
type GetAncestorsOK struct {
	Payload *GetAncestorsOKBody
}

func (o *GetAncestorsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/ancestors][%d] getAncestorsOK  %+v", 200, o.Payload)
}

func (o *GetAncestorsOK) GetPayload() *GetAncestorsOKBody {
	return o.Payload
}

func (o *GetAncestorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAncestorsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAncestorsUnauthorized creates a GetAncestorsUnauthorized with default headers values
func NewGetAncestorsUnauthorized() *GetAncestorsUnauthorized {
	return &GetAncestorsUnauthorized{}
}

/*GetAncestorsUnauthorized handles this case with default header values.

The user cannot be authenticated with the credentials which she/he has used.
*/
type GetAncestorsUnauthorized struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetAncestorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/ancestors][%d] getAncestorsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAncestorsUnauthorized) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAncestorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAncestorsDefault creates a GetAncestorsDefault with default headers values
func NewGetAncestorsDefault(code int) *GetAncestorsDefault {
	return &GetAncestorsDefault{
		_statusCode: code,
	}
}

/*GetAncestorsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetAncestorsDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get ancestors default response
func (o *GetAncestorsDefault) Code() int {
	return o._statusCode
}

func (o *GetAncestorsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/ancestors][%d] getAncestors default  %+v", o._statusCode, o.Payload)
}

func (o *GetAncestorsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetAncestorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetAncestorsOKBody get ancestors o k body
swagger:model GetAncestorsOKBody
*/
type GetAncestorsOKBody struct {

	// data
	// Required: true
	Data []*models.Organization `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get ancestors o k body
func (o *GetAncestorsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetAncestorsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getAncestorsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getAncestorsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetAncestorsOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getAncestorsOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAncestorsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAncestorsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAncestorsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAncestorsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
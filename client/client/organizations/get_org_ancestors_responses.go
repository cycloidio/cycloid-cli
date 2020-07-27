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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetOrgAncestorsReader is a Reader for the GetOrgAncestors structure.
type GetOrgAncestorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgAncestorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgAncestorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrgAncestorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOrgAncestorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrgAncestorsOK creates a GetOrgAncestorsOK with default headers values
func NewGetOrgAncestorsOK() *GetOrgAncestorsOK {
	return &GetOrgAncestorsOK{}
}

/*GetOrgAncestorsOK handles this case with default header values.

Get all the ancestors between the Organization and the User with the shortest path. 0 index is the parent and n is the searched child
*/
type GetOrgAncestorsOK struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *GetOrgAncestorsOKBody
}

func (o *GetOrgAncestorsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/ancestors][%d] getOrgAncestorsOK  %+v", 200, o.Payload)
}

func (o *GetOrgAncestorsOK) GetPayload() *GetOrgAncestorsOKBody {
	return o.Payload
}

func (o *GetOrgAncestorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Length
	contentLength, err := swag.ConvertInt64(response.GetHeader("Content-Length"))
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", response.GetHeader("Content-Length"))
	}
	o.ContentLength = contentLength

	o.Payload = new(GetOrgAncestorsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgAncestorsUnauthorized creates a GetOrgAncestorsUnauthorized with default headers values
func NewGetOrgAncestorsUnauthorized() *GetOrgAncestorsUnauthorized {
	return &GetOrgAncestorsUnauthorized{}
}

/*GetOrgAncestorsUnauthorized handles this case with default header values.

The user cannot be authenticated with the credentials which she/he has used.
*/
type GetOrgAncestorsUnauthorized struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetOrgAncestorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/ancestors][%d] getOrgAncestorsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgAncestorsUnauthorized) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgAncestorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOrgAncestorsDefault creates a GetOrgAncestorsDefault with default headers values
func NewGetOrgAncestorsDefault(code int) *GetOrgAncestorsDefault {
	return &GetOrgAncestorsDefault{
		_statusCode: code,
	}
}

/*GetOrgAncestorsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetOrgAncestorsDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get org ancestors default response
func (o *GetOrgAncestorsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrgAncestorsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/ancestors][%d] getOrgAncestors default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrgAncestorsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetOrgAncestorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrgAncestorsOKBody get org ancestors o k body
swagger:model GetOrgAncestorsOKBody
*/
type GetOrgAncestorsOKBody struct {

	// data
	// Required: true
	Data []*models.OrganizationBasicInfo `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get org ancestors o k body
func (o *GetOrgAncestorsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetOrgAncestorsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrgAncestorsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrgAncestorsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrgAncestorsOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getOrgAncestorsOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrgAncestorsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrgAncestorsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrgAncestorsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrgAncestorsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

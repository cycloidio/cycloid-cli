// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines

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

// GetPipelinesReader is a Reader for the GetPipelines structure.
type GetPipelinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPipelinesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetPipelinesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPipelinesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPipelinesOK creates a GetPipelinesOK with default headers values
func NewGetPipelinesOK() *GetPipelinesOK {
	return &GetPipelinesOK{}
}

/*GetPipelinesOK handles this case with default header values.

List of all the pipelines which authenticated user has access to.
*/
type GetPipelinesOK struct {
	Payload *GetPipelinesOKBody
}

func (o *GetPipelinesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesOK  %+v", 200, o.Payload)
}

func (o *GetPipelinesOK) GetPayload() *GetPipelinesOKBody {
	return o.Payload
}

func (o *GetPipelinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPipelinesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelinesNotFound creates a GetPipelinesNotFound with default headers values
func NewGetPipelinesNotFound() *GetPipelinesNotFound {
	return &GetPipelinesNotFound{}
}

/*GetPipelinesNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetPipelinesNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetPipelinesNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelinesNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelinesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelinesUnprocessableEntity creates a GetPipelinesUnprocessableEntity with default headers values
func NewGetPipelinesUnprocessableEntity() *GetPipelinesUnprocessableEntity {
	return &GetPipelinesUnprocessableEntity{}
}

/*GetPipelinesUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetPipelinesUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetPipelinesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelinesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetPipelinesUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelinesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelinesDefault creates a GetPipelinesDefault with default headers values
func NewGetPipelinesDefault(code int) *GetPipelinesDefault {
	return &GetPipelinesDefault{
		_statusCode: code,
	}
}

/*GetPipelinesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetPipelinesDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get pipelines default response
func (o *GetPipelinesDefault) Code() int {
	return o._statusCode
}

func (o *GetPipelinesDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/pipelines][%d] getPipelines default  %+v", o._statusCode, o.Payload)
}

func (o *GetPipelinesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelinesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetPipelinesOKBody get pipelines o k body
swagger:model GetPipelinesOKBody
*/
type GetPipelinesOKBody struct {

	// data
	// Required: true
	Data []*models.Pipeline `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get pipelines o k body
func (o *GetPipelinesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetPipelinesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPipelinesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPipelinesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPipelinesOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getPipelinesOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPipelinesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPipelinesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPipelinesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPipelinesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

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

// GetProjectPipelinesReader is a Reader for the GetProjectPipelines structure.
type GetProjectPipelinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectPipelinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectPipelinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProjectPipelinesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetProjectPipelinesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetProjectPipelinesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectPipelinesOK creates a GetProjectPipelinesOK with default headers values
func NewGetProjectPipelinesOK() *GetProjectPipelinesOK {
	return &GetProjectPipelinesOK{}
}

/*GetProjectPipelinesOK handles this case with default header values.

List of the pipelines which authenticated user has access to.
*/
type GetProjectPipelinesOK struct {
	Payload *GetProjectPipelinesOKBody
}

func (o *GetProjectPipelinesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] getProjectPipelinesOK  %+v", 200, o.Payload)
}

func (o *GetProjectPipelinesOK) GetPayload() *GetProjectPipelinesOKBody {
	return o.Payload
}

func (o *GetProjectPipelinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProjectPipelinesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectPipelinesNotFound creates a GetProjectPipelinesNotFound with default headers values
func NewGetProjectPipelinesNotFound() *GetProjectPipelinesNotFound {
	return &GetProjectPipelinesNotFound{}
}

/*GetProjectPipelinesNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetProjectPipelinesNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetProjectPipelinesNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] getProjectPipelinesNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectPipelinesNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetProjectPipelinesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectPipelinesUnprocessableEntity creates a GetProjectPipelinesUnprocessableEntity with default headers values
func NewGetProjectPipelinesUnprocessableEntity() *GetProjectPipelinesUnprocessableEntity {
	return &GetProjectPipelinesUnprocessableEntity{}
}

/*GetProjectPipelinesUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetProjectPipelinesUnprocessableEntity struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetProjectPipelinesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] getProjectPipelinesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetProjectPipelinesUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetProjectPipelinesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectPipelinesDefault creates a GetProjectPipelinesDefault with default headers values
func NewGetProjectPipelinesDefault(code int) *GetProjectPipelinesDefault {
	return &GetProjectPipelinesDefault{
		_statusCode: code,
	}
}

/*GetProjectPipelinesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetProjectPipelinesDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get project pipelines default response
func (o *GetProjectPipelinesDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectPipelinesDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines][%d] getProjectPipelines default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectPipelinesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetProjectPipelinesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetProjectPipelinesOKBody get project pipelines o k body
swagger:model GetProjectPipelinesOKBody
*/
type GetProjectPipelinesOKBody struct {

	// data
	// Required: true
	Data []*models.Pipeline `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get project pipelines o k body
func (o *GetProjectPipelinesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetProjectPipelinesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getProjectPipelinesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getProjectPipelinesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetProjectPipelinesOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getProjectPipelinesOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProjectPipelinesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProjectPipelinesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProjectPipelinesOKBody) UnmarshalBinary(b []byte) error {
	var res GetProjectPipelinesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

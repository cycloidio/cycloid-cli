// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs_build

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

// CreateBuildReader is a Reader for the CreateBuild structure.
type CreateBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBuildOK creates a CreateBuildOK with default headers values
func NewCreateBuildOK() *CreateBuildOK {
	return &CreateBuildOK{}
}

/*CreateBuildOK handles this case with default header values.

Create a new build for the pipeline's job and returns its details
*/
type CreateBuildOK struct {
	Payload *CreateBuildOKBody
}

func (o *CreateBuildOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] createBuildOK  %+v", 200, o.Payload)
}

func (o *CreateBuildOK) GetPayload() *CreateBuildOKBody {
	return o.Payload
}

func (o *CreateBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateBuildOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBuildForbidden creates a CreateBuildForbidden with default headers values
func NewCreateBuildForbidden() *CreateBuildForbidden {
	return &CreateBuildForbidden{}
}

/*CreateBuildForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type CreateBuildForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] createBuildForbidden  %+v", 403, o.Payload)
}

func (o *CreateBuildForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBuildNotFound creates a CreateBuildNotFound with default headers values
func NewCreateBuildNotFound() *CreateBuildNotFound {
	return &CreateBuildNotFound{}
}

/*CreateBuildNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type CreateBuildNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *CreateBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] createBuildNotFound  %+v", 404, o.Payload)
}

func (o *CreateBuildNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBuildDefault creates a CreateBuildDefault with default headers values
func NewCreateBuildDefault(code int) *CreateBuildDefault {
	return &CreateBuildDefault{
		_statusCode: code,
	}
}

/*CreateBuildDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type CreateBuildDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the create build default response
func (o *CreateBuildDefault) Code() int {
	return o._statusCode
}

func (o *CreateBuildDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] createBuild default  %+v", o._statusCode, o.Payload)
}

func (o *CreateBuildDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *CreateBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreateBuildOKBody create build o k body
swagger:model CreateBuildOKBody
*/
type CreateBuildOKBody struct {

	// data
	// Required: true
	Data *models.Build `json:"data"`
}

// Validate validates this create build o k body
func (o *CreateBuildOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateBuildOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("createBuildOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createBuildOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateBuildOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBuildOKBody) UnmarshalBinary(b []byte) error {
	var res CreateBuildOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

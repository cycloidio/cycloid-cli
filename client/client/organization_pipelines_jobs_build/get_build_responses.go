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

// GetBuildReader is a Reader for the GetBuild structure.
type GetBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildOK creates a GetBuildOK with default headers values
func NewGetBuildOK() *GetBuildOK {
	return &GetBuildOK{}
}

/*GetBuildOK handles this case with default header values.

The information of the build which has the specified id.
*/
type GetBuildOK struct {
	Payload *GetBuildOKBody
}

func (o *GetBuildOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildOK  %+v", 200, o.Payload)
}

func (o *GetBuildOK) GetPayload() *GetBuildOKBody {
	return o.Payload
}

func (o *GetBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBuildOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildForbidden creates a GetBuildForbidden with default headers values
func NewGetBuildForbidden() *GetBuildForbidden {
	return &GetBuildForbidden{}
}

/*GetBuildForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetBuildForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetBuildForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildForbidden  %+v", 403, o.Payload)
}

func (o *GetBuildForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildNotFound creates a GetBuildNotFound with default headers values
func NewGetBuildNotFound() *GetBuildNotFound {
	return &GetBuildNotFound{}
}

/*GetBuildNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetBuildNotFound struct {
	Payload *models.ErrorPayload
}

func (o *GetBuildNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildNotFound  %+v", 404, o.Payload)
}

func (o *GetBuildNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildDefault creates a GetBuildDefault with default headers values
func NewGetBuildDefault(code int) *GetBuildDefault {
	return &GetBuildDefault{
		_statusCode: code,
	}
}

/*GetBuildDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetBuildDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get build default response
func (o *GetBuildDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuild default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBuildOKBody get build o k body
swagger:model GetBuildOKBody
*/
type GetBuildOKBody struct {

	// data
	// Required: true
	Data *models.Build `json:"data"`
}

// Validate validates this get build o k body
func (o *GetBuildOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBuildOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBuildOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBuildOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBuildOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBuildOKBody) UnmarshalBinary(b []byte) error {
	var res GetBuildOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs_build

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

// GetBuildsReader is a Reader for the GetBuilds structure.
type GetBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBuildsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetBuildsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBuildsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildsOK creates a GetBuildsOK with default headers values
func NewGetBuildsOK() *GetBuildsOK {
	return &GetBuildsOK{}
}

/*GetBuildsOK handles this case with default header values.

List the pipeline job's builds which authenticated user has access to.
*/
type GetBuildsOK struct {
	Payload *GetBuildsOKBody
}

func (o *GetBuildsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] getBuildsOK  %+v", 200, o.Payload)
}

func (o *GetBuildsOK) GetPayload() *GetBuildsOKBody {
	return o.Payload
}

func (o *GetBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBuildsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildsForbidden creates a GetBuildsForbidden with default headers values
func NewGetBuildsForbidden() *GetBuildsForbidden {
	return &GetBuildsForbidden{}
}

/*GetBuildsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetBuildsForbidden struct {
	Payload *models.ErrorPayload
}

func (o *GetBuildsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] getBuildsForbidden  %+v", 403, o.Payload)
}

func (o *GetBuildsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildsNotFound creates a GetBuildsNotFound with default headers values
func NewGetBuildsNotFound() *GetBuildsNotFound {
	return &GetBuildsNotFound{}
}

/*GetBuildsNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetBuildsNotFound struct {
	Payload *models.ErrorPayload
}

func (o *GetBuildsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] getBuildsNotFound  %+v", 404, o.Payload)
}

func (o *GetBuildsNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildsUnprocessableEntity creates a GetBuildsUnprocessableEntity with default headers values
func NewGetBuildsUnprocessableEntity() *GetBuildsUnprocessableEntity {
	return &GetBuildsUnprocessableEntity{}
}

/*GetBuildsUnprocessableEntity handles this case with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetBuildsUnprocessableEntity struct {
	Payload *models.ErrorPayload
}

func (o *GetBuildsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] getBuildsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetBuildsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildsDefault creates a GetBuildsDefault with default headers values
func NewGetBuildsDefault(code int) *GetBuildsDefault {
	return &GetBuildsDefault{
		_statusCode: code,
	}
}

/*GetBuildsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetBuildsDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get builds default response
func (o *GetBuildsDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds][%d] getBuilds default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBuildsOKBody get builds o k body
swagger:model GetBuildsOKBody
*/
type GetBuildsOKBody struct {

	// data
	// Required: true
	Data []*models.Build `json:"data"`

	// pagination concourse
	// Required: true
	PaginationConcourse *models.PaginationConcourse `json:"pagination_concourse"`
}

// Validate validates this get builds o k body
func (o *GetBuildsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePaginationConcourse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBuildsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBuildsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getBuildsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetBuildsOKBody) validatePaginationConcourse(formats strfmt.Registry) error {

	if err := validate.Required("getBuildsOK"+"."+"pagination_concourse", "body", o.PaginationConcourse); err != nil {
		return err
	}

	if o.PaginationConcourse != nil {
		if err := o.PaginationConcourse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBuildsOK" + "." + "pagination_concourse")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBuildsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBuildsOKBody) UnmarshalBinary(b []byte) error {
	var res GetBuildsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

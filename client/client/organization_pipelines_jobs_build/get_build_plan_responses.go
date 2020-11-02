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

// GetBuildPlanReader is a Reader for the GetBuildPlan structure.
type GetBuildPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBuildPlanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBuildPlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildPlanOK creates a GetBuildPlanOK with default headers values
func NewGetBuildPlanOK() *GetBuildPlanOK {
	return &GetBuildPlanOK{}
}

/*GetBuildPlanOK handles this case with default header values.

The information of the build's plan which has the specified id.
*/
type GetBuildPlanOK struct {
	Payload *GetBuildPlanOKBody
}

func (o *GetBuildPlanOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/plan][%d] getBuildPlanOK  %+v", 200, o.Payload)
}

func (o *GetBuildPlanOK) GetPayload() *GetBuildPlanOKBody {
	return o.Payload
}

func (o *GetBuildPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBuildPlanOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildPlanForbidden creates a GetBuildPlanForbidden with default headers values
func NewGetBuildPlanForbidden() *GetBuildPlanForbidden {
	return &GetBuildPlanForbidden{}
}

/*GetBuildPlanForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetBuildPlanForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetBuildPlanForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/plan][%d] getBuildPlanForbidden  %+v", 403, o.Payload)
}

func (o *GetBuildPlanForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildPlanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildPlanNotFound creates a GetBuildPlanNotFound with default headers values
func NewGetBuildPlanNotFound() *GetBuildPlanNotFound {
	return &GetBuildPlanNotFound{}
}

/*GetBuildPlanNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetBuildPlanNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetBuildPlanNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/plan][%d] getBuildPlanNotFound  %+v", 404, o.Payload)
}

func (o *GetBuildPlanNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildPlanDefault creates a GetBuildPlanDefault with default headers values
func NewGetBuildPlanDefault(code int) *GetBuildPlanDefault {
	return &GetBuildPlanDefault{
		_statusCode: code,
	}
}

/*GetBuildPlanDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetBuildPlanDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get build plan default response
func (o *GetBuildPlanDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildPlanDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/plan][%d] getBuildPlan default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildPlanDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildPlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBuildPlanOKBody get build plan o k body
swagger:model GetBuildPlanOKBody
*/
type GetBuildPlanOKBody struct {

	// data
	// Required: true
	Data *models.PublicPlan `json:"data"`
}

// Validate validates this get build plan o k body
func (o *GetBuildPlanOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBuildPlanOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBuildPlanOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBuildPlanOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBuildPlanOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBuildPlanOKBody) UnmarshalBinary(b []byte) error {
	var res GetBuildPlanOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

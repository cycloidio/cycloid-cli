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

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetBuildResourcesReader is a Reader for the GetBuildResources structure.
type GetBuildResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBuildResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBuildResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildResourcesOK creates a GetBuildResourcesOK with default headers values
func NewGetBuildResourcesOK() *GetBuildResourcesOK {
	return &GetBuildResourcesOK{}
}

/*GetBuildResourcesOK handles this case with default header values.

The resources of the build's which has the specified id.
*/
type GetBuildResourcesOK struct {
	Payload *GetBuildResourcesOKBody
}

func (o *GetBuildResourcesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/resources][%d] getBuildResourcesOK  %+v", 200, o.Payload)
}

func (o *GetBuildResourcesOK) GetPayload() *GetBuildResourcesOKBody {
	return o.Payload
}

func (o *GetBuildResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBuildResourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildResourcesForbidden creates a GetBuildResourcesForbidden with default headers values
func NewGetBuildResourcesForbidden() *GetBuildResourcesForbidden {
	return &GetBuildResourcesForbidden{}
}

/*GetBuildResourcesForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetBuildResourcesForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetBuildResourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/resources][%d] getBuildResourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetBuildResourcesForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildResourcesNotFound creates a GetBuildResourcesNotFound with default headers values
func NewGetBuildResourcesNotFound() *GetBuildResourcesNotFound {
	return &GetBuildResourcesNotFound{}
}

/*GetBuildResourcesNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetBuildResourcesNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetBuildResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/resources][%d] getBuildResourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetBuildResourcesNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildResourcesDefault creates a GetBuildResourcesDefault with default headers values
func NewGetBuildResourcesDefault(code int) *GetBuildResourcesDefault {
	return &GetBuildResourcesDefault{
		_statusCode: code,
	}
}

/*GetBuildResourcesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetBuildResourcesDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get build resources default response
func (o *GetBuildResourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildResourcesDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/resources][%d] getBuildResources default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildResourcesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBuildResourcesOKBody get build resources o k body
swagger:model GetBuildResourcesOKBody
*/
type GetBuildResourcesOKBody struct {

	// data
	// Required: true
	Data *models.BuildInputsOutputs `json:"data"`
}

// Validate validates this get build resources o k body
func (o *GetBuildResourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBuildResourcesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBuildResourcesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBuildResourcesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBuildResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBuildResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetBuildResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

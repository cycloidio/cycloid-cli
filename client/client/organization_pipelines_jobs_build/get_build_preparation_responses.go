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

// GetBuildPreparationReader is a Reader for the GetBuildPreparation structure.
type GetBuildPreparationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildPreparationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildPreparationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBuildPreparationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildPreparationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBuildPreparationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildPreparationOK creates a GetBuildPreparationOK with default headers values
func NewGetBuildPreparationOK() *GetBuildPreparationOK {
	return &GetBuildPreparationOK{}
}

/*GetBuildPreparationOK handles this case with default header values.

Return the Preparation
*/
type GetBuildPreparationOK struct {
	Payload *GetBuildPreparationOKBody
}

func (o *GetBuildPreparationOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/preparation][%d] getBuildPreparationOK  %+v", 200, o.Payload)
}

func (o *GetBuildPreparationOK) GetPayload() *GetBuildPreparationOKBody {
	return o.Payload
}

func (o *GetBuildPreparationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBuildPreparationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildPreparationForbidden creates a GetBuildPreparationForbidden with default headers values
func NewGetBuildPreparationForbidden() *GetBuildPreparationForbidden {
	return &GetBuildPreparationForbidden{}
}

/*GetBuildPreparationForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetBuildPreparationForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetBuildPreparationForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/preparation][%d] getBuildPreparationForbidden  %+v", 403, o.Payload)
}

func (o *GetBuildPreparationForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildPreparationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildPreparationNotFound creates a GetBuildPreparationNotFound with default headers values
func NewGetBuildPreparationNotFound() *GetBuildPreparationNotFound {
	return &GetBuildPreparationNotFound{}
}

/*GetBuildPreparationNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetBuildPreparationNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *GetBuildPreparationNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/preparation][%d] getBuildPreparationNotFound  %+v", 404, o.Payload)
}

func (o *GetBuildPreparationNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildPreparationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildPreparationDefault creates a GetBuildPreparationDefault with default headers values
func NewGetBuildPreparationDefault(code int) *GetBuildPreparationDefault {
	return &GetBuildPreparationDefault{
		_statusCode: code,
	}
}

/*GetBuildPreparationDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetBuildPreparationDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get build preparation default response
func (o *GetBuildPreparationDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildPreparationDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/preparation][%d] getBuildPreparation default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildPreparationDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildPreparationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBuildPreparationOKBody get build preparation o k body
swagger:model GetBuildPreparationOKBody
*/
type GetBuildPreparationOKBody struct {

	// data
	// Required: true
	Data *models.Preparation `json:"data"`
}

// Validate validates this get build preparation o k body
func (o *GetBuildPreparationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBuildPreparationOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBuildPreparationOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBuildPreparationOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBuildPreparationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBuildPreparationOKBody) UnmarshalBinary(b []byte) error {
	var res GetBuildPreparationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

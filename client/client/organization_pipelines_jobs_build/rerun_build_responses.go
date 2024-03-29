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

// RerunBuildReader is a Reader for the RerunBuild structure.
type RerunBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RerunBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRerunBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRerunBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRerunBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRerunBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRerunBuildOK creates a RerunBuildOK with default headers values
func NewRerunBuildOK() *RerunBuildOK {
	return &RerunBuildOK{}
}

/*RerunBuildOK handles this case with default header values.

Returns the new build created from the specified build ID.
*/
type RerunBuildOK struct {
	Payload *RerunBuildOKBody
}

func (o *RerunBuildOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildOK  %+v", 200, o.Payload)
}

func (o *RerunBuildOK) GetPayload() *RerunBuildOKBody {
	return o.Payload
}

func (o *RerunBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RerunBuildOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRerunBuildForbidden creates a RerunBuildForbidden with default headers values
func NewRerunBuildForbidden() *RerunBuildForbidden {
	return &RerunBuildForbidden{}
}

/*RerunBuildForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type RerunBuildForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *RerunBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildForbidden  %+v", 403, o.Payload)
}

func (o *RerunBuildForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RerunBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRerunBuildNotFound creates a RerunBuildNotFound with default headers values
func NewRerunBuildNotFound() *RerunBuildNotFound {
	return &RerunBuildNotFound{}
}

/*RerunBuildNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type RerunBuildNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *RerunBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildNotFound  %+v", 404, o.Payload)
}

func (o *RerunBuildNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RerunBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRerunBuildDefault creates a RerunBuildDefault with default headers values
func NewRerunBuildDefault(code int) *RerunBuildDefault {
	return &RerunBuildDefault{
		_statusCode: code,
	}
}

/*RerunBuildDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type RerunBuildDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the rerun build default response
func (o *RerunBuildDefault) Code() int {
	return o._statusCode
}

func (o *RerunBuildDefault) Error() string {
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuild default  %+v", o._statusCode, o.Payload)
}

func (o *RerunBuildDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RerunBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*RerunBuildOKBody rerun build o k body
swagger:model RerunBuildOKBody
*/
type RerunBuildOKBody struct {

	// data
	// Required: true
	Data *models.Build `json:"data"`
}

// Validate validates this rerun build o k body
func (o *RerunBuildOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RerunBuildOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("rerunBuildOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rerunBuildOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RerunBuildOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RerunBuildOKBody) UnmarshalBinary(b []byte) error {
	var res RerunBuildOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

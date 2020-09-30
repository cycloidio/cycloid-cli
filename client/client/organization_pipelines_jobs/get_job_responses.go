// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines_jobs

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

// GetJobReader is a Reader for the GetJob structure.
type GetJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetJobOK creates a GetJobOK with default headers values
func NewGetJobOK() *GetJobOK {
	return &GetJobOK{}
}

/*GetJobOK handles this case with default header values.

The configuration of the job which has the specified name.
*/
type GetJobOK struct {
	Payload *GetJobOKBody
}

func (o *GetJobOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}][%d] getJobOK  %+v", 200, o.Payload)
}

func (o *GetJobOK) GetPayload() *GetJobOKBody {
	return o.Payload
}

func (o *GetJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetJobOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobForbidden creates a GetJobForbidden with default headers values
func NewGetJobForbidden() *GetJobForbidden {
	return &GetJobForbidden{}
}

/*GetJobForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetJobForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetJobForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}][%d] getJobForbidden  %+v", 403, o.Payload)
}

func (o *GetJobForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetJobNotFound creates a GetJobNotFound with default headers values
func NewGetJobNotFound() *GetJobNotFound {
	return &GetJobNotFound{}
}

/*GetJobNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetJobNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetJobNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}][%d] getJobNotFound  %+v", 404, o.Payload)
}

func (o *GetJobNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetJobDefault creates a GetJobDefault with default headers values
func NewGetJobDefault(code int) *GetJobDefault {
	return &GetJobDefault{
		_statusCode: code,
	}
}

/*GetJobDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetJobDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get job default response
func (o *GetJobDefault) Code() int {
	return o._statusCode
}

func (o *GetJobDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}][%d] getJob default  %+v", o._statusCode, o.Payload)
}

func (o *GetJobDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetJobOKBody get job o k body
swagger:model GetJobOKBody
*/
type GetJobOKBody struct {

	// data
	// Required: true
	Data *models.Job `json:"data"`
}

// Validate validates this get job o k body
func (o *GetJobOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetJobOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getJobOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getJobOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetJobOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetJobOKBody) UnmarshalBinary(b []byte) error {
	var res GetJobOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

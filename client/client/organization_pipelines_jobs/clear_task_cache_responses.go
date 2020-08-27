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

// ClearTaskCacheReader is a Reader for the ClearTaskCache structure.
type ClearTaskCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearTaskCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClearTaskCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewClearTaskCacheForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClearTaskCacheNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewClearTaskCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClearTaskCacheOK creates a ClearTaskCacheOK with default headers values
func NewClearTaskCacheOK() *ClearTaskCacheOK {
	return &ClearTaskCacheOK{}
}

/*ClearTaskCacheOK handles this case with default header values.

Cache has been cleared.
*/
type ClearTaskCacheOK struct {
	Payload *ClearTaskCacheOKBody
}

func (o *ClearTaskCacheOK) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/tasks/{step_name}/cache][%d] clearTaskCacheOK  %+v", 200, o.Payload)
}

func (o *ClearTaskCacheOK) GetPayload() *ClearTaskCacheOKBody {
	return o.Payload
}

func (o *ClearTaskCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ClearTaskCacheOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearTaskCacheForbidden creates a ClearTaskCacheForbidden with default headers values
func NewClearTaskCacheForbidden() *ClearTaskCacheForbidden {
	return &ClearTaskCacheForbidden{}
}

/*ClearTaskCacheForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type ClearTaskCacheForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *ClearTaskCacheForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/tasks/{step_name}/cache][%d] clearTaskCacheForbidden  %+v", 403, o.Payload)
}

func (o *ClearTaskCacheForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ClearTaskCacheForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewClearTaskCacheNotFound creates a ClearTaskCacheNotFound with default headers values
func NewClearTaskCacheNotFound() *ClearTaskCacheNotFound {
	return &ClearTaskCacheNotFound{}
}

/*ClearTaskCacheNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type ClearTaskCacheNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *ClearTaskCacheNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/tasks/{step_name}/cache][%d] clearTaskCacheNotFound  %+v", 404, o.Payload)
}

func (o *ClearTaskCacheNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ClearTaskCacheNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewClearTaskCacheDefault creates a ClearTaskCacheDefault with default headers values
func NewClearTaskCacheDefault(code int) *ClearTaskCacheDefault {
	return &ClearTaskCacheDefault{
		_statusCode: code,
	}
}

/*ClearTaskCacheDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type ClearTaskCacheDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the clear task cache default response
func (o *ClearTaskCacheDefault) Code() int {
	return o._statusCode
}

func (o *ClearTaskCacheDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/tasks/{step_name}/cache][%d] clearTaskCache default  %+v", o._statusCode, o.Payload)
}

func (o *ClearTaskCacheDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *ClearTaskCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ClearTaskCacheOKBody clear task cache o k body
swagger:model ClearTaskCacheOKBody
*/
type ClearTaskCacheOKBody struct {

	// data
	// Required: true
	Data *models.ClearTaskCache `json:"data"`
}

// Validate validates this clear task cache o k body
func (o *ClearTaskCacheOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClearTaskCacheOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("clearTaskCacheOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clearTaskCacheOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClearTaskCacheOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClearTaskCacheOKBody) UnmarshalBinary(b []byte) error {
	var res ClearTaskCacheOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
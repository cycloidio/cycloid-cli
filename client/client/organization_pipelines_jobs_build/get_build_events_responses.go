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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/youdeploy-cli/client/models"
)

// GetBuildEventsReader is a Reader for the GetBuildEvents structure.
type GetBuildEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBuildEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBuildEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildEventsOK creates a GetBuildEventsOK with default headers values
func NewGetBuildEventsOK() *GetBuildEventsOK {
	return &GetBuildEventsOK{}
}

/*GetBuildEventsOK handles this case with default header values.

Stream is starting
*/
type GetBuildEventsOK struct {
	ContentType string
}

func (o *GetBuildEventsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/events][%d] getBuildEventsOK ", 200)
}

func (o *GetBuildEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	return nil
}

// NewGetBuildEventsForbidden creates a GetBuildEventsForbidden with default headers values
func NewGetBuildEventsForbidden() *GetBuildEventsForbidden {
	return &GetBuildEventsForbidden{}
}

/*GetBuildEventsForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetBuildEventsForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetBuildEventsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/events][%d] getBuildEventsForbidden  %+v", 403, o.Payload)
}

func (o *GetBuildEventsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildEventsNotFound creates a GetBuildEventsNotFound with default headers values
func NewGetBuildEventsNotFound() *GetBuildEventsNotFound {
	return &GetBuildEventsNotFound{}
}

/*GetBuildEventsNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetBuildEventsNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetBuildEventsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/events][%d] getBuildEventsNotFound  %+v", 404, o.Payload)
}

func (o *GetBuildEventsNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildEventsDefault creates a GetBuildEventsDefault with default headers values
func NewGetBuildEventsDefault(code int) *GetBuildEventsDefault {
	return &GetBuildEventsDefault{
		_statusCode: code,
	}
}

/*GetBuildEventsDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetBuildEventsDefault struct {
	_statusCode int

	Payload *models.ErrorPayload
}

// Code gets the status code for the get build events default response
func (o *GetBuildEventsDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildEventsDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}/events][%d] getBuildEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildEventsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

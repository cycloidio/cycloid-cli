// Code generated by go-swagger; DO NOT EDIT.

package component_pipelines_jobs_builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cycloidio/cycloid-cli/client/models"
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

/*
RerunBuildOK describes a response with status code 200, with default header values.

Returns the new build created from the specified build ID.
*/
type RerunBuildOK struct {
	Payload *RerunBuildOKBody
}

// IsSuccess returns true when this rerun build o k response has a 2xx status code
func (o *RerunBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rerun build o k response has a 3xx status code
func (o *RerunBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rerun build o k response has a 4xx status code
func (o *RerunBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rerun build o k response has a 5xx status code
func (o *RerunBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rerun build o k response a status code equal to that given
func (o *RerunBuildOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rerun build o k response
func (o *RerunBuildOK) Code() int {
	return 200
}

func (o *RerunBuildOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildOK %s", 200, payload)
}

func (o *RerunBuildOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildOK %s", 200, payload)
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

/*
RerunBuildForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type RerunBuildForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this rerun build forbidden response has a 2xx status code
func (o *RerunBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rerun build forbidden response has a 3xx status code
func (o *RerunBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rerun build forbidden response has a 4xx status code
func (o *RerunBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this rerun build forbidden response has a 5xx status code
func (o *RerunBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this rerun build forbidden response a status code equal to that given
func (o *RerunBuildForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the rerun build forbidden response
func (o *RerunBuildForbidden) Code() int {
	return 403
}

func (o *RerunBuildForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildForbidden %s", 403, payload)
}

func (o *RerunBuildForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildForbidden %s", 403, payload)
}

func (o *RerunBuildForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RerunBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

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

/*
RerunBuildNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type RerunBuildNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this rerun build not found response has a 2xx status code
func (o *RerunBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rerun build not found response has a 3xx status code
func (o *RerunBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rerun build not found response has a 4xx status code
func (o *RerunBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rerun build not found response has a 5xx status code
func (o *RerunBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rerun build not found response a status code equal to that given
func (o *RerunBuildNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rerun build not found response
func (o *RerunBuildNotFound) Code() int {
	return 404
}

func (o *RerunBuildNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildNotFound %s", 404, payload)
}

func (o *RerunBuildNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuildNotFound %s", 404, payload)
}

func (o *RerunBuildNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RerunBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

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

/*
RerunBuildDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type RerunBuildDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this rerun build default response has a 2xx status code
func (o *RerunBuildDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rerun build default response has a 3xx status code
func (o *RerunBuildDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rerun build default response has a 4xx status code
func (o *RerunBuildDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rerun build default response has a 5xx status code
func (o *RerunBuildDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rerun build default response a status code equal to that given
func (o *RerunBuildDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rerun build default response
func (o *RerunBuildDefault) Code() int {
	return o._statusCode
}

func (o *RerunBuildDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuild default %s", o._statusCode, payload)
}

func (o *RerunBuildDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] rerunBuild default %s", o._statusCode, payload)
}

func (o *RerunBuildDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *RerunBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Length
	hdrContentLength := response.GetHeader("Content-Length")

	if hdrContentLength != "" {
		valcontentLength, err := swag.ConvertUint64(hdrContentLength)
		if err != nil {
			return errors.InvalidType("Content-Length", "header", "uint64", hdrContentLength)
		}
		o.ContentLength = valcontentLength
	}

	o.Payload = new(models.ErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RerunBuildOKBody rerun build o k body
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rerunBuildOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rerun build o k body based on the context it is used
func (o *RerunBuildOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RerunBuildOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rerunBuildOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rerunBuildOK" + "." + "data")
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

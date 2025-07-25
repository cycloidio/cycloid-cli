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

/*
GetBuildOK describes a response with status code 200, with default header values.

The information of the build which has the specified id.
*/
type GetBuildOK struct {
	Payload *GetBuildOKBody
}

// IsSuccess returns true when this get build o k response has a 2xx status code
func (o *GetBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get build o k response has a 3xx status code
func (o *GetBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build o k response has a 4xx status code
func (o *GetBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get build o k response has a 5xx status code
func (o *GetBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get build o k response a status code equal to that given
func (o *GetBuildOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get build o k response
func (o *GetBuildOK) Code() int {
	return 200
}

func (o *GetBuildOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildOK %s", 200, payload)
}

func (o *GetBuildOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildOK %s", 200, payload)
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

/*
GetBuildForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetBuildForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get build forbidden response has a 2xx status code
func (o *GetBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get build forbidden response has a 3xx status code
func (o *GetBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build forbidden response has a 4xx status code
func (o *GetBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get build forbidden response has a 5xx status code
func (o *GetBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get build forbidden response a status code equal to that given
func (o *GetBuildForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get build forbidden response
func (o *GetBuildForbidden) Code() int {
	return 403
}

func (o *GetBuildForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildForbidden %s", 403, payload)
}

func (o *GetBuildForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildForbidden %s", 403, payload)
}

func (o *GetBuildForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildNotFound creates a GetBuildNotFound with default headers values
func NewGetBuildNotFound() *GetBuildNotFound {
	return &GetBuildNotFound{}
}

/*
GetBuildNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetBuildNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get build not found response has a 2xx status code
func (o *GetBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get build not found response has a 3xx status code
func (o *GetBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build not found response has a 4xx status code
func (o *GetBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get build not found response has a 5xx status code
func (o *GetBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get build not found response a status code equal to that given
func (o *GetBuildNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get build not found response
func (o *GetBuildNotFound) Code() int {
	return 404
}

func (o *GetBuildNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildNotFound %s", 404, payload)
}

func (o *GetBuildNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuildNotFound %s", 404, payload)
}

func (o *GetBuildNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBuildDefault creates a GetBuildDefault with default headers values
func NewGetBuildDefault(code int) *GetBuildDefault {
	return &GetBuildDefault{
		_statusCode: code,
	}
}

/*
GetBuildDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetBuildDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get build default response has a 2xx status code
func (o *GetBuildDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get build default response has a 3xx status code
func (o *GetBuildDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get build default response has a 4xx status code
func (o *GetBuildDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get build default response has a 5xx status code
func (o *GetBuildDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get build default response a status code equal to that given
func (o *GetBuildDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get build default response
func (o *GetBuildDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuild default %s", o._statusCode, payload)
}

func (o *GetBuildDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/jobs/{job_name}/builds/{build_id}][%d] getBuild default %s", o._statusCode, payload)
}

func (o *GetBuildDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetBuildOKBody get build o k body
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getBuildOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get build o k body based on the context it is used
func (o *GetBuildOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBuildOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getBuildOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getBuildOK" + "." + "data")
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

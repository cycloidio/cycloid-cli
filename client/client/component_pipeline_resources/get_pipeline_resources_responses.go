// Code generated by go-swagger; DO NOT EDIT.

package component_pipeline_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// GetPipelineResourcesReader is a Reader for the GetPipelineResources structure.
type GetPipelineResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPipelineResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPipelineResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPipelineResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPipelineResourcesOK creates a GetPipelineResourcesOK with default headers values
func NewGetPipelineResourcesOK() *GetPipelineResourcesOK {
	return &GetPipelineResourcesOK{}
}

/*
GetPipelineResourcesOK describes a response with status code 200, with default header values.

The resources of the pipeline's which has the specified name.
*/
type GetPipelineResourcesOK struct {
	Payload *GetPipelineResourcesOKBody
}

// IsSuccess returns true when this get pipeline resources o k response has a 2xx status code
func (o *GetPipelineResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipeline resources o k response has a 3xx status code
func (o *GetPipelineResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline resources o k response has a 4xx status code
func (o *GetPipelineResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline resources o k response has a 5xx status code
func (o *GetPipelineResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline resources o k response a status code equal to that given
func (o *GetPipelineResourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pipeline resources o k response
func (o *GetPipelineResourcesOK) Code() int {
	return 200
}

func (o *GetPipelineResourcesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResourcesOK %s", 200, payload)
}

func (o *GetPipelineResourcesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResourcesOK %s", 200, payload)
}

func (o *GetPipelineResourcesOK) GetPayload() *GetPipelineResourcesOKBody {
	return o.Payload
}

func (o *GetPipelineResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPipelineResourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineResourcesForbidden creates a GetPipelineResourcesForbidden with default headers values
func NewGetPipelineResourcesForbidden() *GetPipelineResourcesForbidden {
	return &GetPipelineResourcesForbidden{}
}

/*
GetPipelineResourcesForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetPipelineResourcesForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get pipeline resources forbidden response has a 2xx status code
func (o *GetPipelineResourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline resources forbidden response has a 3xx status code
func (o *GetPipelineResourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline resources forbidden response has a 4xx status code
func (o *GetPipelineResourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline resources forbidden response has a 5xx status code
func (o *GetPipelineResourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline resources forbidden response a status code equal to that given
func (o *GetPipelineResourcesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get pipeline resources forbidden response
func (o *GetPipelineResourcesForbidden) Code() int {
	return 403
}

func (o *GetPipelineResourcesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResourcesForbidden %s", 403, payload)
}

func (o *GetPipelineResourcesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResourcesForbidden %s", 403, payload)
}

func (o *GetPipelineResourcesForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelineResourcesNotFound creates a GetPipelineResourcesNotFound with default headers values
func NewGetPipelineResourcesNotFound() *GetPipelineResourcesNotFound {
	return &GetPipelineResourcesNotFound{}
}

/*
GetPipelineResourcesNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetPipelineResourcesNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get pipeline resources not found response has a 2xx status code
func (o *GetPipelineResourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline resources not found response has a 3xx status code
func (o *GetPipelineResourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline resources not found response has a 4xx status code
func (o *GetPipelineResourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline resources not found response has a 5xx status code
func (o *GetPipelineResourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline resources not found response a status code equal to that given
func (o *GetPipelineResourcesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get pipeline resources not found response
func (o *GetPipelineResourcesNotFound) Code() int {
	return 404
}

func (o *GetPipelineResourcesNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResourcesNotFound %s", 404, payload)
}

func (o *GetPipelineResourcesNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResourcesNotFound %s", 404, payload)
}

func (o *GetPipelineResourcesNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelineResourcesDefault creates a GetPipelineResourcesDefault with default headers values
func NewGetPipelineResourcesDefault(code int) *GetPipelineResourcesDefault {
	return &GetPipelineResourcesDefault{
		_statusCode: code,
	}
}

/*
GetPipelineResourcesDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetPipelineResourcesDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get pipeline resources default response has a 2xx status code
func (o *GetPipelineResourcesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get pipeline resources default response has a 3xx status code
func (o *GetPipelineResourcesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get pipeline resources default response has a 4xx status code
func (o *GetPipelineResourcesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get pipeline resources default response has a 5xx status code
func (o *GetPipelineResourcesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get pipeline resources default response a status code equal to that given
func (o *GetPipelineResourcesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get pipeline resources default response
func (o *GetPipelineResourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetPipelineResourcesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResources default %s", o._statusCode, payload)
}

func (o *GetPipelineResourcesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/environments/{environment_canonical}/components/{component_canonical}/pipelines/{inpath_pipeline_name}/resources][%d] getPipelineResources default %s", o._statusCode, payload)
}

func (o *GetPipelineResourcesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetPipelineResourcesOKBody get pipeline resources o k body
swagger:model GetPipelineResourcesOKBody
*/
type GetPipelineResourcesOKBody struct {

	// data
	// Required: true
	Data []*models.Resource `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get pipeline resources o k body
func (o *GetPipelineResourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPipelineResourcesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPipelineResourcesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPipelineResourcesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPipelineResourcesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPipelineResourcesOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getPipelineResourcesOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPipelineResourcesOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPipelineResourcesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get pipeline resources o k body based on the context it is used
func (o *GetPipelineResourcesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPipelineResourcesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPipelineResourcesOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPipelineResourcesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPipelineResourcesOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPipelineResourcesOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPipelineResourcesOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPipelineResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPipelineResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPipelineResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

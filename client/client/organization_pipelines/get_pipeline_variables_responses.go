// Code generated by go-swagger; DO NOT EDIT.

package organization_pipelines

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

// GetPipelineVariablesReader is a Reader for the GetPipelineVariables structure.
type GetPipelineVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPipelineVariablesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPipelineVariablesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPipelineVariablesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPipelineVariablesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPipelineVariablesOK creates a GetPipelineVariablesOK with default headers values
func NewGetPipelineVariablesOK() *GetPipelineVariablesOK {
	return &GetPipelineVariablesOK{}
}

/*GetPipelineVariablesOK handles this case with default header values.

This endpoint returns the variables of the pipeline.
*/
type GetPipelineVariablesOK struct {
	Payload *GetPipelineVariablesOKBody
}

func (o *GetPipelineVariablesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/variables][%d] getPipelineVariablesOK  %+v", 200, o.Payload)
}

func (o *GetPipelineVariablesOK) GetPayload() *GetPipelineVariablesOKBody {
	return o.Payload
}

func (o *GetPipelineVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPipelineVariablesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineVariablesForbidden creates a GetPipelineVariablesForbidden with default headers values
func NewGetPipelineVariablesForbidden() *GetPipelineVariablesForbidden {
	return &GetPipelineVariablesForbidden{}
}

/*GetPipelineVariablesForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetPipelineVariablesForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetPipelineVariablesForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/variables][%d] getPipelineVariablesForbidden  %+v", 403, o.Payload)
}

func (o *GetPipelineVariablesForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineVariablesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelineVariablesNotFound creates a GetPipelineVariablesNotFound with default headers values
func NewGetPipelineVariablesNotFound() *GetPipelineVariablesNotFound {
	return &GetPipelineVariablesNotFound{}
}

/*GetPipelineVariablesNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetPipelineVariablesNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

func (o *GetPipelineVariablesNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/variables][%d] getPipelineVariablesNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineVariablesNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineVariablesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPipelineVariablesConflict creates a GetPipelineVariablesConflict with default headers values
func NewGetPipelineVariablesConflict() *GetPipelineVariablesConflict {
	return &GetPipelineVariablesConflict{}
}

/*GetPipelineVariablesConflict handles this case with default header values.

Project has no config repository configured
*/
type GetPipelineVariablesConflict struct {
}

func (o *GetPipelineVariablesConflict) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/variables][%d] getPipelineVariablesConflict ", 409)
}

func (o *GetPipelineVariablesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPipelineVariablesDefault creates a GetPipelineVariablesDefault with default headers values
func NewGetPipelineVariablesDefault(code int) *GetPipelineVariablesDefault {
	return &GetPipelineVariablesDefault{
		_statusCode: code,
	}
}

/*GetPipelineVariablesDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetPipelineVariablesDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength int64

	Payload *models.ErrorPayload
}

// Code gets the status code for the get pipeline variables default response
func (o *GetPipelineVariablesDefault) Code() int {
	return o._statusCode
}

func (o *GetPipelineVariablesDefault) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects/{project_canonical}/pipelines/{inpath_pipeline_name}/variables][%d] getPipelineVariables default  %+v", o._statusCode, o.Payload)
}

func (o *GetPipelineVariablesDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetPipelineVariablesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetPipelineVariablesOKBody get pipeline variables o k body
swagger:model GetPipelineVariablesOKBody
*/
type GetPipelineVariablesOKBody struct {

	// data
	// Required: true
	Data *models.PipelineVariables `json:"data"`
}

// Validate validates this get pipeline variables o k body
func (o *GetPipelineVariablesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPipelineVariablesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPipelineVariablesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPipelineVariablesOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPipelineVariablesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPipelineVariablesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPipelineVariablesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

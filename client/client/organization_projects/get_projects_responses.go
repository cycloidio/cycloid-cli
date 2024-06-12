// Code generated by go-swagger; DO NOT EDIT.

package organization_projects

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

// GetProjectsReader is a Reader for the GetProjects structure.
type GetProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetProjectsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectsOK creates a GetProjectsOK with default headers values
func NewGetProjectsOK() *GetProjectsOK {
	return &GetProjectsOK{}
}

/*
GetProjectsOK describes a response with status code 200, with default header values.

List of the projects which the organization has.
*/
type GetProjectsOK struct {
	Payload *GetProjectsOKBody
}

// IsSuccess returns true when this get projects o k response has a 2xx status code
func (o *GetProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get projects o k response has a 3xx status code
func (o *GetProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects o k response has a 4xx status code
func (o *GetProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get projects o k response has a 5xx status code
func (o *GetProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects o k response a status code equal to that given
func (o *GetProjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get projects o k response
func (o *GetProjectsOK) Code() int {
	return 200
}

func (o *GetProjectsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsOK %s", 200, payload)
}

func (o *GetProjectsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsOK %s", 200, payload)
}

func (o *GetProjectsOK) GetPayload() *GetProjectsOKBody {
	return o.Payload
}

func (o *GetProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetProjectsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsForbidden creates a GetProjectsForbidden with default headers values
func NewGetProjectsForbidden() *GetProjectsForbidden {
	return &GetProjectsForbidden{}
}

/*
GetProjectsForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetProjectsForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get projects forbidden response has a 2xx status code
func (o *GetProjectsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects forbidden response has a 3xx status code
func (o *GetProjectsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects forbidden response has a 4xx status code
func (o *GetProjectsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects forbidden response has a 5xx status code
func (o *GetProjectsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects forbidden response a status code equal to that given
func (o *GetProjectsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get projects forbidden response
func (o *GetProjectsForbidden) Code() int {
	return 403
}

func (o *GetProjectsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsForbidden %s", 403, payload)
}

func (o *GetProjectsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsForbidden %s", 403, payload)
}

func (o *GetProjectsForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectsNotFound creates a GetProjectsNotFound with default headers values
func NewGetProjectsNotFound() *GetProjectsNotFound {
	return &GetProjectsNotFound{}
}

/*
GetProjectsNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetProjectsNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get projects not found response has a 2xx status code
func (o *GetProjectsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects not found response has a 3xx status code
func (o *GetProjectsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects not found response has a 4xx status code
func (o *GetProjectsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects not found response has a 5xx status code
func (o *GetProjectsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects not found response a status code equal to that given
func (o *GetProjectsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get projects not found response
func (o *GetProjectsNotFound) Code() int {
	return 404
}

func (o *GetProjectsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsNotFound %s", 404, payload)
}

func (o *GetProjectsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsNotFound %s", 404, payload)
}

func (o *GetProjectsNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectsUnprocessableEntity creates a GetProjectsUnprocessableEntity with default headers values
func NewGetProjectsUnprocessableEntity() *GetProjectsUnprocessableEntity {
	return &GetProjectsUnprocessableEntity{}
}

/*
GetProjectsUnprocessableEntity describes a response with status code 422, with default header values.

All the custom errors that are generated from the Cycloid API
*/
type GetProjectsUnprocessableEntity struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get projects unprocessable entity response has a 2xx status code
func (o *GetProjectsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects unprocessable entity response has a 3xx status code
func (o *GetProjectsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects unprocessable entity response has a 4xx status code
func (o *GetProjectsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects unprocessable entity response has a 5xx status code
func (o *GetProjectsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects unprocessable entity response a status code equal to that given
func (o *GetProjectsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get projects unprocessable entity response
func (o *GetProjectsUnprocessableEntity) Code() int {
	return 422
}

func (o *GetProjectsUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsUnprocessableEntity %s", 422, payload)
}

func (o *GetProjectsUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjectsUnprocessableEntity %s", 422, payload)
}

func (o *GetProjectsUnprocessableEntity) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetProjectsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectsDefault creates a GetProjectsDefault with default headers values
func NewGetProjectsDefault(code int) *GetProjectsDefault {
	return &GetProjectsDefault{
		_statusCode: code,
	}
}

/*
GetProjectsDefault describes a response with status code -1, with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type GetProjectsDefault struct {
	_statusCode int

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get projects default response has a 2xx status code
func (o *GetProjectsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get projects default response has a 3xx status code
func (o *GetProjectsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get projects default response has a 4xx status code
func (o *GetProjectsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get projects default response has a 5xx status code
func (o *GetProjectsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get projects default response a status code equal to that given
func (o *GetProjectsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get projects default response
func (o *GetProjectsDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjects default %s", o._statusCode, payload)
}

func (o *GetProjectsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/projects][%d] getProjects default %s", o._statusCode, payload)
}

func (o *GetProjectsDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetProjectsOKBody get projects o k body
swagger:model GetProjectsOKBody
*/
type GetProjectsOKBody struct {

	// data
	// Required: true
	Data []*models.Project `json:"data"`

	// pagination
	// Required: true
	Pagination *models.Pagination `json:"pagination"`
}

// Validate validates this get projects o k body
func (o *GetProjectsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetProjectsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getProjectsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getProjectsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getProjectsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetProjectsOKBody) validatePagination(formats strfmt.Registry) error {

	if err := validate.Required("getProjectsOK"+"."+"pagination", "body", o.Pagination); err != nil {
		return err
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProjectsOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getProjectsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get projects o k body based on the context it is used
func (o *GetProjectsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetProjectsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getProjectsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getProjectsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetProjectsOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {

		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getProjectsOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getProjectsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProjectsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProjectsOKBody) UnmarshalBinary(b []byte) error {
	var res GetProjectsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// GetSummaryReader is a Reader for the GetSummary structure.
type GetSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organization_canonical}/summary] getSummary", response, response.Code())
	}
}

// NewGetSummaryOK creates a GetSummaryOK with default headers values
func NewGetSummaryOK() *GetSummaryOK {
	return &GetSummaryOK{}
}

/*
GetSummaryOK describes a response with status code 200, with default header values.

The summary object
*/
type GetSummaryOK struct {
	Payload *GetSummaryOKBody
}

// IsSuccess returns true when this get summary o k response has a 2xx status code
func (o *GetSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get summary o k response has a 3xx status code
func (o *GetSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary o k response has a 4xx status code
func (o *GetSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get summary o k response has a 5xx status code
func (o *GetSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get summary o k response a status code equal to that given
func (o *GetSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get summary o k response
func (o *GetSummaryOK) Code() int {
	return 200
}

func (o *GetSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/summary][%d] getSummaryOK %s", 200, payload)
}

func (o *GetSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/summary][%d] getSummaryOK %s", 200, payload)
}

func (o *GetSummaryOK) GetPayload() *GetSummaryOKBody {
	return o.Payload
}

func (o *GetSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSummaryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSummaryForbidden creates a GetSummaryForbidden with default headers values
func NewGetSummaryForbidden() *GetSummaryForbidden {
	return &GetSummaryForbidden{}
}

/*
GetSummaryForbidden describes a response with status code 403, with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type GetSummaryForbidden struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get summary forbidden response has a 2xx status code
func (o *GetSummaryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get summary forbidden response has a 3xx status code
func (o *GetSummaryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary forbidden response has a 4xx status code
func (o *GetSummaryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get summary forbidden response has a 5xx status code
func (o *GetSummaryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get summary forbidden response a status code equal to that given
func (o *GetSummaryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get summary forbidden response
func (o *GetSummaryForbidden) Code() int {
	return 403
}

func (o *GetSummaryForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/summary][%d] getSummaryForbidden %s", 403, payload)
}

func (o *GetSummaryForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/summary][%d] getSummaryForbidden %s", 403, payload)
}

func (o *GetSummaryForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetSummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSummaryNotFound creates a GetSummaryNotFound with default headers values
func NewGetSummaryNotFound() *GetSummaryNotFound {
	return &GetSummaryNotFound{}
}

/*
GetSummaryNotFound describes a response with status code 404, with default header values.

The response sent when any of the entities present in the path is not found.
*/
type GetSummaryNotFound struct {

	/* The length of the response body in octets (8-bit bytes).

	   Format: uint64
	*/
	ContentLength uint64

	Payload *models.ErrorPayload
}

// IsSuccess returns true when this get summary not found response has a 2xx status code
func (o *GetSummaryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get summary not found response has a 3xx status code
func (o *GetSummaryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get summary not found response has a 4xx status code
func (o *GetSummaryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get summary not found response has a 5xx status code
func (o *GetSummaryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get summary not found response a status code equal to that given
func (o *GetSummaryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get summary not found response
func (o *GetSummaryNotFound) Code() int {
	return 404
}

func (o *GetSummaryNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/summary][%d] getSummaryNotFound %s", 404, payload)
}

func (o *GetSummaryNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{organization_canonical}/summary][%d] getSummaryNotFound %s", 404, payload)
}

func (o *GetSummaryNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *GetSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetSummaryOKBody get summary o k body
swagger:model GetSummaryOKBody
*/
type GetSummaryOKBody struct {

	// summary
	// Required: true
	Summary *models.Summary `json:"summary"`
}

// Validate validates this get summary o k body
func (o *GetSummaryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSummaryOKBody) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("getSummaryOK"+"."+"summary", "body", o.Summary); err != nil {
		return err
	}

	if o.Summary != nil {
		if err := o.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSummaryOK" + "." + "summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSummaryOK" + "." + "summary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get summary o k body based on the context it is used
func (o *GetSummaryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSummaryOKBody) contextValidateSummary(ctx context.Context, formats strfmt.Registry) error {

	if o.Summary != nil {

		if err := o.Summary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSummaryOK" + "." + "summary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getSummaryOK" + "." + "summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSummaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSummaryOKBody) UnmarshalBinary(b []byte) error {
	var res GetSummaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

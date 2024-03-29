// Code generated by go-swagger; DO NOT EDIT.

package organization_kpis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cycloidio/cycloid-cli/client/models"
)

// DeleteKPIFavoriteReader is a Reader for the DeleteKPIFavorite structure.
type DeleteKPIFavoriteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKPIFavoriteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteKPIFavoriteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteKPIFavoriteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKPIFavoriteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteKPIFavoriteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteKPIFavoriteNoContent creates a DeleteKPIFavoriteNoContent with default headers values
func NewDeleteKPIFavoriteNoContent() *DeleteKPIFavoriteNoContent {
	return &DeleteKPIFavoriteNoContent{}
}

/*DeleteKPIFavoriteNoContent handles this case with default header values.

The kpi has been removed from user favorites list.
*/
type DeleteKPIFavoriteNoContent struct {
}

func (o *DeleteKPIFavoriteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}/favorites][%d] deleteKPIFavoriteNoContent ", 204)
}

func (o *DeleteKPIFavoriteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKPIFavoriteForbidden creates a DeleteKPIFavoriteForbidden with default headers values
func NewDeleteKPIFavoriteForbidden() *DeleteKPIFavoriteForbidden {
	return &DeleteKPIFavoriteForbidden{}
}

/*DeleteKPIFavoriteForbidden handles this case with default header values.

The authenticated user cannot perform the operation because, it doesn't have permissions for such operation.
*/
type DeleteKPIFavoriteForbidden struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteKPIFavoriteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}/favorites][%d] deleteKPIFavoriteForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKPIFavoriteForbidden) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteKPIFavoriteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteKPIFavoriteNotFound creates a DeleteKPIFavoriteNotFound with default headers values
func NewDeleteKPIFavoriteNotFound() *DeleteKPIFavoriteNotFound {
	return &DeleteKPIFavoriteNotFound{}
}

/*DeleteKPIFavoriteNotFound handles this case with default header values.

The response sent when any of the entities present in the path is not found.
*/
type DeleteKPIFavoriteNotFound struct {
	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

func (o *DeleteKPIFavoriteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}/favorites][%d] deleteKPIFavoriteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKPIFavoriteNotFound) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteKPIFavoriteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteKPIFavoriteDefault creates a DeleteKPIFavoriteDefault with default headers values
func NewDeleteKPIFavoriteDefault(code int) *DeleteKPIFavoriteDefault {
	return &DeleteKPIFavoriteDefault{
		_statusCode: code,
	}
}

/*DeleteKPIFavoriteDefault handles this case with default header values.

The response sent when an unexpected error happened, as known as an internal server error.
*/
type DeleteKPIFavoriteDefault struct {
	_statusCode int

	/*The length of the response body in octets (8-bit bytes).
	 */
	ContentLength uint64

	Payload *models.ErrorPayload
}

// Code gets the status code for the delete k p i favorite default response
func (o *DeleteKPIFavoriteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteKPIFavoriteDefault) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_canonical}/kpis/{kpi_canonical}/favorites][%d] deleteKPIFavorite default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteKPIFavoriteDefault) GetPayload() *models.ErrorPayload {
	return o.Payload
}

func (o *DeleteKPIFavoriteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

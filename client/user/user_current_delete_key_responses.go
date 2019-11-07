// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCurrentDeleteKeyReader is a Reader for the UserCurrentDeleteKey structure.
type UserCurrentDeleteKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentDeleteKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserCurrentDeleteKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserCurrentDeleteKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserCurrentDeleteKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentDeleteKeyNoContent creates a UserCurrentDeleteKeyNoContent with default headers values
func NewUserCurrentDeleteKeyNoContent() *UserCurrentDeleteKeyNoContent {
	return &UserCurrentDeleteKeyNoContent{}
}

/*UserCurrentDeleteKeyNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type UserCurrentDeleteKeyNoContent struct {
}

func (o *UserCurrentDeleteKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/keys/{id}][%d] userCurrentDeleteKeyNoContent ", 204)
}

func (o *UserCurrentDeleteKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserCurrentDeleteKeyForbidden creates a UserCurrentDeleteKeyForbidden with default headers values
func NewUserCurrentDeleteKeyForbidden() *UserCurrentDeleteKeyForbidden {
	return &UserCurrentDeleteKeyForbidden{}
}

/*UserCurrentDeleteKeyForbidden handles this case with default header values.

APIForbiddenError is a forbidden error response
*/
type UserCurrentDeleteKeyForbidden struct {
	Message string

	URL string
}

func (o *UserCurrentDeleteKeyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /user/keys/{id}][%d] userCurrentDeleteKeyForbidden ", 403)
}

func (o *UserCurrentDeleteKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

// NewUserCurrentDeleteKeyNotFound creates a UserCurrentDeleteKeyNotFound with default headers values
func NewUserCurrentDeleteKeyNotFound() *UserCurrentDeleteKeyNotFound {
	return &UserCurrentDeleteKeyNotFound{}
}

/*UserCurrentDeleteKeyNotFound handles this case with default header values.

APINotFound is a not found empty response
*/
type UserCurrentDeleteKeyNotFound struct {
}

func (o *UserCurrentDeleteKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/keys/{id}][%d] userCurrentDeleteKeyNotFound ", 404)
}

func (o *UserCurrentDeleteKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
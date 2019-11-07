// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCurrentDeleteGPGKeyReader is a Reader for the UserCurrentDeleteGPGKey structure.
type UserCurrentDeleteGPGKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentDeleteGPGKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserCurrentDeleteGPGKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserCurrentDeleteGPGKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentDeleteGPGKeyNoContent creates a UserCurrentDeleteGPGKeyNoContent with default headers values
func NewUserCurrentDeleteGPGKeyNoContent() *UserCurrentDeleteGPGKeyNoContent {
	return &UserCurrentDeleteGPGKeyNoContent{}
}

/*UserCurrentDeleteGPGKeyNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type UserCurrentDeleteGPGKeyNoContent struct {
}

func (o *UserCurrentDeleteGPGKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/gpg_keys/{id}][%d] userCurrentDeleteGPGKeyNoContent ", 204)
}

func (o *UserCurrentDeleteGPGKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserCurrentDeleteGPGKeyForbidden creates a UserCurrentDeleteGPGKeyForbidden with default headers values
func NewUserCurrentDeleteGPGKeyForbidden() *UserCurrentDeleteGPGKeyForbidden {
	return &UserCurrentDeleteGPGKeyForbidden{}
}

/*UserCurrentDeleteGPGKeyForbidden handles this case with default header values.

APIForbiddenError is a forbidden error response
*/
type UserCurrentDeleteGPGKeyForbidden struct {
	Message string

	URL string
}

func (o *UserCurrentDeleteGPGKeyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /user/gpg_keys/{id}][%d] userCurrentDeleteGPGKeyForbidden ", 403)
}

func (o *UserCurrentDeleteGPGKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

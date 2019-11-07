// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "gitea.com/sapk/explore/models"
)

// UserGetReader is a Reader for the UserGet structure.
type UserGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUserGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserGetOK creates a UserGetOK with default headers values
func NewUserGetOK() *UserGetOK {
	return &UserGetOK{}
}

/*UserGetOK handles this case with default header values.

User
*/
type UserGetOK struct {
	Payload *models.User
}

func (o *UserGetOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}][%d] userGetOK  %+v", 200, o.Payload)
}

func (o *UserGetOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UserGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetNotFound creates a UserGetNotFound with default headers values
func NewUserGetNotFound() *UserGetNotFound {
	return &UserGetNotFound{}
}

/*UserGetNotFound handles this case with default header values.

APINotFound is a not found empty response
*/
type UserGetNotFound struct {
}

func (o *UserGetNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}][%d] userGetNotFound ", 404)
}

func (o *UserGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
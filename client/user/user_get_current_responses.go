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

// UserGetCurrentReader is a Reader for the UserGetCurrent structure.
type UserGetCurrentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetCurrentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetCurrentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserGetCurrentOK creates a UserGetCurrentOK with default headers values
func NewUserGetCurrentOK() *UserGetCurrentOK {
	return &UserGetCurrentOK{}
}

/*UserGetCurrentOK handles this case with default header values.

User
*/
type UserGetCurrentOK struct {
	Payload *models.User
}

func (o *UserGetCurrentOK) Error() string {
	return fmt.Sprintf("[GET /user][%d] userGetCurrentOK  %+v", 200, o.Payload)
}

func (o *UserGetCurrentOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UserGetCurrentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// UserListKeysReader is a Reader for the UserListKeys structure.
type UserListKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserListKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserListKeysOK creates a UserListKeysOK with default headers values
func NewUserListKeysOK() *UserListKeysOK {
	return &UserListKeysOK{}
}

/*UserListKeysOK handles this case with default header values.

PublicKeyList
*/
type UserListKeysOK struct {
	Payload []*models.PublicKey
}

func (o *UserListKeysOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/keys][%d] userListKeysOK  %+v", 200, o.Payload)
}

func (o *UserListKeysOK) GetPayload() []*models.PublicKey {
	return o.Payload
}

func (o *UserListKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
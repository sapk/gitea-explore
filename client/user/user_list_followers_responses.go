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

// UserListFollowersReader is a Reader for the UserListFollowers structure.
type UserListFollowersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListFollowersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserListFollowersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserListFollowersOK creates a UserListFollowersOK with default headers values
func NewUserListFollowersOK() *UserListFollowersOK {
	return &UserListFollowersOK{}
}

/*UserListFollowersOK handles this case with default header values.

UserList
*/
type UserListFollowersOK struct {
	Payload []*models.User
}

func (o *UserListFollowersOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/followers][%d] userListFollowersOK  %+v", 200, o.Payload)
}

func (o *UserListFollowersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *UserListFollowersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
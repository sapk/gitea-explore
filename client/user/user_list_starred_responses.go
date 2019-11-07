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

// UserListStarredReader is a Reader for the UserListStarred structure.
type UserListStarredReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListStarredReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserListStarredOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserListStarredOK creates a UserListStarredOK with default headers values
func NewUserListStarredOK() *UserListStarredOK {
	return &UserListStarredOK{}
}

/*UserListStarredOK handles this case with default header values.

RepositoryList
*/
type UserListStarredOK struct {
	Payload []*models.Repository
}

func (o *UserListStarredOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/starred][%d] userListStarredOK  %+v", 200, o.Payload)
}

func (o *UserListStarredOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *UserListStarredOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

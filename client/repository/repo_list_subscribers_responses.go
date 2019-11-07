// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "gitea.com/sapk/explore/models"
)

// RepoListSubscribersReader is a Reader for the RepoListSubscribers structure.
type RepoListSubscribersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListSubscribersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListSubscribersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoListSubscribersOK creates a RepoListSubscribersOK with default headers values
func NewRepoListSubscribersOK() *RepoListSubscribersOK {
	return &RepoListSubscribersOK{}
}

/*RepoListSubscribersOK handles this case with default header values.

UserList
*/
type RepoListSubscribersOK struct {
	Payload []*models.User
}

func (o *RepoListSubscribersOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/subscribers][%d] repoListSubscribersOK  %+v", 200, o.Payload)
}

func (o *RepoListSubscribersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *RepoListSubscribersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// RepoCreateHookReader is a Reader for the RepoCreateHook structure.
type RepoCreateHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoCreateHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoCreateHookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoCreateHookCreated creates a RepoCreateHookCreated with default headers values
func NewRepoCreateHookCreated() *RepoCreateHookCreated {
	return &RepoCreateHookCreated{}
}

/*RepoCreateHookCreated handles this case with default header values.

Hook
*/
type RepoCreateHookCreated struct {
	Payload *models.Hook
}

func (o *RepoCreateHookCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks][%d] repoCreateHookCreated  %+v", 201, o.Payload)
}

func (o *RepoCreateHookCreated) GetPayload() *models.Hook {
	return o.Payload
}

func (o *RepoCreateHookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

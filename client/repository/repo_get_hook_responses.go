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

// RepoGetHookReader is a Reader for the RepoGetHook structure.
type RepoGetHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoGetHookOK creates a RepoGetHookOK with default headers values
func NewRepoGetHookOK() *RepoGetHookOK {
	return &RepoGetHookOK{}
}

/*RepoGetHookOK handles this case with default header values.

Hook
*/
type RepoGetHookOK struct {
	Payload *models.Hook
}

func (o *RepoGetHookOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/{id}][%d] repoGetHookOK  %+v", 200, o.Payload)
}

func (o *RepoGetHookOK) GetPayload() *models.Hook {
	return o.Payload
}

func (o *RepoGetHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

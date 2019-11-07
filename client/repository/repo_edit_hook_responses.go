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

// RepoEditHookReader is a Reader for the RepoEditHook structure.
type RepoEditHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoEditHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoEditHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoEditHookOK creates a RepoEditHookOK with default headers values
func NewRepoEditHookOK() *RepoEditHookOK {
	return &RepoEditHookOK{}
}

/*RepoEditHookOK handles this case with default header values.

Hook
*/
type RepoEditHookOK struct {
	Payload *models.Hook
}

func (o *RepoEditHookOK) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/{id}][%d] repoEditHookOK  %+v", 200, o.Payload)
}

func (o *RepoEditHookOK) GetPayload() *models.Hook {
	return o.Payload
}

func (o *RepoEditHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
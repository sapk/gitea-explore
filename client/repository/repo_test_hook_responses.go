// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepoTestHookReader is a Reader for the RepoTestHook structure.
type RepoTestHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoTestHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoTestHookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoTestHookNoContent creates a RepoTestHookNoContent with default headers values
func NewRepoTestHookNoContent() *RepoTestHookNoContent {
	return &RepoTestHookNoContent{}
}

/*RepoTestHookNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type RepoTestHookNoContent struct {
}

func (o *RepoTestHookNoContent) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks/{id}/tests][%d] repoTestHookNoContent ", 204)
}

func (o *RepoTestHookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

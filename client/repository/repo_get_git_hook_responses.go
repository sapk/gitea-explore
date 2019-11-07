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

// RepoGetGitHookReader is a Reader for the RepoGetGitHook structure.
type RepoGetGitHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetGitHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetGitHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetGitHookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoGetGitHookOK creates a RepoGetGitHookOK with default headers values
func NewRepoGetGitHookOK() *RepoGetGitHookOK {
	return &RepoGetGitHookOK{}
}

/*RepoGetGitHookOK handles this case with default header values.

GitHook
*/
type RepoGetGitHookOK struct {
	Payload *models.GitHook
}

func (o *RepoGetGitHookOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/git/{id}][%d] repoGetGitHookOK  %+v", 200, o.Payload)
}

func (o *RepoGetGitHookOK) GetPayload() *models.GitHook {
	return o.Payload
}

func (o *RepoGetGitHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetGitHookNotFound creates a RepoGetGitHookNotFound with default headers values
func NewRepoGetGitHookNotFound() *RepoGetGitHookNotFound {
	return &RepoGetGitHookNotFound{}
}

/*RepoGetGitHookNotFound handles this case with default header values.

APINotFound is a not found empty response
*/
type RepoGetGitHookNotFound struct {
}

func (o *RepoGetGitHookNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/git/{id}][%d] repoGetGitHookNotFound ", 404)
}

func (o *RepoGetGitHookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
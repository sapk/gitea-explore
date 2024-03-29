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

// RepoListAllGitRefsReader is a Reader for the RepoListAllGitRefs structure.
type RepoListAllGitRefsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListAllGitRefsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListAllGitRefsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoListAllGitRefsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoListAllGitRefsOK creates a RepoListAllGitRefsOK with default headers values
func NewRepoListAllGitRefsOK() *RepoListAllGitRefsOK {
	return &RepoListAllGitRefsOK{}
}

/*RepoListAllGitRefsOK handles this case with default header values.

ReferenceList
*/
type RepoListAllGitRefsOK struct {
	Payload []*models.Reference
}

func (o *RepoListAllGitRefsOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs][%d] repoListAllGitRefsOK  %+v", 200, o.Payload)
}

func (o *RepoListAllGitRefsOK) GetPayload() []*models.Reference {
	return o.Payload
}

func (o *RepoListAllGitRefsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListAllGitRefsNotFound creates a RepoListAllGitRefsNotFound with default headers values
func NewRepoListAllGitRefsNotFound() *RepoListAllGitRefsNotFound {
	return &RepoListAllGitRefsNotFound{}
}

/*RepoListAllGitRefsNotFound handles this case with default header values.

APINotFound is a not found empty response
*/
type RepoListAllGitRefsNotFound struct {
}

func (o *RepoListAllGitRefsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs][%d] repoListAllGitRefsNotFound ", 404)
}

func (o *RepoListAllGitRefsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

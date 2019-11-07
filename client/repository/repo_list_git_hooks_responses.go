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

// RepoListGitHooksReader is a Reader for the RepoListGitHooks structure.
type RepoListGitHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListGitHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListGitHooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoListGitHooksOK creates a RepoListGitHooksOK with default headers values
func NewRepoListGitHooksOK() *RepoListGitHooksOK {
	return &RepoListGitHooksOK{}
}

/*RepoListGitHooksOK handles this case with default header values.

GitHookList
*/
type RepoListGitHooksOK struct {
	Payload []*models.GitHook
}

func (o *RepoListGitHooksOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/git][%d] repoListGitHooksOK  %+v", 200, o.Payload)
}

func (o *RepoListGitHooksOK) GetPayload() []*models.GitHook {
	return o.Payload
}

func (o *RepoListGitHooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

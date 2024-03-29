// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepoUpdateTopicsReader is a Reader for the RepoUpdateTopics structure.
type RepoUpdateTopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoUpdateTopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoUpdateTopicsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoUpdateTopicsNoContent creates a RepoUpdateTopicsNoContent with default headers values
func NewRepoUpdateTopicsNoContent() *RepoUpdateTopicsNoContent {
	return &RepoUpdateTopicsNoContent{}
}

/*RepoUpdateTopicsNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type RepoUpdateTopicsNoContent struct {
}

func (o *RepoUpdateTopicsNoContent) Error() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/topics][%d] repoUpdateTopicsNoContent ", 204)
}

func (o *RepoUpdateTopicsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

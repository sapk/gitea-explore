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

// RepoListTagsReader is a Reader for the RepoListTags structure.
type RepoListTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoListTagsOK creates a RepoListTagsOK with default headers values
func NewRepoListTagsOK() *RepoListTagsOK {
	return &RepoListTagsOK{}
}

/*RepoListTagsOK handles this case with default header values.

TagList
*/
type RepoListTagsOK struct {
	Payload []*models.Tag
}

func (o *RepoListTagsOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/tags][%d] repoListTagsOK  %+v", 200, o.Payload)
}

func (o *RepoListTagsOK) GetPayload() []*models.Tag {
	return o.Payload
}

func (o *RepoListTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

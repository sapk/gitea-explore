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

// RepoListStatusesReader is a Reader for the RepoListStatuses structure.
type RepoListStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoListStatusesOK creates a RepoListStatusesOK with default headers values
func NewRepoListStatusesOK() *RepoListStatusesOK {
	return &RepoListStatusesOK{}
}

/*RepoListStatusesOK handles this case with default header values.

StatusList
*/
type RepoListStatusesOK struct {
	Payload []*models.Status
}

func (o *RepoListStatusesOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/statuses/{sha}][%d] repoListStatusesOK  %+v", 200, o.Payload)
}

func (o *RepoListStatusesOK) GetPayload() []*models.Status {
	return o.Payload
}

func (o *RepoListStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
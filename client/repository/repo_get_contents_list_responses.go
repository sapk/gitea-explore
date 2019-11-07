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

// RepoGetContentsListReader is a Reader for the RepoGetContentsList structure.
type RepoGetContentsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetContentsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetContentsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoGetContentsListOK creates a RepoGetContentsListOK with default headers values
func NewRepoGetContentsListOK() *RepoGetContentsListOK {
	return &RepoGetContentsListOK{}
}

/*RepoGetContentsListOK handles this case with default header values.

ContentsListResponse
*/
type RepoGetContentsListOK struct {
	Payload []*models.ContentsResponse
}

func (o *RepoGetContentsListOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/contents][%d] repoGetContentsListOK  %+v", 200, o.Payload)
}

func (o *RepoGetContentsListOK) GetPayload() []*models.ContentsResponse {
	return o.Payload
}

func (o *RepoGetContentsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

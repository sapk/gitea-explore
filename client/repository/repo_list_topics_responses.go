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

// RepoListTopicsReader is a Reader for the RepoListTopics structure.
type RepoListTopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListTopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListTopicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoListTopicsOK creates a RepoListTopicsOK with default headers values
func NewRepoListTopicsOK() *RepoListTopicsOK {
	return &RepoListTopicsOK{}
}

/*RepoListTopicsOK handles this case with default header values.

TopicNames
*/
type RepoListTopicsOK struct {
	Payload *models.TopicName
}

func (o *RepoListTopicsOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/topics][%d] repoListTopicsOK  %+v", 200, o.Payload)
}

func (o *RepoListTopicsOK) GetPayload() *models.TopicName {
	return o.Payload
}

func (o *RepoListTopicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopicName)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

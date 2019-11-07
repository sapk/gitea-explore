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

// RepoSearchReader is a Reader for the RepoSearch structure.
type RepoSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewRepoSearchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoSearchOK creates a RepoSearchOK with default headers values
func NewRepoSearchOK() *RepoSearchOK {
	return &RepoSearchOK{}
}

/*RepoSearchOK handles this case with default header values.

SearchResults
*/
type RepoSearchOK struct {
	Payload *models.SearchResults
}

func (o *RepoSearchOK) Error() string {
	return fmt.Sprintf("[GET /repos/search][%d] repoSearchOK  %+v", 200, o.Payload)
}

func (o *RepoSearchOK) GetPayload() *models.SearchResults {
	return o.Payload
}

func (o *RepoSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoSearchUnprocessableEntity creates a RepoSearchUnprocessableEntity with default headers values
func NewRepoSearchUnprocessableEntity() *RepoSearchUnprocessableEntity {
	return &RepoSearchUnprocessableEntity{}
}

/*RepoSearchUnprocessableEntity handles this case with default header values.

APIValidationError is error format response related to input validation
*/
type RepoSearchUnprocessableEntity struct {
	Message string

	URL string
}

func (o *RepoSearchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /repos/search][%d] repoSearchUnprocessableEntity ", 422)
}

func (o *RepoSearchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

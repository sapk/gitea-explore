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

// RepoEditReader is a Reader for the RepoEdit structure.
type RepoEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoEditUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoEditOK creates a RepoEditOK with default headers values
func NewRepoEditOK() *RepoEditOK {
	return &RepoEditOK{}
}

/*RepoEditOK handles this case with default header values.

Repository
*/
type RepoEditOK struct {
	Payload *models.Repository
}

func (o *RepoEditOK) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditOK  %+v", 200, o.Payload)
}

func (o *RepoEditOK) GetPayload() *models.Repository {
	return o.Payload
}

func (o *RepoEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoEditForbidden creates a RepoEditForbidden with default headers values
func NewRepoEditForbidden() *RepoEditForbidden {
	return &RepoEditForbidden{}
}

/*RepoEditForbidden handles this case with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoEditForbidden struct {
	Message string

	URL string
}

func (o *RepoEditForbidden) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditForbidden ", 403)
}

func (o *RepoEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

// NewRepoEditUnprocessableEntity creates a RepoEditUnprocessableEntity with default headers values
func NewRepoEditUnprocessableEntity() *RepoEditUnprocessableEntity {
	return &RepoEditUnprocessableEntity{}
}

/*RepoEditUnprocessableEntity handles this case with default header values.

APIValidationError is error format response related to input validation
*/
type RepoEditUnprocessableEntity struct {
	Message string

	URL string
}

func (o *RepoEditUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}][%d] repoEditUnprocessableEntity ", 422)
}

func (o *RepoEditUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header message
	o.Message = response.GetHeader("message")

	// response header url
	o.URL = response.GetHeader("url")

	return nil
}

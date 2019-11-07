// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "gitea.com/sapk/explore/models"
)

// IssueEditIssueDeadlineReader is a Reader for the IssueEditIssueDeadline structure.
type IssueEditIssueDeadlineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueEditIssueDeadlineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIssueEditIssueDeadlineCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueEditIssueDeadlineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueEditIssueDeadlineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueEditIssueDeadlineCreated creates a IssueEditIssueDeadlineCreated with default headers values
func NewIssueEditIssueDeadlineCreated() *IssueEditIssueDeadlineCreated {
	return &IssueEditIssueDeadlineCreated{}
}

/*IssueEditIssueDeadlineCreated handles this case with default header values.

IssueDeadline
*/
type IssueEditIssueDeadlineCreated struct {
	Payload *models.IssueDeadline
}

func (o *IssueEditIssueDeadlineCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineCreated  %+v", 201, o.Payload)
}

func (o *IssueEditIssueDeadlineCreated) GetPayload() *models.IssueDeadline {
	return o.Payload
}

func (o *IssueEditIssueDeadlineCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IssueDeadline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueEditIssueDeadlineForbidden creates a IssueEditIssueDeadlineForbidden with default headers values
func NewIssueEditIssueDeadlineForbidden() *IssueEditIssueDeadlineForbidden {
	return &IssueEditIssueDeadlineForbidden{}
}

/*IssueEditIssueDeadlineForbidden handles this case with default header values.

Not repo writer
*/
type IssueEditIssueDeadlineForbidden struct {
}

func (o *IssueEditIssueDeadlineForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineForbidden ", 403)
}

func (o *IssueEditIssueDeadlineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueEditIssueDeadlineNotFound creates a IssueEditIssueDeadlineNotFound with default headers values
func NewIssueEditIssueDeadlineNotFound() *IssueEditIssueDeadlineNotFound {
	return &IssueEditIssueDeadlineNotFound{}
}

/*IssueEditIssueDeadlineNotFound handles this case with default header values.

Issue not found
*/
type IssueEditIssueDeadlineNotFound struct {
}

func (o *IssueEditIssueDeadlineNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineNotFound ", 404)
}

func (o *IssueEditIssueDeadlineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
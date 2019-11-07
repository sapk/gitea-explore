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

// IssueGetRepoCommentsReader is a Reader for the IssueGetRepoComments structure.
type IssueGetRepoCommentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueGetRepoCommentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueGetRepoCommentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueGetRepoCommentsOK creates a IssueGetRepoCommentsOK with default headers values
func NewIssueGetRepoCommentsOK() *IssueGetRepoCommentsOK {
	return &IssueGetRepoCommentsOK{}
}

/*IssueGetRepoCommentsOK handles this case with default header values.

CommentList
*/
type IssueGetRepoCommentsOK struct {
	Payload []*models.Comment
}

func (o *IssueGetRepoCommentsOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/comments][%d] issueGetRepoCommentsOK  %+v", 200, o.Payload)
}

func (o *IssueGetRepoCommentsOK) GetPayload() []*models.Comment {
	return o.Payload
}

func (o *IssueGetRepoCommentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// IssueCreateMilestoneReader is a Reader for the IssueCreateMilestone structure.
type IssueCreateMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueCreateMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIssueCreateMilestoneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueCreateMilestoneCreated creates a IssueCreateMilestoneCreated with default headers values
func NewIssueCreateMilestoneCreated() *IssueCreateMilestoneCreated {
	return &IssueCreateMilestoneCreated{}
}

/*IssueCreateMilestoneCreated handles this case with default header values.

Milestone
*/
type IssueCreateMilestoneCreated struct {
	Payload *models.Milestone
}

func (o *IssueCreateMilestoneCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/milestones][%d] issueCreateMilestoneCreated  %+v", 201, o.Payload)
}

func (o *IssueCreateMilestoneCreated) GetPayload() *models.Milestone {
	return o.Payload
}

func (o *IssueCreateMilestoneCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Milestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
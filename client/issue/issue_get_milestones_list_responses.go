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

// IssueGetMilestonesListReader is a Reader for the IssueGetMilestonesList structure.
type IssueGetMilestonesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueGetMilestonesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueGetMilestonesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueGetMilestonesListOK creates a IssueGetMilestonesListOK with default headers values
func NewIssueGetMilestonesListOK() *IssueGetMilestonesListOK {
	return &IssueGetMilestonesListOK{}
}

/*IssueGetMilestonesListOK handles this case with default header values.

MilestoneList
*/
type IssueGetMilestonesListOK struct {
	Payload []*models.Milestone
}

func (o *IssueGetMilestonesListOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/milestones][%d] issueGetMilestonesListOK  %+v", 200, o.Payload)
}

func (o *IssueGetMilestonesListOK) GetPayload() []*models.Milestone {
	return o.Payload
}

func (o *IssueGetMilestonesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

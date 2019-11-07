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

// IssueCreateLabelReader is a Reader for the IssueCreateLabel structure.
type IssueCreateLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueCreateLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIssueCreateLabelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueCreateLabelCreated creates a IssueCreateLabelCreated with default headers values
func NewIssueCreateLabelCreated() *IssueCreateLabelCreated {
	return &IssueCreateLabelCreated{}
}

/*IssueCreateLabelCreated handles this case with default header values.

Label
*/
type IssueCreateLabelCreated struct {
	Payload *models.Label
}

func (o *IssueCreateLabelCreated) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/labels][%d] issueCreateLabelCreated  %+v", 201, o.Payload)
}

func (o *IssueCreateLabelCreated) GetPayload() *models.Label {
	return o.Payload
}

func (o *IssueCreateLabelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Label)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

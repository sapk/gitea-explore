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

// IssueGetLabelReader is a Reader for the IssueGetLabel structure.
type IssueGetLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueGetLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueGetLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueGetLabelOK creates a IssueGetLabelOK with default headers values
func NewIssueGetLabelOK() *IssueGetLabelOK {
	return &IssueGetLabelOK{}
}

/*IssueGetLabelOK handles this case with default header values.

Label
*/
type IssueGetLabelOK struct {
	Payload *models.Label
}

func (o *IssueGetLabelOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/labels/{id}][%d] issueGetLabelOK  %+v", 200, o.Payload)
}

func (o *IssueGetLabelOK) GetPayload() *models.Label {
	return o.Payload
}

func (o *IssueGetLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Label)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
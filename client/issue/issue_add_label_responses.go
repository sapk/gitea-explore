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

// IssueAddLabelReader is a Reader for the IssueAddLabel structure.
type IssueAddLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueAddLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueAddLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIssueAddLabelOK creates a IssueAddLabelOK with default headers values
func NewIssueAddLabelOK() *IssueAddLabelOK {
	return &IssueAddLabelOK{}
}

/*IssueAddLabelOK handles this case with default header values.

LabelList
*/
type IssueAddLabelOK struct {
	Payload []*models.Label
}

func (o *IssueAddLabelOK) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/labels][%d] issueAddLabelOK  %+v", 200, o.Payload)
}

func (o *IssueAddLabelOK) GetPayload() []*models.Label {
	return o.Payload
}

func (o *IssueAddLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

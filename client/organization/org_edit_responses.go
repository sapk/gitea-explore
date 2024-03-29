// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "gitea.com/sapk/explore/models"
)

// OrgEditReader is a Reader for the OrgEdit structure.
type OrgEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgEditOK creates a OrgEditOK with default headers values
func NewOrgEditOK() *OrgEditOK {
	return &OrgEditOK{}
}

/*OrgEditOK handles this case with default header values.

Organization
*/
type OrgEditOK struct {
	Payload *models.Organization
}

func (o *OrgEditOK) Error() string {
	return fmt.Sprintf("[PATCH /orgs/{org}][%d] orgEditOK  %+v", 200, o.Payload)
}

func (o *OrgEditOK) GetPayload() *models.Organization {
	return o.Payload
}

func (o *OrgEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

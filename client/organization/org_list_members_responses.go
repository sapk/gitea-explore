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

// OrgListMembersReader is a Reader for the OrgListMembers structure.
type OrgListMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgListMembersOK creates a OrgListMembersOK with default headers values
func NewOrgListMembersOK() *OrgListMembersOK {
	return &OrgListMembersOK{}
}

/*OrgListMembersOK handles this case with default header values.

UserList
*/
type OrgListMembersOK struct {
	Payload []*models.User
}

func (o *OrgListMembersOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/members][%d] orgListMembersOK  %+v", 200, o.Payload)
}

func (o *OrgListMembersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *OrgListMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
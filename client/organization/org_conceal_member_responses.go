// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgConcealMemberReader is a Reader for the OrgConcealMember structure.
type OrgConcealMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgConcealMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrgConcealMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgConcealMemberNoContent creates a OrgConcealMemberNoContent with default headers values
func NewOrgConcealMemberNoContent() *OrgConcealMemberNoContent {
	return &OrgConcealMemberNoContent{}
}

/*OrgConcealMemberNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type OrgConcealMemberNoContent struct {
}

func (o *OrgConcealMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/public_members/{username}][%d] orgConcealMemberNoContent ", 204)
}

func (o *OrgConcealMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// OrgDeleteTeamReader is a Reader for the OrgDeleteTeam structure.
type OrgDeleteTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgDeleteTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOrgDeleteTeamNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgDeleteTeamNoContent creates a OrgDeleteTeamNoContent with default headers values
func NewOrgDeleteTeamNoContent() *OrgDeleteTeamNoContent {
	return &OrgDeleteTeamNoContent{}
}

/*OrgDeleteTeamNoContent handles this case with default header values.

team deleted
*/
type OrgDeleteTeamNoContent struct {
}

func (o *OrgDeleteTeamNoContent) Error() string {
	return fmt.Sprintf("[DELETE /teams/{id}][%d] orgDeleteTeamNoContent ", 204)
}

func (o *OrgDeleteTeamNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
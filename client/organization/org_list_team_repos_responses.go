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

// OrgListTeamReposReader is a Reader for the OrgListTeamRepos structure.
type OrgListTeamReposReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListTeamReposReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListTeamReposOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgListTeamReposOK creates a OrgListTeamReposOK with default headers values
func NewOrgListTeamReposOK() *OrgListTeamReposOK {
	return &OrgListTeamReposOK{}
}

/*OrgListTeamReposOK handles this case with default header values.

RepositoryList
*/
type OrgListTeamReposOK struct {
	Payload []*models.Repository
}

func (o *OrgListTeamReposOK) Error() string {
	return fmt.Sprintf("[GET /teams/{id}/repos][%d] orgListTeamReposOK  %+v", 200, o.Payload)
}

func (o *OrgListTeamReposOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *OrgListTeamReposOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

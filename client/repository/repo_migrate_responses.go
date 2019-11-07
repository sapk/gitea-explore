// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "gitea.com/sapk/explore/models"
)

// RepoMigrateReader is a Reader for the RepoMigrate structure.
type RepoMigrateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoMigrateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRepoMigrateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoMigrateCreated creates a RepoMigrateCreated with default headers values
func NewRepoMigrateCreated() *RepoMigrateCreated {
	return &RepoMigrateCreated{}
}

/*RepoMigrateCreated handles this case with default header values.

Repository
*/
type RepoMigrateCreated struct {
	Payload *models.Repository
}

func (o *RepoMigrateCreated) Error() string {
	return fmt.Sprintf("[POST /repos/migrate][%d] repoMigrateCreated  %+v", 201, o.Payload)
}

func (o *RepoMigrateCreated) GetPayload() *models.Repository {
	return o.Payload
}

func (o *RepoMigrateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

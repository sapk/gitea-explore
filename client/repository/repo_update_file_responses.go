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

// RepoUpdateFileReader is a Reader for the RepoUpdateFile structure.
type RepoUpdateFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoUpdateFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoUpdateFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoUpdateFileOK creates a RepoUpdateFileOK with default headers values
func NewRepoUpdateFileOK() *RepoUpdateFileOK {
	return &RepoUpdateFileOK{}
}

/*RepoUpdateFileOK handles this case with default header values.

FileResponse
*/
type RepoUpdateFileOK struct {
	Payload *models.FileResponse
}

func (o *RepoUpdateFileOK) Error() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/contents/{filepath}][%d] repoUpdateFileOK  %+v", 200, o.Payload)
}

func (o *RepoUpdateFileOK) GetPayload() *models.FileResponse {
	return o.Payload
}

func (o *RepoUpdateFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

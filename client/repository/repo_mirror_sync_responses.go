// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepoMirrorSyncReader is a Reader for the RepoMirrorSync structure.
type RepoMirrorSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoMirrorSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoMirrorSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoMirrorSyncOK creates a RepoMirrorSyncOK with default headers values
func NewRepoMirrorSyncOK() *RepoMirrorSyncOK {
	return &RepoMirrorSyncOK{}
}

/*RepoMirrorSyncOK handles this case with default header values.

APIEmpty is an empty response
*/
type RepoMirrorSyncOK struct {
}

func (o *RepoMirrorSyncOK) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/mirror-sync][%d] repoMirrorSyncOK ", 200)
}

func (o *RepoMirrorSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

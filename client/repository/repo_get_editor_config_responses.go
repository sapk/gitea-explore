// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepoGetEditorConfigReader is a Reader for the RepoGetEditorConfig structure.
type RepoGetEditorConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetEditorConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetEditorConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoGetEditorConfigOK creates a RepoGetEditorConfigOK with default headers values
func NewRepoGetEditorConfigOK() *RepoGetEditorConfigOK {
	return &RepoGetEditorConfigOK{}
}

/*RepoGetEditorConfigOK handles this case with default header values.

success
*/
type RepoGetEditorConfigOK struct {
}

func (o *RepoGetEditorConfigOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/editorconfig/{filepath}][%d] repoGetEditorConfigOK ", 200)
}

func (o *RepoGetEditorConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

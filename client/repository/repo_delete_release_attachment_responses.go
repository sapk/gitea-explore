// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepoDeleteReleaseAttachmentReader is a Reader for the RepoDeleteReleaseAttachment structure.
type RepoDeleteReleaseAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDeleteReleaseAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoDeleteReleaseAttachmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRepoDeleteReleaseAttachmentNoContent creates a RepoDeleteReleaseAttachmentNoContent with default headers values
func NewRepoDeleteReleaseAttachmentNoContent() *RepoDeleteReleaseAttachmentNoContent {
	return &RepoDeleteReleaseAttachmentNoContent{}
}

/*RepoDeleteReleaseAttachmentNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type RepoDeleteReleaseAttachmentNoContent struct {
}

func (o *RepoDeleteReleaseAttachmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/releases/{id}/assets/{attachment_id}][%d] repoDeleteReleaseAttachmentNoContent ", 204)
}

func (o *RepoDeleteReleaseAttachmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
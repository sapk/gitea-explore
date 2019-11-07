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

// OrgGetHookReader is a Reader for the OrgGetHook structure.
type OrgGetHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgGetHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgGetHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgGetHookOK creates a OrgGetHookOK with default headers values
func NewOrgGetHookOK() *OrgGetHookOK {
	return &OrgGetHookOK{}
}

/*OrgGetHookOK handles this case with default header values.

Hook
*/
type OrgGetHookOK struct {
	Payload *models.Hook
}

func (o *OrgGetHookOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org}/hooks/{id}][%d] orgGetHookOK  %+v", 200, o.Payload)
}

func (o *OrgGetHookOK) GetPayload() *models.Hook {
	return o.Payload
}

func (o *OrgGetHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
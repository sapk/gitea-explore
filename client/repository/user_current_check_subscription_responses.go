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

// UserCurrentCheckSubscriptionReader is a Reader for the UserCurrentCheckSubscription structure.
type UserCurrentCheckSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentCheckSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserCurrentCheckSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentCheckSubscriptionOK creates a UserCurrentCheckSubscriptionOK with default headers values
func NewUserCurrentCheckSubscriptionOK() *UserCurrentCheckSubscriptionOK {
	return &UserCurrentCheckSubscriptionOK{}
}

/*UserCurrentCheckSubscriptionOK handles this case with default header values.

WatchInfo
*/
type UserCurrentCheckSubscriptionOK struct {
	Payload *models.WatchInfo
}

func (o *UserCurrentCheckSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/subscription][%d] userCurrentCheckSubscriptionOK  %+v", 200, o.Payload)
}

func (o *UserCurrentCheckSubscriptionOK) GetPayload() *models.WatchInfo {
	return o.Payload
}

func (o *UserCurrentCheckSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WatchInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
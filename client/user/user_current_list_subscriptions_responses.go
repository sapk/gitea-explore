// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "gitea.com/sapk/explore/models"
)

// UserCurrentListSubscriptionsReader is a Reader for the UserCurrentListSubscriptions structure.
type UserCurrentListSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentListSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserCurrentListSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentListSubscriptionsOK creates a UserCurrentListSubscriptionsOK with default headers values
func NewUserCurrentListSubscriptionsOK() *UserCurrentListSubscriptionsOK {
	return &UserCurrentListSubscriptionsOK{}
}

/*UserCurrentListSubscriptionsOK handles this case with default header values.

RepositoryList
*/
type UserCurrentListSubscriptionsOK struct {
	Payload []*models.Repository
}

func (o *UserCurrentListSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /user/subscriptions][%d] userCurrentListSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *UserCurrentListSubscriptionsOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *UserCurrentListSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

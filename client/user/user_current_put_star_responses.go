// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserCurrentPutStarReader is a Reader for the UserCurrentPutStar structure.
type UserCurrentPutStarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentPutStarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserCurrentPutStarNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCurrentPutStarNoContent creates a UserCurrentPutStarNoContent with default headers values
func NewUserCurrentPutStarNoContent() *UserCurrentPutStarNoContent {
	return &UserCurrentPutStarNoContent{}
}

/*UserCurrentPutStarNoContent handles this case with default header values.

APIEmpty is an empty response
*/
type UserCurrentPutStarNoContent struct {
}

func (o *UserCurrentPutStarNoContent) Error() string {
	return fmt.Sprintf("[PUT /user/starred/{owner}/{repo}][%d] userCurrentPutStarNoContent ", 204)
}

func (o *UserCurrentPutStarNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

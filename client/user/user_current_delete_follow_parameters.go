// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserCurrentDeleteFollowParams creates a new UserCurrentDeleteFollowParams object
// with the default values initialized.
func NewUserCurrentDeleteFollowParams() *UserCurrentDeleteFollowParams {
	var ()
	return &UserCurrentDeleteFollowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentDeleteFollowParamsWithTimeout creates a new UserCurrentDeleteFollowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCurrentDeleteFollowParamsWithTimeout(timeout time.Duration) *UserCurrentDeleteFollowParams {
	var ()
	return &UserCurrentDeleteFollowParams{

		timeout: timeout,
	}
}

// NewUserCurrentDeleteFollowParamsWithContext creates a new UserCurrentDeleteFollowParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCurrentDeleteFollowParamsWithContext(ctx context.Context) *UserCurrentDeleteFollowParams {
	var ()
	return &UserCurrentDeleteFollowParams{

		Context: ctx,
	}
}

// NewUserCurrentDeleteFollowParamsWithHTTPClient creates a new UserCurrentDeleteFollowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCurrentDeleteFollowParamsWithHTTPClient(client *http.Client) *UserCurrentDeleteFollowParams {
	var ()
	return &UserCurrentDeleteFollowParams{
		HTTPClient: client,
	}
}

/*UserCurrentDeleteFollowParams contains all the parameters to send to the API endpoint
for the user current delete follow operation typically these are written to a http.Request
*/
type UserCurrentDeleteFollowParams struct {

	/*Username
	  username of user to unfollow

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) WithTimeout(timeout time.Duration) *UserCurrentDeleteFollowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) WithContext(ctx context.Context) *UserCurrentDeleteFollowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) WithHTTPClient(client *http.Client) *UserCurrentDeleteFollowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) WithUsername(username string) *UserCurrentDeleteFollowParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user current delete follow params
func (o *UserCurrentDeleteFollowParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentDeleteFollowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewUserListStarredParams creates a new UserListStarredParams object
// with the default values initialized.
func NewUserListStarredParams() *UserListStarredParams {
	var ()
	return &UserListStarredParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserListStarredParamsWithTimeout creates a new UserListStarredParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserListStarredParamsWithTimeout(timeout time.Duration) *UserListStarredParams {
	var ()
	return &UserListStarredParams{

		timeout: timeout,
	}
}

// NewUserListStarredParamsWithContext creates a new UserListStarredParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserListStarredParamsWithContext(ctx context.Context) *UserListStarredParams {
	var ()
	return &UserListStarredParams{

		Context: ctx,
	}
}

// NewUserListStarredParamsWithHTTPClient creates a new UserListStarredParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserListStarredParamsWithHTTPClient(client *http.Client) *UserListStarredParams {
	var ()
	return &UserListStarredParams{
		HTTPClient: client,
	}
}

/*UserListStarredParams contains all the parameters to send to the API endpoint
for the user list starred operation typically these are written to a http.Request
*/
type UserListStarredParams struct {

	/*Username
	  username of user

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user list starred params
func (o *UserListStarredParams) WithTimeout(timeout time.Duration) *UserListStarredParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list starred params
func (o *UserListStarredParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list starred params
func (o *UserListStarredParams) WithContext(ctx context.Context) *UserListStarredParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list starred params
func (o *UserListStarredParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list starred params
func (o *UserListStarredParams) WithHTTPClient(client *http.Client) *UserListStarredParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list starred params
func (o *UserListStarredParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the user list starred params
func (o *UserListStarredParams) WithUsername(username string) *UserListStarredParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user list starred params
func (o *UserListStarredParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserListStarredParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

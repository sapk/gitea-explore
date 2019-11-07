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

// NewUserListReposParams creates a new UserListReposParams object
// with the default values initialized.
func NewUserListReposParams() *UserListReposParams {
	var ()
	return &UserListReposParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserListReposParamsWithTimeout creates a new UserListReposParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserListReposParamsWithTimeout(timeout time.Duration) *UserListReposParams {
	var ()
	return &UserListReposParams{

		timeout: timeout,
	}
}

// NewUserListReposParamsWithContext creates a new UserListReposParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserListReposParamsWithContext(ctx context.Context) *UserListReposParams {
	var ()
	return &UserListReposParams{

		Context: ctx,
	}
}

// NewUserListReposParamsWithHTTPClient creates a new UserListReposParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserListReposParamsWithHTTPClient(client *http.Client) *UserListReposParams {
	var ()
	return &UserListReposParams{
		HTTPClient: client,
	}
}

/*UserListReposParams contains all the parameters to send to the API endpoint
for the user list repos operation typically these are written to a http.Request
*/
type UserListReposParams struct {

	/*Username
	  username of user

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user list repos params
func (o *UserListReposParams) WithTimeout(timeout time.Duration) *UserListReposParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list repos params
func (o *UserListReposParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list repos params
func (o *UserListReposParams) WithContext(ctx context.Context) *UserListReposParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list repos params
func (o *UserListReposParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list repos params
func (o *UserListReposParams) WithHTTPClient(client *http.Client) *UserListReposParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list repos params
func (o *UserListReposParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the user list repos params
func (o *UserListReposParams) WithUsername(username string) *UserListReposParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user list repos params
func (o *UserListReposParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserListReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

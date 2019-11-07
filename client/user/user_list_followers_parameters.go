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

// NewUserListFollowersParams creates a new UserListFollowersParams object
// with the default values initialized.
func NewUserListFollowersParams() *UserListFollowersParams {
	var ()
	return &UserListFollowersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserListFollowersParamsWithTimeout creates a new UserListFollowersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserListFollowersParamsWithTimeout(timeout time.Duration) *UserListFollowersParams {
	var ()
	return &UserListFollowersParams{

		timeout: timeout,
	}
}

// NewUserListFollowersParamsWithContext creates a new UserListFollowersParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserListFollowersParamsWithContext(ctx context.Context) *UserListFollowersParams {
	var ()
	return &UserListFollowersParams{

		Context: ctx,
	}
}

// NewUserListFollowersParamsWithHTTPClient creates a new UserListFollowersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserListFollowersParamsWithHTTPClient(client *http.Client) *UserListFollowersParams {
	var ()
	return &UserListFollowersParams{
		HTTPClient: client,
	}
}

/*UserListFollowersParams contains all the parameters to send to the API endpoint
for the user list followers operation typically these are written to a http.Request
*/
type UserListFollowersParams struct {

	/*Username
	  username of user

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user list followers params
func (o *UserListFollowersParams) WithTimeout(timeout time.Duration) *UserListFollowersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user list followers params
func (o *UserListFollowersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user list followers params
func (o *UserListFollowersParams) WithContext(ctx context.Context) *UserListFollowersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user list followers params
func (o *UserListFollowersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user list followers params
func (o *UserListFollowersParams) WithHTTPClient(client *http.Client) *UserListFollowersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user list followers params
func (o *UserListFollowersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the user list followers params
func (o *UserListFollowersParams) WithUsername(username string) *UserListFollowersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user list followers params
func (o *UserListFollowersParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserListFollowersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
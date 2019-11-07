// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminDeleteUserParams creates a new AdminDeleteUserParams object
// with the default values initialized.
func NewAdminDeleteUserParams() *AdminDeleteUserParams {
	var ()
	return &AdminDeleteUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeleteUserParamsWithTimeout creates a new AdminDeleteUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminDeleteUserParamsWithTimeout(timeout time.Duration) *AdminDeleteUserParams {
	var ()
	return &AdminDeleteUserParams{

		timeout: timeout,
	}
}

// NewAdminDeleteUserParamsWithContext creates a new AdminDeleteUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminDeleteUserParamsWithContext(ctx context.Context) *AdminDeleteUserParams {
	var ()
	return &AdminDeleteUserParams{

		Context: ctx,
	}
}

// NewAdminDeleteUserParamsWithHTTPClient creates a new AdminDeleteUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminDeleteUserParamsWithHTTPClient(client *http.Client) *AdminDeleteUserParams {
	var ()
	return &AdminDeleteUserParams{
		HTTPClient: client,
	}
}

/*AdminDeleteUserParams contains all the parameters to send to the API endpoint
for the admin delete user operation typically these are written to a http.Request
*/
type AdminDeleteUserParams struct {

	/*Username
	  username of user to delete

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin delete user params
func (o *AdminDeleteUserParams) WithTimeout(timeout time.Duration) *AdminDeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete user params
func (o *AdminDeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete user params
func (o *AdminDeleteUserParams) WithContext(ctx context.Context) *AdminDeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete user params
func (o *AdminDeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin delete user params
func (o *AdminDeleteUserParams) WithHTTPClient(client *http.Client) *AdminDeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete user params
func (o *AdminDeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the admin delete user params
func (o *AdminDeleteUserParams) WithUsername(username string) *AdminDeleteUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the admin delete user params
func (o *AdminDeleteUserParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
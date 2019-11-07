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

// NewUserGetTokensParams creates a new UserGetTokensParams object
// with the default values initialized.
func NewUserGetTokensParams() *UserGetTokensParams {
	var ()
	return &UserGetTokensParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetTokensParamsWithTimeout creates a new UserGetTokensParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetTokensParamsWithTimeout(timeout time.Duration) *UserGetTokensParams {
	var ()
	return &UserGetTokensParams{

		timeout: timeout,
	}
}

// NewUserGetTokensParamsWithContext creates a new UserGetTokensParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetTokensParamsWithContext(ctx context.Context) *UserGetTokensParams {
	var ()
	return &UserGetTokensParams{

		Context: ctx,
	}
}

// NewUserGetTokensParamsWithHTTPClient creates a new UserGetTokensParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetTokensParamsWithHTTPClient(client *http.Client) *UserGetTokensParams {
	var ()
	return &UserGetTokensParams{
		HTTPClient: client,
	}
}

/*UserGetTokensParams contains all the parameters to send to the API endpoint
for the user get tokens operation typically these are written to a http.Request
*/
type UserGetTokensParams struct {

	/*Username
	  username of user

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user get tokens params
func (o *UserGetTokensParams) WithTimeout(timeout time.Duration) *UserGetTokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get tokens params
func (o *UserGetTokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get tokens params
func (o *UserGetTokensParams) WithContext(ctx context.Context) *UserGetTokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get tokens params
func (o *UserGetTokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get tokens params
func (o *UserGetTokensParams) WithHTTPClient(client *http.Client) *UserGetTokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get tokens params
func (o *UserGetTokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the user get tokens params
func (o *UserGetTokensParams) WithUsername(username string) *UserGetTokensParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user get tokens params
func (o *UserGetTokensParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetTokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

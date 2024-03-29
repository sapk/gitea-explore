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

// NewUserCurrentListStarredParams creates a new UserCurrentListStarredParams object
// with the default values initialized.
func NewUserCurrentListStarredParams() *UserCurrentListStarredParams {

	return &UserCurrentListStarredParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentListStarredParamsWithTimeout creates a new UserCurrentListStarredParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCurrentListStarredParamsWithTimeout(timeout time.Duration) *UserCurrentListStarredParams {

	return &UserCurrentListStarredParams{

		timeout: timeout,
	}
}

// NewUserCurrentListStarredParamsWithContext creates a new UserCurrentListStarredParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCurrentListStarredParamsWithContext(ctx context.Context) *UserCurrentListStarredParams {

	return &UserCurrentListStarredParams{

		Context: ctx,
	}
}

// NewUserCurrentListStarredParamsWithHTTPClient creates a new UserCurrentListStarredParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCurrentListStarredParamsWithHTTPClient(client *http.Client) *UserCurrentListStarredParams {

	return &UserCurrentListStarredParams{
		HTTPClient: client,
	}
}

/*UserCurrentListStarredParams contains all the parameters to send to the API endpoint
for the user current list starred operation typically these are written to a http.Request
*/
type UserCurrentListStarredParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user current list starred params
func (o *UserCurrentListStarredParams) WithTimeout(timeout time.Duration) *UserCurrentListStarredParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current list starred params
func (o *UserCurrentListStarredParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current list starred params
func (o *UserCurrentListStarredParams) WithContext(ctx context.Context) *UserCurrentListStarredParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current list starred params
func (o *UserCurrentListStarredParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current list starred params
func (o *UserCurrentListStarredParams) WithHTTPClient(client *http.Client) *UserCurrentListStarredParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current list starred params
func (o *UserCurrentListStarredParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentListStarredParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

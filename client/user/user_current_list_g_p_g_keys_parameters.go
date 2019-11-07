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

// NewUserCurrentListGPGKeysParams creates a new UserCurrentListGPGKeysParams object
// with the default values initialized.
func NewUserCurrentListGPGKeysParams() *UserCurrentListGPGKeysParams {

	return &UserCurrentListGPGKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentListGPGKeysParamsWithTimeout creates a new UserCurrentListGPGKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCurrentListGPGKeysParamsWithTimeout(timeout time.Duration) *UserCurrentListGPGKeysParams {

	return &UserCurrentListGPGKeysParams{

		timeout: timeout,
	}
}

// NewUserCurrentListGPGKeysParamsWithContext creates a new UserCurrentListGPGKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCurrentListGPGKeysParamsWithContext(ctx context.Context) *UserCurrentListGPGKeysParams {

	return &UserCurrentListGPGKeysParams{

		Context: ctx,
	}
}

// NewUserCurrentListGPGKeysParamsWithHTTPClient creates a new UserCurrentListGPGKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCurrentListGPGKeysParamsWithHTTPClient(client *http.Client) *UserCurrentListGPGKeysParams {

	return &UserCurrentListGPGKeysParams{
		HTTPClient: client,
	}
}

/*UserCurrentListGPGKeysParams contains all the parameters to send to the API endpoint
for the user current list g p g keys operation typically these are written to a http.Request
*/
type UserCurrentListGPGKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user current list g p g keys params
func (o *UserCurrentListGPGKeysParams) WithTimeout(timeout time.Duration) *UserCurrentListGPGKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current list g p g keys params
func (o *UserCurrentListGPGKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current list g p g keys params
func (o *UserCurrentListGPGKeysParams) WithContext(ctx context.Context) *UserCurrentListGPGKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current list g p g keys params
func (o *UserCurrentListGPGKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current list g p g keys params
func (o *UserCurrentListGPGKeysParams) WithHTTPClient(client *http.Client) *UserCurrentListGPGKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current list g p g keys params
func (o *UserCurrentListGPGKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentListGPGKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewUserCurrentCheckStarringParams creates a new UserCurrentCheckStarringParams object
// with the default values initialized.
func NewUserCurrentCheckStarringParams() *UserCurrentCheckStarringParams {
	var ()
	return &UserCurrentCheckStarringParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCurrentCheckStarringParamsWithTimeout creates a new UserCurrentCheckStarringParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCurrentCheckStarringParamsWithTimeout(timeout time.Duration) *UserCurrentCheckStarringParams {
	var ()
	return &UserCurrentCheckStarringParams{

		timeout: timeout,
	}
}

// NewUserCurrentCheckStarringParamsWithContext creates a new UserCurrentCheckStarringParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCurrentCheckStarringParamsWithContext(ctx context.Context) *UserCurrentCheckStarringParams {
	var ()
	return &UserCurrentCheckStarringParams{

		Context: ctx,
	}
}

// NewUserCurrentCheckStarringParamsWithHTTPClient creates a new UserCurrentCheckStarringParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCurrentCheckStarringParamsWithHTTPClient(client *http.Client) *UserCurrentCheckStarringParams {
	var ()
	return &UserCurrentCheckStarringParams{
		HTTPClient: client,
	}
}

/*UserCurrentCheckStarringParams contains all the parameters to send to the API endpoint
for the user current check starring operation typically these are written to a http.Request
*/
type UserCurrentCheckStarringParams struct {

	/*Owner
	  owner of the repo

	*/
	Owner string
	/*Repo
	  name of the repo

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user current check starring params
func (o *UserCurrentCheckStarringParams) WithTimeout(timeout time.Duration) *UserCurrentCheckStarringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user current check starring params
func (o *UserCurrentCheckStarringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user current check starring params
func (o *UserCurrentCheckStarringParams) WithContext(ctx context.Context) *UserCurrentCheckStarringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user current check starring params
func (o *UserCurrentCheckStarringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user current check starring params
func (o *UserCurrentCheckStarringParams) WithHTTPClient(client *http.Client) *UserCurrentCheckStarringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user current check starring params
func (o *UserCurrentCheckStarringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the user current check starring params
func (o *UserCurrentCheckStarringParams) WithOwner(owner string) *UserCurrentCheckStarringParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the user current check starring params
func (o *UserCurrentCheckStarringParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the user current check starring params
func (o *UserCurrentCheckStarringParams) WithRepo(repo string) *UserCurrentCheckStarringParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the user current check starring params
func (o *UserCurrentCheckStarringParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *UserCurrentCheckStarringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

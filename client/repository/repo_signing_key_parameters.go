// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewRepoSigningKeyParams creates a new RepoSigningKeyParams object
// with the default values initialized.
func NewRepoSigningKeyParams() *RepoSigningKeyParams {
	var ()
	return &RepoSigningKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoSigningKeyParamsWithTimeout creates a new RepoSigningKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoSigningKeyParamsWithTimeout(timeout time.Duration) *RepoSigningKeyParams {
	var ()
	return &RepoSigningKeyParams{

		timeout: timeout,
	}
}

// NewRepoSigningKeyParamsWithContext creates a new RepoSigningKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoSigningKeyParamsWithContext(ctx context.Context) *RepoSigningKeyParams {
	var ()
	return &RepoSigningKeyParams{

		Context: ctx,
	}
}

// NewRepoSigningKeyParamsWithHTTPClient creates a new RepoSigningKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoSigningKeyParamsWithHTTPClient(client *http.Client) *RepoSigningKeyParams {
	var ()
	return &RepoSigningKeyParams{
		HTTPClient: client,
	}
}

/*RepoSigningKeyParams contains all the parameters to send to the API endpoint
for the repo signing key operation typically these are written to a http.Request
*/
type RepoSigningKeyParams struct {

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

// WithTimeout adds the timeout to the repo signing key params
func (o *RepoSigningKeyParams) WithTimeout(timeout time.Duration) *RepoSigningKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo signing key params
func (o *RepoSigningKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo signing key params
func (o *RepoSigningKeyParams) WithContext(ctx context.Context) *RepoSigningKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo signing key params
func (o *RepoSigningKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo signing key params
func (o *RepoSigningKeyParams) WithHTTPClient(client *http.Client) *RepoSigningKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo signing key params
func (o *RepoSigningKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo signing key params
func (o *RepoSigningKeyParams) WithOwner(owner string) *RepoSigningKeyParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo signing key params
func (o *RepoSigningKeyParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo signing key params
func (o *RepoSigningKeyParams) WithRepo(repo string) *RepoSigningKeyParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo signing key params
func (o *RepoSigningKeyParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoSigningKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
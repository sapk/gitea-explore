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

// NewRepoListTopicsParams creates a new RepoListTopicsParams object
// with the default values initialized.
func NewRepoListTopicsParams() *RepoListTopicsParams {
	var ()
	return &RepoListTopicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListTopicsParamsWithTimeout creates a new RepoListTopicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoListTopicsParamsWithTimeout(timeout time.Duration) *RepoListTopicsParams {
	var ()
	return &RepoListTopicsParams{

		timeout: timeout,
	}
}

// NewRepoListTopicsParamsWithContext creates a new RepoListTopicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoListTopicsParamsWithContext(ctx context.Context) *RepoListTopicsParams {
	var ()
	return &RepoListTopicsParams{

		Context: ctx,
	}
}

// NewRepoListTopicsParamsWithHTTPClient creates a new RepoListTopicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoListTopicsParamsWithHTTPClient(client *http.Client) *RepoListTopicsParams {
	var ()
	return &RepoListTopicsParams{
		HTTPClient: client,
	}
}

/*RepoListTopicsParams contains all the parameters to send to the API endpoint
for the repo list topics operation typically these are written to a http.Request
*/
type RepoListTopicsParams struct {

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

// WithTimeout adds the timeout to the repo list topics params
func (o *RepoListTopicsParams) WithTimeout(timeout time.Duration) *RepoListTopicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list topics params
func (o *RepoListTopicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list topics params
func (o *RepoListTopicsParams) WithContext(ctx context.Context) *RepoListTopicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list topics params
func (o *RepoListTopicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list topics params
func (o *RepoListTopicsParams) WithHTTPClient(client *http.Client) *RepoListTopicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list topics params
func (o *RepoListTopicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list topics params
func (o *RepoListTopicsParams) WithOwner(owner string) *RepoListTopicsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list topics params
func (o *RepoListTopicsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list topics params
func (o *RepoListTopicsParams) WithRepo(repo string) *RepoListTopicsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list topics params
func (o *RepoListTopicsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListTopicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

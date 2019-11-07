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

// NewRepoGetSingleCommitParams creates a new RepoGetSingleCommitParams object
// with the default values initialized.
func NewRepoGetSingleCommitParams() *RepoGetSingleCommitParams {
	var ()
	return &RepoGetSingleCommitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetSingleCommitParamsWithTimeout creates a new RepoGetSingleCommitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoGetSingleCommitParamsWithTimeout(timeout time.Duration) *RepoGetSingleCommitParams {
	var ()
	return &RepoGetSingleCommitParams{

		timeout: timeout,
	}
}

// NewRepoGetSingleCommitParamsWithContext creates a new RepoGetSingleCommitParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoGetSingleCommitParamsWithContext(ctx context.Context) *RepoGetSingleCommitParams {
	var ()
	return &RepoGetSingleCommitParams{

		Context: ctx,
	}
}

// NewRepoGetSingleCommitParamsWithHTTPClient creates a new RepoGetSingleCommitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoGetSingleCommitParamsWithHTTPClient(client *http.Client) *RepoGetSingleCommitParams {
	var ()
	return &RepoGetSingleCommitParams{
		HTTPClient: client,
	}
}

/*RepoGetSingleCommitParams contains all the parameters to send to the API endpoint
for the repo get single commit operation typically these are written to a http.Request
*/
type RepoGetSingleCommitParams struct {

	/*Owner
	  owner of the repo

	*/
	Owner string
	/*Repo
	  name of the repo

	*/
	Repo string
	/*Sha
	  the commit hash

	*/
	Sha string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithTimeout(timeout time.Duration) *RepoGetSingleCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithContext(ctx context.Context) *RepoGetSingleCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithHTTPClient(client *http.Client) *RepoGetSingleCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithOwner(owner string) *RepoGetSingleCommitParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithRepo(repo string) *RepoGetSingleCommitParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSha adds the sha to the repo get single commit params
func (o *RepoGetSingleCommitParams) WithSha(sha string) *RepoGetSingleCommitParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the repo get single commit params
func (o *RepoGetSingleCommitParams) SetSha(sha string) {
	o.Sha = sha
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetSingleCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param sha
	if err := r.SetPathParam("sha", o.Sha); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

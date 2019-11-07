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

// NewRepoListGitHooksParams creates a new RepoListGitHooksParams object
// with the default values initialized.
func NewRepoListGitHooksParams() *RepoListGitHooksParams {
	var ()
	return &RepoListGitHooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListGitHooksParamsWithTimeout creates a new RepoListGitHooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoListGitHooksParamsWithTimeout(timeout time.Duration) *RepoListGitHooksParams {
	var ()
	return &RepoListGitHooksParams{

		timeout: timeout,
	}
}

// NewRepoListGitHooksParamsWithContext creates a new RepoListGitHooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoListGitHooksParamsWithContext(ctx context.Context) *RepoListGitHooksParams {
	var ()
	return &RepoListGitHooksParams{

		Context: ctx,
	}
}

// NewRepoListGitHooksParamsWithHTTPClient creates a new RepoListGitHooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoListGitHooksParamsWithHTTPClient(client *http.Client) *RepoListGitHooksParams {
	var ()
	return &RepoListGitHooksParams{
		HTTPClient: client,
	}
}

/*RepoListGitHooksParams contains all the parameters to send to the API endpoint
for the repo list git hooks operation typically these are written to a http.Request
*/
type RepoListGitHooksParams struct {

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

// WithTimeout adds the timeout to the repo list git hooks params
func (o *RepoListGitHooksParams) WithTimeout(timeout time.Duration) *RepoListGitHooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list git hooks params
func (o *RepoListGitHooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list git hooks params
func (o *RepoListGitHooksParams) WithContext(ctx context.Context) *RepoListGitHooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list git hooks params
func (o *RepoListGitHooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list git hooks params
func (o *RepoListGitHooksParams) WithHTTPClient(client *http.Client) *RepoListGitHooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list git hooks params
func (o *RepoListGitHooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list git hooks params
func (o *RepoListGitHooksParams) WithOwner(owner string) *RepoListGitHooksParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list git hooks params
func (o *RepoListGitHooksParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list git hooks params
func (o *RepoListGitHooksParams) WithRepo(repo string) *RepoListGitHooksParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list git hooks params
func (o *RepoListGitHooksParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListGitHooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
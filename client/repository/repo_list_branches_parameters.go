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

// NewRepoListBranchesParams creates a new RepoListBranchesParams object
// with the default values initialized.
func NewRepoListBranchesParams() *RepoListBranchesParams {
	var ()
	return &RepoListBranchesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListBranchesParamsWithTimeout creates a new RepoListBranchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoListBranchesParamsWithTimeout(timeout time.Duration) *RepoListBranchesParams {
	var ()
	return &RepoListBranchesParams{

		timeout: timeout,
	}
}

// NewRepoListBranchesParamsWithContext creates a new RepoListBranchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoListBranchesParamsWithContext(ctx context.Context) *RepoListBranchesParams {
	var ()
	return &RepoListBranchesParams{

		Context: ctx,
	}
}

// NewRepoListBranchesParamsWithHTTPClient creates a new RepoListBranchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoListBranchesParamsWithHTTPClient(client *http.Client) *RepoListBranchesParams {
	var ()
	return &RepoListBranchesParams{
		HTTPClient: client,
	}
}

/*RepoListBranchesParams contains all the parameters to send to the API endpoint
for the repo list branches operation typically these are written to a http.Request
*/
type RepoListBranchesParams struct {

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

// WithTimeout adds the timeout to the repo list branches params
func (o *RepoListBranchesParams) WithTimeout(timeout time.Duration) *RepoListBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list branches params
func (o *RepoListBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list branches params
func (o *RepoListBranchesParams) WithContext(ctx context.Context) *RepoListBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list branches params
func (o *RepoListBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list branches params
func (o *RepoListBranchesParams) WithHTTPClient(client *http.Client) *RepoListBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list branches params
func (o *RepoListBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list branches params
func (o *RepoListBranchesParams) WithOwner(owner string) *RepoListBranchesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list branches params
func (o *RepoListBranchesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list branches params
func (o *RepoListBranchesParams) WithRepo(repo string) *RepoListBranchesParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list branches params
func (o *RepoListBranchesParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

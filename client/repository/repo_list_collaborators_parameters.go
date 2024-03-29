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

// NewRepoListCollaboratorsParams creates a new RepoListCollaboratorsParams object
// with the default values initialized.
func NewRepoListCollaboratorsParams() *RepoListCollaboratorsParams {
	var ()
	return &RepoListCollaboratorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoListCollaboratorsParamsWithTimeout creates a new RepoListCollaboratorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoListCollaboratorsParamsWithTimeout(timeout time.Duration) *RepoListCollaboratorsParams {
	var ()
	return &RepoListCollaboratorsParams{

		timeout: timeout,
	}
}

// NewRepoListCollaboratorsParamsWithContext creates a new RepoListCollaboratorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoListCollaboratorsParamsWithContext(ctx context.Context) *RepoListCollaboratorsParams {
	var ()
	return &RepoListCollaboratorsParams{

		Context: ctx,
	}
}

// NewRepoListCollaboratorsParamsWithHTTPClient creates a new RepoListCollaboratorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoListCollaboratorsParamsWithHTTPClient(client *http.Client) *RepoListCollaboratorsParams {
	var ()
	return &RepoListCollaboratorsParams{
		HTTPClient: client,
	}
}

/*RepoListCollaboratorsParams contains all the parameters to send to the API endpoint
for the repo list collaborators operation typically these are written to a http.Request
*/
type RepoListCollaboratorsParams struct {

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

// WithTimeout adds the timeout to the repo list collaborators params
func (o *RepoListCollaboratorsParams) WithTimeout(timeout time.Duration) *RepoListCollaboratorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo list collaborators params
func (o *RepoListCollaboratorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo list collaborators params
func (o *RepoListCollaboratorsParams) WithContext(ctx context.Context) *RepoListCollaboratorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo list collaborators params
func (o *RepoListCollaboratorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo list collaborators params
func (o *RepoListCollaboratorsParams) WithHTTPClient(client *http.Client) *RepoListCollaboratorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo list collaborators params
func (o *RepoListCollaboratorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo list collaborators params
func (o *RepoListCollaboratorsParams) WithOwner(owner string) *RepoListCollaboratorsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo list collaborators params
func (o *RepoListCollaboratorsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo list collaborators params
func (o *RepoListCollaboratorsParams) WithRepo(repo string) *RepoListCollaboratorsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo list collaborators params
func (o *RepoListCollaboratorsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoListCollaboratorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

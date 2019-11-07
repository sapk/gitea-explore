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

// NewRepoCheckCollaboratorParams creates a new RepoCheckCollaboratorParams object
// with the default values initialized.
func NewRepoCheckCollaboratorParams() *RepoCheckCollaboratorParams {
	var ()
	return &RepoCheckCollaboratorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCheckCollaboratorParamsWithTimeout creates a new RepoCheckCollaboratorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoCheckCollaboratorParamsWithTimeout(timeout time.Duration) *RepoCheckCollaboratorParams {
	var ()
	return &RepoCheckCollaboratorParams{

		timeout: timeout,
	}
}

// NewRepoCheckCollaboratorParamsWithContext creates a new RepoCheckCollaboratorParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoCheckCollaboratorParamsWithContext(ctx context.Context) *RepoCheckCollaboratorParams {
	var ()
	return &RepoCheckCollaboratorParams{

		Context: ctx,
	}
}

// NewRepoCheckCollaboratorParamsWithHTTPClient creates a new RepoCheckCollaboratorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoCheckCollaboratorParamsWithHTTPClient(client *http.Client) *RepoCheckCollaboratorParams {
	var ()
	return &RepoCheckCollaboratorParams{
		HTTPClient: client,
	}
}

/*RepoCheckCollaboratorParams contains all the parameters to send to the API endpoint
for the repo check collaborator operation typically these are written to a http.Request
*/
type RepoCheckCollaboratorParams struct {

	/*Collaborator
	  username of the collaborator

	*/
	Collaborator string
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

// WithTimeout adds the timeout to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) WithTimeout(timeout time.Duration) *RepoCheckCollaboratorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) WithContext(ctx context.Context) *RepoCheckCollaboratorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) WithHTTPClient(client *http.Client) *RepoCheckCollaboratorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollaborator adds the collaborator to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) WithCollaborator(collaborator string) *RepoCheckCollaboratorParams {
	o.SetCollaborator(collaborator)
	return o
}

// SetCollaborator adds the collaborator to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) SetCollaborator(collaborator string) {
	o.Collaborator = collaborator
}

// WithOwner adds the owner to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) WithOwner(owner string) *RepoCheckCollaboratorParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) WithRepo(repo string) *RepoCheckCollaboratorParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo check collaborator params
func (o *RepoCheckCollaboratorParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCheckCollaboratorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collaborator
	if err := r.SetPathParam("collaborator", o.Collaborator); err != nil {
		return err
	}

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
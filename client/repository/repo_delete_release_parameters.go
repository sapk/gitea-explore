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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRepoDeleteReleaseParams creates a new RepoDeleteReleaseParams object
// with the default values initialized.
func NewRepoDeleteReleaseParams() *RepoDeleteReleaseParams {
	var ()
	return &RepoDeleteReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteReleaseParamsWithTimeout creates a new RepoDeleteReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoDeleteReleaseParamsWithTimeout(timeout time.Duration) *RepoDeleteReleaseParams {
	var ()
	return &RepoDeleteReleaseParams{

		timeout: timeout,
	}
}

// NewRepoDeleteReleaseParamsWithContext creates a new RepoDeleteReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoDeleteReleaseParamsWithContext(ctx context.Context) *RepoDeleteReleaseParams {
	var ()
	return &RepoDeleteReleaseParams{

		Context: ctx,
	}
}

// NewRepoDeleteReleaseParamsWithHTTPClient creates a new RepoDeleteReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoDeleteReleaseParamsWithHTTPClient(client *http.Client) *RepoDeleteReleaseParams {
	var ()
	return &RepoDeleteReleaseParams{
		HTTPClient: client,
	}
}

/*RepoDeleteReleaseParams contains all the parameters to send to the API endpoint
for the repo delete release operation typically these are written to a http.Request
*/
type RepoDeleteReleaseParams struct {

	/*ID
	  id of the release to delete

	*/
	ID int64
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

// WithTimeout adds the timeout to the repo delete release params
func (o *RepoDeleteReleaseParams) WithTimeout(timeout time.Duration) *RepoDeleteReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete release params
func (o *RepoDeleteReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete release params
func (o *RepoDeleteReleaseParams) WithContext(ctx context.Context) *RepoDeleteReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete release params
func (o *RepoDeleteReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete release params
func (o *RepoDeleteReleaseParams) WithHTTPClient(client *http.Client) *RepoDeleteReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete release params
func (o *RepoDeleteReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the repo delete release params
func (o *RepoDeleteReleaseParams) WithID(id int64) *RepoDeleteReleaseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repo delete release params
func (o *RepoDeleteReleaseParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the repo delete release params
func (o *RepoDeleteReleaseParams) WithOwner(owner string) *RepoDeleteReleaseParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete release params
func (o *RepoDeleteReleaseParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete release params
func (o *RepoDeleteReleaseParams) WithRepo(repo string) *RepoDeleteReleaseParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete release params
func (o *RepoDeleteReleaseParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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

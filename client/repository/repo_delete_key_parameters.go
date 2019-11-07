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

// NewRepoDeleteKeyParams creates a new RepoDeleteKeyParams object
// with the default values initialized.
func NewRepoDeleteKeyParams() *RepoDeleteKeyParams {
	var ()
	return &RepoDeleteKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteKeyParamsWithTimeout creates a new RepoDeleteKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoDeleteKeyParamsWithTimeout(timeout time.Duration) *RepoDeleteKeyParams {
	var ()
	return &RepoDeleteKeyParams{

		timeout: timeout,
	}
}

// NewRepoDeleteKeyParamsWithContext creates a new RepoDeleteKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoDeleteKeyParamsWithContext(ctx context.Context) *RepoDeleteKeyParams {
	var ()
	return &RepoDeleteKeyParams{

		Context: ctx,
	}
}

// NewRepoDeleteKeyParamsWithHTTPClient creates a new RepoDeleteKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoDeleteKeyParamsWithHTTPClient(client *http.Client) *RepoDeleteKeyParams {
	var ()
	return &RepoDeleteKeyParams{
		HTTPClient: client,
	}
}

/*RepoDeleteKeyParams contains all the parameters to send to the API endpoint
for the repo delete key operation typically these are written to a http.Request
*/
type RepoDeleteKeyParams struct {

	/*ID
	  id of the key to delete

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

// WithTimeout adds the timeout to the repo delete key params
func (o *RepoDeleteKeyParams) WithTimeout(timeout time.Duration) *RepoDeleteKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete key params
func (o *RepoDeleteKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete key params
func (o *RepoDeleteKeyParams) WithContext(ctx context.Context) *RepoDeleteKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete key params
func (o *RepoDeleteKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete key params
func (o *RepoDeleteKeyParams) WithHTTPClient(client *http.Client) *RepoDeleteKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete key params
func (o *RepoDeleteKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the repo delete key params
func (o *RepoDeleteKeyParams) WithID(id int64) *RepoDeleteKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repo delete key params
func (o *RepoDeleteKeyParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the repo delete key params
func (o *RepoDeleteKeyParams) WithOwner(owner string) *RepoDeleteKeyParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete key params
func (o *RepoDeleteKeyParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete key params
func (o *RepoDeleteKeyParams) WithRepo(repo string) *RepoDeleteKeyParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete key params
func (o *RepoDeleteKeyParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

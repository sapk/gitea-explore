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

// NewRepoGetReleaseParams creates a new RepoGetReleaseParams object
// with the default values initialized.
func NewRepoGetReleaseParams() *RepoGetReleaseParams {
	var ()
	return &RepoGetReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetReleaseParamsWithTimeout creates a new RepoGetReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoGetReleaseParamsWithTimeout(timeout time.Duration) *RepoGetReleaseParams {
	var ()
	return &RepoGetReleaseParams{

		timeout: timeout,
	}
}

// NewRepoGetReleaseParamsWithContext creates a new RepoGetReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoGetReleaseParamsWithContext(ctx context.Context) *RepoGetReleaseParams {
	var ()
	return &RepoGetReleaseParams{

		Context: ctx,
	}
}

// NewRepoGetReleaseParamsWithHTTPClient creates a new RepoGetReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoGetReleaseParamsWithHTTPClient(client *http.Client) *RepoGetReleaseParams {
	var ()
	return &RepoGetReleaseParams{
		HTTPClient: client,
	}
}

/*RepoGetReleaseParams contains all the parameters to send to the API endpoint
for the repo get release operation typically these are written to a http.Request
*/
type RepoGetReleaseParams struct {

	/*ID
	  id of the release to get

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

// WithTimeout adds the timeout to the repo get release params
func (o *RepoGetReleaseParams) WithTimeout(timeout time.Duration) *RepoGetReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get release params
func (o *RepoGetReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get release params
func (o *RepoGetReleaseParams) WithContext(ctx context.Context) *RepoGetReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get release params
func (o *RepoGetReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get release params
func (o *RepoGetReleaseParams) WithHTTPClient(client *http.Client) *RepoGetReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get release params
func (o *RepoGetReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the repo get release params
func (o *RepoGetReleaseParams) WithID(id int64) *RepoGetReleaseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repo get release params
func (o *RepoGetReleaseParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the repo get release params
func (o *RepoGetReleaseParams) WithOwner(owner string) *RepoGetReleaseParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get release params
func (o *RepoGetReleaseParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo get release params
func (o *RepoGetReleaseParams) WithRepo(repo string) *RepoGetReleaseParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get release params
func (o *RepoGetReleaseParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

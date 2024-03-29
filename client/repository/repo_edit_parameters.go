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

	models "gitea.com/sapk/explore/models"
)

// NewRepoEditParams creates a new RepoEditParams object
// with the default values initialized.
func NewRepoEditParams() *RepoEditParams {
	var ()
	return &RepoEditParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoEditParamsWithTimeout creates a new RepoEditParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoEditParamsWithTimeout(timeout time.Duration) *RepoEditParams {
	var ()
	return &RepoEditParams{

		timeout: timeout,
	}
}

// NewRepoEditParamsWithContext creates a new RepoEditParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoEditParamsWithContext(ctx context.Context) *RepoEditParams {
	var ()
	return &RepoEditParams{

		Context: ctx,
	}
}

// NewRepoEditParamsWithHTTPClient creates a new RepoEditParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoEditParamsWithHTTPClient(client *http.Client) *RepoEditParams {
	var ()
	return &RepoEditParams{
		HTTPClient: client,
	}
}

/*RepoEditParams contains all the parameters to send to the API endpoint
for the repo edit operation typically these are written to a http.Request
*/
type RepoEditParams struct {

	/*Body
	  Properties of a repo that you can edit

	*/
	Body *models.EditRepoOption
	/*Owner
	  owner of the repo to edit

	*/
	Owner string
	/*Repo
	  name of the repo to edit

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repo edit params
func (o *RepoEditParams) WithTimeout(timeout time.Duration) *RepoEditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo edit params
func (o *RepoEditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo edit params
func (o *RepoEditParams) WithContext(ctx context.Context) *RepoEditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo edit params
func (o *RepoEditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo edit params
func (o *RepoEditParams) WithHTTPClient(client *http.Client) *RepoEditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo edit params
func (o *RepoEditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo edit params
func (o *RepoEditParams) WithBody(body *models.EditRepoOption) *RepoEditParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo edit params
func (o *RepoEditParams) SetBody(body *models.EditRepoOption) {
	o.Body = body
}

// WithOwner adds the owner to the repo edit params
func (o *RepoEditParams) WithOwner(owner string) *RepoEditParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo edit params
func (o *RepoEditParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo edit params
func (o *RepoEditParams) WithRepo(repo string) *RepoEditParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo edit params
func (o *RepoEditParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoEditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

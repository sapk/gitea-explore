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

// NewRepoCreateFileParams creates a new RepoCreateFileParams object
// with the default values initialized.
func NewRepoCreateFileParams() *RepoCreateFileParams {
	var ()
	return &RepoCreateFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoCreateFileParamsWithTimeout creates a new RepoCreateFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoCreateFileParamsWithTimeout(timeout time.Duration) *RepoCreateFileParams {
	var ()
	return &RepoCreateFileParams{

		timeout: timeout,
	}
}

// NewRepoCreateFileParamsWithContext creates a new RepoCreateFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoCreateFileParamsWithContext(ctx context.Context) *RepoCreateFileParams {
	var ()
	return &RepoCreateFileParams{

		Context: ctx,
	}
}

// NewRepoCreateFileParamsWithHTTPClient creates a new RepoCreateFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoCreateFileParamsWithHTTPClient(client *http.Client) *RepoCreateFileParams {
	var ()
	return &RepoCreateFileParams{
		HTTPClient: client,
	}
}

/*RepoCreateFileParams contains all the parameters to send to the API endpoint
for the repo create file operation typically these are written to a http.Request
*/
type RepoCreateFileParams struct {

	/*Body*/
	Body *models.CreateFileOptions
	/*Filepath
	  path of the file to create

	*/
	Filepath string
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

// WithTimeout adds the timeout to the repo create file params
func (o *RepoCreateFileParams) WithTimeout(timeout time.Duration) *RepoCreateFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo create file params
func (o *RepoCreateFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo create file params
func (o *RepoCreateFileParams) WithContext(ctx context.Context) *RepoCreateFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo create file params
func (o *RepoCreateFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo create file params
func (o *RepoCreateFileParams) WithHTTPClient(client *http.Client) *RepoCreateFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo create file params
func (o *RepoCreateFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo create file params
func (o *RepoCreateFileParams) WithBody(body *models.CreateFileOptions) *RepoCreateFileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo create file params
func (o *RepoCreateFileParams) SetBody(body *models.CreateFileOptions) {
	o.Body = body
}

// WithFilepath adds the filepath to the repo create file params
func (o *RepoCreateFileParams) WithFilepath(filepath string) *RepoCreateFileParams {
	o.SetFilepath(filepath)
	return o
}

// SetFilepath adds the filepath to the repo create file params
func (o *RepoCreateFileParams) SetFilepath(filepath string) {
	o.Filepath = filepath
}

// WithOwner adds the owner to the repo create file params
func (o *RepoCreateFileParams) WithOwner(owner string) *RepoCreateFileParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo create file params
func (o *RepoCreateFileParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo create file params
func (o *RepoCreateFileParams) WithRepo(repo string) *RepoCreateFileParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo create file params
func (o *RepoCreateFileParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoCreateFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param filepath
	if err := r.SetPathParam("filepath", o.Filepath); err != nil {
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
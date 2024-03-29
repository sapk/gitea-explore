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

// NewRepoDeleteFileParams creates a new RepoDeleteFileParams object
// with the default values initialized.
func NewRepoDeleteFileParams() *RepoDeleteFileParams {
	var ()
	return &RepoDeleteFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoDeleteFileParamsWithTimeout creates a new RepoDeleteFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoDeleteFileParamsWithTimeout(timeout time.Duration) *RepoDeleteFileParams {
	var ()
	return &RepoDeleteFileParams{

		timeout: timeout,
	}
}

// NewRepoDeleteFileParamsWithContext creates a new RepoDeleteFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoDeleteFileParamsWithContext(ctx context.Context) *RepoDeleteFileParams {
	var ()
	return &RepoDeleteFileParams{

		Context: ctx,
	}
}

// NewRepoDeleteFileParamsWithHTTPClient creates a new RepoDeleteFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoDeleteFileParamsWithHTTPClient(client *http.Client) *RepoDeleteFileParams {
	var ()
	return &RepoDeleteFileParams{
		HTTPClient: client,
	}
}

/*RepoDeleteFileParams contains all the parameters to send to the API endpoint
for the repo delete file operation typically these are written to a http.Request
*/
type RepoDeleteFileParams struct {

	/*Body*/
	Body *models.DeleteFileOptions
	/*Filepath
	  path of the file to delete

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

// WithTimeout adds the timeout to the repo delete file params
func (o *RepoDeleteFileParams) WithTimeout(timeout time.Duration) *RepoDeleteFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo delete file params
func (o *RepoDeleteFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo delete file params
func (o *RepoDeleteFileParams) WithContext(ctx context.Context) *RepoDeleteFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo delete file params
func (o *RepoDeleteFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo delete file params
func (o *RepoDeleteFileParams) WithHTTPClient(client *http.Client) *RepoDeleteFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo delete file params
func (o *RepoDeleteFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repo delete file params
func (o *RepoDeleteFileParams) WithBody(body *models.DeleteFileOptions) *RepoDeleteFileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repo delete file params
func (o *RepoDeleteFileParams) SetBody(body *models.DeleteFileOptions) {
	o.Body = body
}

// WithFilepath adds the filepath to the repo delete file params
func (o *RepoDeleteFileParams) WithFilepath(filepath string) *RepoDeleteFileParams {
	o.SetFilepath(filepath)
	return o
}

// SetFilepath adds the filepath to the repo delete file params
func (o *RepoDeleteFileParams) SetFilepath(filepath string) {
	o.Filepath = filepath
}

// WithOwner adds the owner to the repo delete file params
func (o *RepoDeleteFileParams) WithOwner(owner string) *RepoDeleteFileParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo delete file params
func (o *RepoDeleteFileParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the repo delete file params
func (o *RepoDeleteFileParams) WithRepo(repo string) *RepoDeleteFileParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo delete file params
func (o *RepoDeleteFileParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoDeleteFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

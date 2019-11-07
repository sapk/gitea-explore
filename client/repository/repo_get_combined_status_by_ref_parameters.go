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

// NewRepoGetCombinedStatusByRefParams creates a new RepoGetCombinedStatusByRefParams object
// with the default values initialized.
func NewRepoGetCombinedStatusByRefParams() *RepoGetCombinedStatusByRefParams {
	var ()
	return &RepoGetCombinedStatusByRefParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoGetCombinedStatusByRefParamsWithTimeout creates a new RepoGetCombinedStatusByRefParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoGetCombinedStatusByRefParamsWithTimeout(timeout time.Duration) *RepoGetCombinedStatusByRefParams {
	var ()
	return &RepoGetCombinedStatusByRefParams{

		timeout: timeout,
	}
}

// NewRepoGetCombinedStatusByRefParamsWithContext creates a new RepoGetCombinedStatusByRefParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoGetCombinedStatusByRefParamsWithContext(ctx context.Context) *RepoGetCombinedStatusByRefParams {
	var ()
	return &RepoGetCombinedStatusByRefParams{

		Context: ctx,
	}
}

// NewRepoGetCombinedStatusByRefParamsWithHTTPClient creates a new RepoGetCombinedStatusByRefParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoGetCombinedStatusByRefParamsWithHTTPClient(client *http.Client) *RepoGetCombinedStatusByRefParams {
	var ()
	return &RepoGetCombinedStatusByRefParams{
		HTTPClient: client,
	}
}

/*RepoGetCombinedStatusByRefParams contains all the parameters to send to the API endpoint
for the repo get combined status by ref operation typically these are written to a http.Request
*/
type RepoGetCombinedStatusByRefParams struct {

	/*Owner
	  owner of the repo

	*/
	Owner string
	/*Page
	  page number of results

	*/
	Page *int64
	/*Ref
	  name of branch/tag/commit

	*/
	Ref string
	/*Repo
	  name of the repo

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) WithTimeout(timeout time.Duration) *RepoGetCombinedStatusByRefParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) WithContext(ctx context.Context) *RepoGetCombinedStatusByRefParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) WithHTTPClient(client *http.Client) *RepoGetCombinedStatusByRefParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) WithOwner(owner string) *RepoGetCombinedStatusByRefParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) WithPage(page *int64) *RepoGetCombinedStatusByRefParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) SetPage(page *int64) {
	o.Page = page
}

// WithRef adds the ref to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) WithRef(ref string) *RepoGetCombinedStatusByRefParams {
	o.SetRef(ref)
	return o
}

// SetRef adds the ref to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) SetRef(ref string) {
	o.Ref = ref
}

// WithRepo adds the repo to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) WithRepo(repo string) *RepoGetCombinedStatusByRefParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the repo get combined status by ref params
func (o *RepoGetCombinedStatusByRefParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *RepoGetCombinedStatusByRefParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	// path param ref
	if err := r.SetPathParam("ref", o.Ref); err != nil {
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
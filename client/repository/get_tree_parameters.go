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

// NewGetTreeParams creates a new GetTreeParams object
// with the default values initialized.
func NewGetTreeParams() *GetTreeParams {
	var ()
	return &GetTreeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTreeParamsWithTimeout creates a new GetTreeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTreeParamsWithTimeout(timeout time.Duration) *GetTreeParams {
	var ()
	return &GetTreeParams{

		timeout: timeout,
	}
}

// NewGetTreeParamsWithContext creates a new GetTreeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTreeParamsWithContext(ctx context.Context) *GetTreeParams {
	var ()
	return &GetTreeParams{

		Context: ctx,
	}
}

// NewGetTreeParamsWithHTTPClient creates a new GetTreeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTreeParamsWithHTTPClient(client *http.Client) *GetTreeParams {
	var ()
	return &GetTreeParams{
		HTTPClient: client,
	}
}

/*GetTreeParams contains all the parameters to send to the API endpoint
for the get tree operation typically these are written to a http.Request
*/
type GetTreeParams struct {

	/*Owner
	  owner of the repo

	*/
	Owner string
	/*Page
	  page number; the 'truncated' field in the response will be true if there are still more items after this page, false if the last page

	*/
	Page *int64
	/*PerPage
	  number of items per page; default is 1000 or what is set in app.ini as DEFAULT_GIT_TREES_PER_PAGE

	*/
	PerPage *int64
	/*Recursive
	  show all directories and files

	*/
	Recursive *bool
	/*Repo
	  name of the repo

	*/
	Repo string
	/*Sha
	  sha of the commit

	*/
	Sha string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tree params
func (o *GetTreeParams) WithTimeout(timeout time.Duration) *GetTreeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tree params
func (o *GetTreeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tree params
func (o *GetTreeParams) WithContext(ctx context.Context) *GetTreeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tree params
func (o *GetTreeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tree params
func (o *GetTreeParams) WithHTTPClient(client *http.Client) *GetTreeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tree params
func (o *GetTreeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get tree params
func (o *GetTreeParams) WithOwner(owner string) *GetTreeParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get tree params
func (o *GetTreeParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPage adds the page to the get tree params
func (o *GetTreeParams) WithPage(page *int64) *GetTreeParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get tree params
func (o *GetTreeParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get tree params
func (o *GetTreeParams) WithPerPage(perPage *int64) *GetTreeParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get tree params
func (o *GetTreeParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithRecursive adds the recursive to the get tree params
func (o *GetTreeParams) WithRecursive(recursive *bool) *GetTreeParams {
	o.SetRecursive(recursive)
	return o
}

// SetRecursive adds the recursive to the get tree params
func (o *GetTreeParams) SetRecursive(recursive *bool) {
	o.Recursive = recursive
}

// WithRepo adds the repo to the get tree params
func (o *GetTreeParams) WithRepo(repo string) *GetTreeParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the get tree params
func (o *GetTreeParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithSha adds the sha to the get tree params
func (o *GetTreeParams) WithSha(sha string) *GetTreeParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the get tree params
func (o *GetTreeParams) SetSha(sha string) {
	o.Sha = sha
}

// WriteToRequest writes these params to a swagger request
func (o *GetTreeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive bool
		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := swag.FormatBool(qrRecursive)
		if qRecursive != "" {
			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
				return err
			}
		}

	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	// path param sha
	if err := r.SetPathParam("sha", o.Sha); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
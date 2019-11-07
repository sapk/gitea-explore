// Code generated by go-swagger; DO NOT EDIT.

package issue

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

	models "gitea.com/sapk/explore/models"
)

// NewIssueReplaceLabelsParams creates a new IssueReplaceLabelsParams object
// with the default values initialized.
func NewIssueReplaceLabelsParams() *IssueReplaceLabelsParams {
	var ()
	return &IssueReplaceLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIssueReplaceLabelsParamsWithTimeout creates a new IssueReplaceLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIssueReplaceLabelsParamsWithTimeout(timeout time.Duration) *IssueReplaceLabelsParams {
	var ()
	return &IssueReplaceLabelsParams{

		timeout: timeout,
	}
}

// NewIssueReplaceLabelsParamsWithContext creates a new IssueReplaceLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIssueReplaceLabelsParamsWithContext(ctx context.Context) *IssueReplaceLabelsParams {
	var ()
	return &IssueReplaceLabelsParams{

		Context: ctx,
	}
}

// NewIssueReplaceLabelsParamsWithHTTPClient creates a new IssueReplaceLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIssueReplaceLabelsParamsWithHTTPClient(client *http.Client) *IssueReplaceLabelsParams {
	var ()
	return &IssueReplaceLabelsParams{
		HTTPClient: client,
	}
}

/*IssueReplaceLabelsParams contains all the parameters to send to the API endpoint
for the issue replace labels operation typically these are written to a http.Request
*/
type IssueReplaceLabelsParams struct {

	/*Body*/
	Body *models.IssueLabelsOption
	/*Index
	  index of the issue

	*/
	Index int64
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

// WithTimeout adds the timeout to the issue replace labels params
func (o *IssueReplaceLabelsParams) WithTimeout(timeout time.Duration) *IssueReplaceLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue replace labels params
func (o *IssueReplaceLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue replace labels params
func (o *IssueReplaceLabelsParams) WithContext(ctx context.Context) *IssueReplaceLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue replace labels params
func (o *IssueReplaceLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue replace labels params
func (o *IssueReplaceLabelsParams) WithHTTPClient(client *http.Client) *IssueReplaceLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue replace labels params
func (o *IssueReplaceLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue replace labels params
func (o *IssueReplaceLabelsParams) WithBody(body *models.IssueLabelsOption) *IssueReplaceLabelsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue replace labels params
func (o *IssueReplaceLabelsParams) SetBody(body *models.IssueLabelsOption) {
	o.Body = body
}

// WithIndex adds the index to the issue replace labels params
func (o *IssueReplaceLabelsParams) WithIndex(index int64) *IssueReplaceLabelsParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the issue replace labels params
func (o *IssueReplaceLabelsParams) SetIndex(index int64) {
	o.Index = index
}

// WithOwner adds the owner to the issue replace labels params
func (o *IssueReplaceLabelsParams) WithOwner(owner string) *IssueReplaceLabelsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue replace labels params
func (o *IssueReplaceLabelsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue replace labels params
func (o *IssueReplaceLabelsParams) WithRepo(repo string) *IssueReplaceLabelsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue replace labels params
func (o *IssueReplaceLabelsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueReplaceLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
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
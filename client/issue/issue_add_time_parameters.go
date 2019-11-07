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

// NewIssueAddTimeParams creates a new IssueAddTimeParams object
// with the default values initialized.
func NewIssueAddTimeParams() *IssueAddTimeParams {
	var ()
	return &IssueAddTimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIssueAddTimeParamsWithTimeout creates a new IssueAddTimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIssueAddTimeParamsWithTimeout(timeout time.Duration) *IssueAddTimeParams {
	var ()
	return &IssueAddTimeParams{

		timeout: timeout,
	}
}

// NewIssueAddTimeParamsWithContext creates a new IssueAddTimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewIssueAddTimeParamsWithContext(ctx context.Context) *IssueAddTimeParams {
	var ()
	return &IssueAddTimeParams{

		Context: ctx,
	}
}

// NewIssueAddTimeParamsWithHTTPClient creates a new IssueAddTimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIssueAddTimeParamsWithHTTPClient(client *http.Client) *IssueAddTimeParams {
	var ()
	return &IssueAddTimeParams{
		HTTPClient: client,
	}
}

/*IssueAddTimeParams contains all the parameters to send to the API endpoint
for the issue add time operation typically these are written to a http.Request
*/
type IssueAddTimeParams struct {

	/*Body*/
	Body *models.AddTimeOption
	/*ID
	  index of the issue to add tracked time to

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

// WithTimeout adds the timeout to the issue add time params
func (o *IssueAddTimeParams) WithTimeout(timeout time.Duration) *IssueAddTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue add time params
func (o *IssueAddTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue add time params
func (o *IssueAddTimeParams) WithContext(ctx context.Context) *IssueAddTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue add time params
func (o *IssueAddTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue add time params
func (o *IssueAddTimeParams) WithHTTPClient(client *http.Client) *IssueAddTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue add time params
func (o *IssueAddTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue add time params
func (o *IssueAddTimeParams) WithBody(body *models.AddTimeOption) *IssueAddTimeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue add time params
func (o *IssueAddTimeParams) SetBody(body *models.AddTimeOption) {
	o.Body = body
}

// WithID adds the id to the issue add time params
func (o *IssueAddTimeParams) WithID(id int64) *IssueAddTimeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue add time params
func (o *IssueAddTimeParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the issue add time params
func (o *IssueAddTimeParams) WithOwner(owner string) *IssueAddTimeParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue add time params
func (o *IssueAddTimeParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue add time params
func (o *IssueAddTimeParams) WithRepo(repo string) *IssueAddTimeParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue add time params
func (o *IssueAddTimeParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueAddTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

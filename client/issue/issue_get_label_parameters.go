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
)

// NewIssueGetLabelParams creates a new IssueGetLabelParams object
// with the default values initialized.
func NewIssueGetLabelParams() *IssueGetLabelParams {
	var ()
	return &IssueGetLabelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIssueGetLabelParamsWithTimeout creates a new IssueGetLabelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIssueGetLabelParamsWithTimeout(timeout time.Duration) *IssueGetLabelParams {
	var ()
	return &IssueGetLabelParams{

		timeout: timeout,
	}
}

// NewIssueGetLabelParamsWithContext creates a new IssueGetLabelParams object
// with the default values initialized, and the ability to set a context for a request
func NewIssueGetLabelParamsWithContext(ctx context.Context) *IssueGetLabelParams {
	var ()
	return &IssueGetLabelParams{

		Context: ctx,
	}
}

// NewIssueGetLabelParamsWithHTTPClient creates a new IssueGetLabelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIssueGetLabelParamsWithHTTPClient(client *http.Client) *IssueGetLabelParams {
	var ()
	return &IssueGetLabelParams{
		HTTPClient: client,
	}
}

/*IssueGetLabelParams contains all the parameters to send to the API endpoint
for the issue get label operation typically these are written to a http.Request
*/
type IssueGetLabelParams struct {

	/*ID
	  id of the label to get

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

// WithTimeout adds the timeout to the issue get label params
func (o *IssueGetLabelParams) WithTimeout(timeout time.Duration) *IssueGetLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue get label params
func (o *IssueGetLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue get label params
func (o *IssueGetLabelParams) WithContext(ctx context.Context) *IssueGetLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue get label params
func (o *IssueGetLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue get label params
func (o *IssueGetLabelParams) WithHTTPClient(client *http.Client) *IssueGetLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue get label params
func (o *IssueGetLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the issue get label params
func (o *IssueGetLabelParams) WithID(id int64) *IssueGetLabelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue get label params
func (o *IssueGetLabelParams) SetID(id int64) {
	o.ID = id
}

// WithOwner adds the owner to the issue get label params
func (o *IssueGetLabelParams) WithOwner(owner string) *IssueGetLabelParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue get label params
func (o *IssueGetLabelParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue get label params
func (o *IssueGetLabelParams) WithRepo(repo string) *IssueGetLabelParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue get label params
func (o *IssueGetLabelParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *IssueGetLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
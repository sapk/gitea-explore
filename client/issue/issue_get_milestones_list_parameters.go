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

	strfmt "github.com/go-openapi/strfmt"
)

// NewIssueGetMilestonesListParams creates a new IssueGetMilestonesListParams object
// with the default values initialized.
func NewIssueGetMilestonesListParams() *IssueGetMilestonesListParams {
	var ()
	return &IssueGetMilestonesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIssueGetMilestonesListParamsWithTimeout creates a new IssueGetMilestonesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIssueGetMilestonesListParamsWithTimeout(timeout time.Duration) *IssueGetMilestonesListParams {
	var ()
	return &IssueGetMilestonesListParams{

		timeout: timeout,
	}
}

// NewIssueGetMilestonesListParamsWithContext creates a new IssueGetMilestonesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIssueGetMilestonesListParamsWithContext(ctx context.Context) *IssueGetMilestonesListParams {
	var ()
	return &IssueGetMilestonesListParams{

		Context: ctx,
	}
}

// NewIssueGetMilestonesListParamsWithHTTPClient creates a new IssueGetMilestonesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIssueGetMilestonesListParamsWithHTTPClient(client *http.Client) *IssueGetMilestonesListParams {
	var ()
	return &IssueGetMilestonesListParams{
		HTTPClient: client,
	}
}

/*IssueGetMilestonesListParams contains all the parameters to send to the API endpoint
for the issue get milestones list operation typically these are written to a http.Request
*/
type IssueGetMilestonesListParams struct {

	/*Owner
	  owner of the repo

	*/
	Owner string
	/*Repo
	  name of the repo

	*/
	Repo string
	/*State
	  Milestone state, Recognised values are open, closed and all. Defaults to "open"

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the issue get milestones list params
func (o *IssueGetMilestonesListParams) WithTimeout(timeout time.Duration) *IssueGetMilestonesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue get milestones list params
func (o *IssueGetMilestonesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue get milestones list params
func (o *IssueGetMilestonesListParams) WithContext(ctx context.Context) *IssueGetMilestonesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue get milestones list params
func (o *IssueGetMilestonesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue get milestones list params
func (o *IssueGetMilestonesListParams) WithHTTPClient(client *http.Client) *IssueGetMilestonesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue get milestones list params
func (o *IssueGetMilestonesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the issue get milestones list params
func (o *IssueGetMilestonesListParams) WithOwner(owner string) *IssueGetMilestonesListParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the issue get milestones list params
func (o *IssueGetMilestonesListParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithRepo adds the repo to the issue get milestones list params
func (o *IssueGetMilestonesListParams) WithRepo(repo string) *IssueGetMilestonesListParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the issue get milestones list params
func (o *IssueGetMilestonesListParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithState adds the state to the issue get milestones list params
func (o *IssueGetMilestonesListParams) WithState(state *string) *IssueGetMilestonesListParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the issue get milestones list params
func (o *IssueGetMilestonesListParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *IssueGetMilestonesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

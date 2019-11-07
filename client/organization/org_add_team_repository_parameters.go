// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewOrgAddTeamRepositoryParams creates a new OrgAddTeamRepositoryParams object
// with the default values initialized.
func NewOrgAddTeamRepositoryParams() *OrgAddTeamRepositoryParams {
	var ()
	return &OrgAddTeamRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgAddTeamRepositoryParamsWithTimeout creates a new OrgAddTeamRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgAddTeamRepositoryParamsWithTimeout(timeout time.Duration) *OrgAddTeamRepositoryParams {
	var ()
	return &OrgAddTeamRepositoryParams{

		timeout: timeout,
	}
}

// NewOrgAddTeamRepositoryParamsWithContext creates a new OrgAddTeamRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgAddTeamRepositoryParamsWithContext(ctx context.Context) *OrgAddTeamRepositoryParams {
	var ()
	return &OrgAddTeamRepositoryParams{

		Context: ctx,
	}
}

// NewOrgAddTeamRepositoryParamsWithHTTPClient creates a new OrgAddTeamRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgAddTeamRepositoryParamsWithHTTPClient(client *http.Client) *OrgAddTeamRepositoryParams {
	var ()
	return &OrgAddTeamRepositoryParams{
		HTTPClient: client,
	}
}

/*OrgAddTeamRepositoryParams contains all the parameters to send to the API endpoint
for the org add team repository operation typically these are written to a http.Request
*/
type OrgAddTeamRepositoryParams struct {

	/*ID
	  id of the team

	*/
	ID int64
	/*Org
	  organization that owns the repo to add

	*/
	Org string
	/*Repo
	  name of the repo to add

	*/
	Repo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org add team repository params
func (o *OrgAddTeamRepositoryParams) WithTimeout(timeout time.Duration) *OrgAddTeamRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org add team repository params
func (o *OrgAddTeamRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org add team repository params
func (o *OrgAddTeamRepositoryParams) WithContext(ctx context.Context) *OrgAddTeamRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org add team repository params
func (o *OrgAddTeamRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org add team repository params
func (o *OrgAddTeamRepositoryParams) WithHTTPClient(client *http.Client) *OrgAddTeamRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org add team repository params
func (o *OrgAddTeamRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the org add team repository params
func (o *OrgAddTeamRepositoryParams) WithID(id int64) *OrgAddTeamRepositoryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org add team repository params
func (o *OrgAddTeamRepositoryParams) SetID(id int64) {
	o.ID = id
}

// WithOrg adds the org to the org add team repository params
func (o *OrgAddTeamRepositoryParams) WithOrg(org string) *OrgAddTeamRepositoryParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org add team repository params
func (o *OrgAddTeamRepositoryParams) SetOrg(org string) {
	o.Org = org
}

// WithRepo adds the repo to the org add team repository params
func (o *OrgAddTeamRepositoryParams) WithRepo(repo string) *OrgAddTeamRepositoryParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the org add team repository params
func (o *OrgAddTeamRepositoryParams) SetRepo(repo string) {
	o.Repo = repo
}

// WriteToRequest writes these params to a swagger request
func (o *OrgAddTeamRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
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
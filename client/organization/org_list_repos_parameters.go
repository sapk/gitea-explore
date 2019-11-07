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

	strfmt "github.com/go-openapi/strfmt"
)

// NewOrgListReposParams creates a new OrgListReposParams object
// with the default values initialized.
func NewOrgListReposParams() *OrgListReposParams {
	var ()
	return &OrgListReposParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListReposParamsWithTimeout creates a new OrgListReposParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgListReposParamsWithTimeout(timeout time.Duration) *OrgListReposParams {
	var ()
	return &OrgListReposParams{

		timeout: timeout,
	}
}

// NewOrgListReposParamsWithContext creates a new OrgListReposParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgListReposParamsWithContext(ctx context.Context) *OrgListReposParams {
	var ()
	return &OrgListReposParams{

		Context: ctx,
	}
}

// NewOrgListReposParamsWithHTTPClient creates a new OrgListReposParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgListReposParamsWithHTTPClient(client *http.Client) *OrgListReposParams {
	var ()
	return &OrgListReposParams{
		HTTPClient: client,
	}
}

/*OrgListReposParams contains all the parameters to send to the API endpoint
for the org list repos operation typically these are written to a http.Request
*/
type OrgListReposParams struct {

	/*Org
	  name of the organization

	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org list repos params
func (o *OrgListReposParams) WithTimeout(timeout time.Duration) *OrgListReposParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list repos params
func (o *OrgListReposParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list repos params
func (o *OrgListReposParams) WithContext(ctx context.Context) *OrgListReposParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list repos params
func (o *OrgListReposParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list repos params
func (o *OrgListReposParams) WithHTTPClient(client *http.Client) *OrgListReposParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list repos params
func (o *OrgListReposParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the org list repos params
func (o *OrgListReposParams) WithOrg(org string) *OrgListReposParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org list repos params
func (o *OrgListReposParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

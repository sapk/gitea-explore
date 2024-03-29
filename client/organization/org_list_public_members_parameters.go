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

// NewOrgListPublicMembersParams creates a new OrgListPublicMembersParams object
// with the default values initialized.
func NewOrgListPublicMembersParams() *OrgListPublicMembersParams {
	var ()
	return &OrgListPublicMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListPublicMembersParamsWithTimeout creates a new OrgListPublicMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgListPublicMembersParamsWithTimeout(timeout time.Duration) *OrgListPublicMembersParams {
	var ()
	return &OrgListPublicMembersParams{

		timeout: timeout,
	}
}

// NewOrgListPublicMembersParamsWithContext creates a new OrgListPublicMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgListPublicMembersParamsWithContext(ctx context.Context) *OrgListPublicMembersParams {
	var ()
	return &OrgListPublicMembersParams{

		Context: ctx,
	}
}

// NewOrgListPublicMembersParamsWithHTTPClient creates a new OrgListPublicMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgListPublicMembersParamsWithHTTPClient(client *http.Client) *OrgListPublicMembersParams {
	var ()
	return &OrgListPublicMembersParams{
		HTTPClient: client,
	}
}

/*OrgListPublicMembersParams contains all the parameters to send to the API endpoint
for the org list public members operation typically these are written to a http.Request
*/
type OrgListPublicMembersParams struct {

	/*Org
	  name of the organization

	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org list public members params
func (o *OrgListPublicMembersParams) WithTimeout(timeout time.Duration) *OrgListPublicMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list public members params
func (o *OrgListPublicMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list public members params
func (o *OrgListPublicMembersParams) WithContext(ctx context.Context) *OrgListPublicMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list public members params
func (o *OrgListPublicMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list public members params
func (o *OrgListPublicMembersParams) WithHTTPClient(client *http.Client) *OrgListPublicMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list public members params
func (o *OrgListPublicMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the org list public members params
func (o *OrgListPublicMembersParams) WithOrg(org string) *OrgListPublicMembersParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org list public members params
func (o *OrgListPublicMembersParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListPublicMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

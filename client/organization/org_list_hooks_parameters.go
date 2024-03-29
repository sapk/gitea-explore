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

// NewOrgListHooksParams creates a new OrgListHooksParams object
// with the default values initialized.
func NewOrgListHooksParams() *OrgListHooksParams {
	var ()
	return &OrgListHooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListHooksParamsWithTimeout creates a new OrgListHooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgListHooksParamsWithTimeout(timeout time.Duration) *OrgListHooksParams {
	var ()
	return &OrgListHooksParams{

		timeout: timeout,
	}
}

// NewOrgListHooksParamsWithContext creates a new OrgListHooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgListHooksParamsWithContext(ctx context.Context) *OrgListHooksParams {
	var ()
	return &OrgListHooksParams{

		Context: ctx,
	}
}

// NewOrgListHooksParamsWithHTTPClient creates a new OrgListHooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgListHooksParamsWithHTTPClient(client *http.Client) *OrgListHooksParams {
	var ()
	return &OrgListHooksParams{
		HTTPClient: client,
	}
}

/*OrgListHooksParams contains all the parameters to send to the API endpoint
for the org list hooks operation typically these are written to a http.Request
*/
type OrgListHooksParams struct {

	/*Org
	  name of the organization

	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org list hooks params
func (o *OrgListHooksParams) WithTimeout(timeout time.Duration) *OrgListHooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list hooks params
func (o *OrgListHooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list hooks params
func (o *OrgListHooksParams) WithContext(ctx context.Context) *OrgListHooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list hooks params
func (o *OrgListHooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list hooks params
func (o *OrgListHooksParams) WithHTTPClient(client *http.Client) *OrgListHooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list hooks params
func (o *OrgListHooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrg adds the org to the org list hooks params
func (o *OrgListHooksParams) WithOrg(org string) *OrgListHooksParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org list hooks params
func (o *OrgListHooksParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListHooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

	models "gitea.com/sapk/explore/models"
)

// NewOrgEditParams creates a new OrgEditParams object
// with the default values initialized.
func NewOrgEditParams() *OrgEditParams {
	var ()
	return &OrgEditParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgEditParamsWithTimeout creates a new OrgEditParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgEditParamsWithTimeout(timeout time.Duration) *OrgEditParams {
	var ()
	return &OrgEditParams{

		timeout: timeout,
	}
}

// NewOrgEditParamsWithContext creates a new OrgEditParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgEditParamsWithContext(ctx context.Context) *OrgEditParams {
	var ()
	return &OrgEditParams{

		Context: ctx,
	}
}

// NewOrgEditParamsWithHTTPClient creates a new OrgEditParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgEditParamsWithHTTPClient(client *http.Client) *OrgEditParams {
	var ()
	return &OrgEditParams{
		HTTPClient: client,
	}
}

/*OrgEditParams contains all the parameters to send to the API endpoint
for the org edit operation typically these are written to a http.Request
*/
type OrgEditParams struct {

	/*Body*/
	Body *models.EditOrgOption
	/*Org
	  name of the organization to edit

	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org edit params
func (o *OrgEditParams) WithTimeout(timeout time.Duration) *OrgEditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org edit params
func (o *OrgEditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org edit params
func (o *OrgEditParams) WithContext(ctx context.Context) *OrgEditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org edit params
func (o *OrgEditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org edit params
func (o *OrgEditParams) WithHTTPClient(client *http.Client) *OrgEditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org edit params
func (o *OrgEditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the org edit params
func (o *OrgEditParams) WithBody(body *models.EditOrgOption) *OrgEditParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the org edit params
func (o *OrgEditParams) SetBody(body *models.EditOrgOption) {
	o.Body = body
}

// WithOrg adds the org to the org edit params
func (o *OrgEditParams) WithOrg(org string) *OrgEditParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the org edit params
func (o *OrgEditParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *OrgEditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

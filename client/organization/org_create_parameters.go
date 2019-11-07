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

// NewOrgCreateParams creates a new OrgCreateParams object
// with the default values initialized.
func NewOrgCreateParams() *OrgCreateParams {
	var ()
	return &OrgCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgCreateParamsWithTimeout creates a new OrgCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgCreateParamsWithTimeout(timeout time.Duration) *OrgCreateParams {
	var ()
	return &OrgCreateParams{

		timeout: timeout,
	}
}

// NewOrgCreateParamsWithContext creates a new OrgCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgCreateParamsWithContext(ctx context.Context) *OrgCreateParams {
	var ()
	return &OrgCreateParams{

		Context: ctx,
	}
}

// NewOrgCreateParamsWithHTTPClient creates a new OrgCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgCreateParamsWithHTTPClient(client *http.Client) *OrgCreateParams {
	var ()
	return &OrgCreateParams{
		HTTPClient: client,
	}
}

/*OrgCreateParams contains all the parameters to send to the API endpoint
for the org create operation typically these are written to a http.Request
*/
type OrgCreateParams struct {

	/*Organization*/
	Organization *models.CreateOrgOption

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org create params
func (o *OrgCreateParams) WithTimeout(timeout time.Duration) *OrgCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org create params
func (o *OrgCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org create params
func (o *OrgCreateParams) WithContext(ctx context.Context) *OrgCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org create params
func (o *OrgCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org create params
func (o *OrgCreateParams) WithHTTPClient(client *http.Client) *OrgCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org create params
func (o *OrgCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganization adds the organization to the org create params
func (o *OrgCreateParams) WithOrganization(organization *models.CreateOrgOption) *OrgCreateParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the org create params
func (o *OrgCreateParams) SetOrganization(organization *models.CreateOrgOption) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *OrgCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Organization != nil {
		if err := r.SetBodyParam(o.Organization); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
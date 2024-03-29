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

// NewOrgListTeamMembersParams creates a new OrgListTeamMembersParams object
// with the default values initialized.
func NewOrgListTeamMembersParams() *OrgListTeamMembersParams {
	var ()
	return &OrgListTeamMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrgListTeamMembersParamsWithTimeout creates a new OrgListTeamMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrgListTeamMembersParamsWithTimeout(timeout time.Duration) *OrgListTeamMembersParams {
	var ()
	return &OrgListTeamMembersParams{

		timeout: timeout,
	}
}

// NewOrgListTeamMembersParamsWithContext creates a new OrgListTeamMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrgListTeamMembersParamsWithContext(ctx context.Context) *OrgListTeamMembersParams {
	var ()
	return &OrgListTeamMembersParams{

		Context: ctx,
	}
}

// NewOrgListTeamMembersParamsWithHTTPClient creates a new OrgListTeamMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOrgListTeamMembersParamsWithHTTPClient(client *http.Client) *OrgListTeamMembersParams {
	var ()
	return &OrgListTeamMembersParams{
		HTTPClient: client,
	}
}

/*OrgListTeamMembersParams contains all the parameters to send to the API endpoint
for the org list team members operation typically these are written to a http.Request
*/
type OrgListTeamMembersParams struct {

	/*ID
	  id of the team

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the org list team members params
func (o *OrgListTeamMembersParams) WithTimeout(timeout time.Duration) *OrgListTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the org list team members params
func (o *OrgListTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the org list team members params
func (o *OrgListTeamMembersParams) WithContext(ctx context.Context) *OrgListTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the org list team members params
func (o *OrgListTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the org list team members params
func (o *OrgListTeamMembersParams) WithHTTPClient(client *http.Client) *OrgListTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the org list team members params
func (o *OrgListTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the org list team members params
func (o *OrgListTeamMembersParams) WithID(id int64) *OrgListTeamMembersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the org list team members params
func (o *OrgListTeamMembersParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrgListTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminCreateRepoParams creates a new AdminCreateRepoParams object
// with the default values initialized.
func NewAdminCreateRepoParams() *AdminCreateRepoParams {
	var ()
	return &AdminCreateRepoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminCreateRepoParamsWithTimeout creates a new AdminCreateRepoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminCreateRepoParamsWithTimeout(timeout time.Duration) *AdminCreateRepoParams {
	var ()
	return &AdminCreateRepoParams{

		timeout: timeout,
	}
}

// NewAdminCreateRepoParamsWithContext creates a new AdminCreateRepoParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminCreateRepoParamsWithContext(ctx context.Context) *AdminCreateRepoParams {
	var ()
	return &AdminCreateRepoParams{

		Context: ctx,
	}
}

// NewAdminCreateRepoParamsWithHTTPClient creates a new AdminCreateRepoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminCreateRepoParamsWithHTTPClient(client *http.Client) *AdminCreateRepoParams {
	var ()
	return &AdminCreateRepoParams{
		HTTPClient: client,
	}
}

/*AdminCreateRepoParams contains all the parameters to send to the API endpoint
for the admin create repo operation typically these are written to a http.Request
*/
type AdminCreateRepoParams struct {

	/*Repository*/
	Repository *models.CreateRepoOption
	/*Username
	  username of the user. This user will own the created repository

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin create repo params
func (o *AdminCreateRepoParams) WithTimeout(timeout time.Duration) *AdminCreateRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin create repo params
func (o *AdminCreateRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin create repo params
func (o *AdminCreateRepoParams) WithContext(ctx context.Context) *AdminCreateRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin create repo params
func (o *AdminCreateRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin create repo params
func (o *AdminCreateRepoParams) WithHTTPClient(client *http.Client) *AdminCreateRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin create repo params
func (o *AdminCreateRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the admin create repo params
func (o *AdminCreateRepoParams) WithRepository(repository *models.CreateRepoOption) *AdminCreateRepoParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the admin create repo params
func (o *AdminCreateRepoParams) SetRepository(repository *models.CreateRepoOption) {
	o.Repository = repository
}

// WithUsername adds the username to the admin create repo params
func (o *AdminCreateRepoParams) WithUsername(username string) *AdminCreateRepoParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the admin create repo params
func (o *AdminCreateRepoParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *AdminCreateRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Repository != nil {
		if err := r.SetBodyParam(o.Repository); err != nil {
			return err
		}
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

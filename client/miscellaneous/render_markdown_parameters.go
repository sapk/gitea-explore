// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

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

// NewRenderMarkdownParams creates a new RenderMarkdownParams object
// with the default values initialized.
func NewRenderMarkdownParams() *RenderMarkdownParams {
	var ()
	return &RenderMarkdownParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRenderMarkdownParamsWithTimeout creates a new RenderMarkdownParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRenderMarkdownParamsWithTimeout(timeout time.Duration) *RenderMarkdownParams {
	var ()
	return &RenderMarkdownParams{

		timeout: timeout,
	}
}

// NewRenderMarkdownParamsWithContext creates a new RenderMarkdownParams object
// with the default values initialized, and the ability to set a context for a request
func NewRenderMarkdownParamsWithContext(ctx context.Context) *RenderMarkdownParams {
	var ()
	return &RenderMarkdownParams{

		Context: ctx,
	}
}

// NewRenderMarkdownParamsWithHTTPClient creates a new RenderMarkdownParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRenderMarkdownParamsWithHTTPClient(client *http.Client) *RenderMarkdownParams {
	var ()
	return &RenderMarkdownParams{
		HTTPClient: client,
	}
}

/*RenderMarkdownParams contains all the parameters to send to the API endpoint
for the render markdown operation typically these are written to a http.Request
*/
type RenderMarkdownParams struct {

	/*Body*/
	Body *models.MarkdownOption

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the render markdown params
func (o *RenderMarkdownParams) WithTimeout(timeout time.Duration) *RenderMarkdownParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the render markdown params
func (o *RenderMarkdownParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the render markdown params
func (o *RenderMarkdownParams) WithContext(ctx context.Context) *RenderMarkdownParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the render markdown params
func (o *RenderMarkdownParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the render markdown params
func (o *RenderMarkdownParams) WithHTTPClient(client *http.Client) *RenderMarkdownParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the render markdown params
func (o *RenderMarkdownParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the render markdown params
func (o *RenderMarkdownParams) WithBody(body *models.MarkdownOption) *RenderMarkdownParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the render markdown params
func (o *RenderMarkdownParams) SetBody(body *models.MarkdownOption) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RenderMarkdownParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewRepoSearchParams creates a new RepoSearchParams object
// with the default values initialized.
func NewRepoSearchParams() *RepoSearchParams {
	var ()
	return &RepoSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRepoSearchParamsWithTimeout creates a new RepoSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRepoSearchParamsWithTimeout(timeout time.Duration) *RepoSearchParams {
	var ()
	return &RepoSearchParams{

		timeout: timeout,
	}
}

// NewRepoSearchParamsWithContext creates a new RepoSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewRepoSearchParamsWithContext(ctx context.Context) *RepoSearchParams {
	var ()
	return &RepoSearchParams{

		Context: ctx,
	}
}

// NewRepoSearchParamsWithHTTPClient creates a new RepoSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRepoSearchParamsWithHTTPClient(client *http.Client) *RepoSearchParams {
	var ()
	return &RepoSearchParams{
		HTTPClient: client,
	}
}

/*RepoSearchParams contains all the parameters to send to the API endpoint
for the repo search operation typically these are written to a http.Request
*/
type RepoSearchParams struct {

	/*Exclusive
	  if `uid` is given, search only for repos that the user owns

	*/
	Exclusive *bool
	/*IncludeDesc
	  include search of keyword within repository description

	*/
	IncludeDesc *bool
	/*Limit
	  page size of results, maximum page size is 50

	*/
	Limit *int64
	/*Mode
	  type of repository to search for. Supported values are "fork", "source", "mirror" and "collaborative"

	*/
	Mode *string
	/*Order
	  sort order, either "asc" (ascending) or "desc" (descending). Default is "asc", ignored if "sort" is not specified.

	*/
	Order *string
	/*Page
	  page number of results to return (1-based)

	*/
	Page *int64
	/*Private
	  include private repositories this user has access to (defaults to true)

	*/
	Private *bool
	/*Q
	  keyword

	*/
	Q *string
	/*Sort
	  sort repos by attribute. Supported values are "alpha", "created", "updated", "size", and "id". Default is "alpha"

	*/
	Sort *string
	/*StarredBy
	  search only for repos that the user with the given id has starred

	*/
	StarredBy *int64
	/*Topic
	  Limit search to repositories with keyword as topic

	*/
	Topic *bool
	/*UID
	  search only for repos that the user with the given id owns or contributes to

	*/
	UID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the repo search params
func (o *RepoSearchParams) WithTimeout(timeout time.Duration) *RepoSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repo search params
func (o *RepoSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repo search params
func (o *RepoSearchParams) WithContext(ctx context.Context) *RepoSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repo search params
func (o *RepoSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repo search params
func (o *RepoSearchParams) WithHTTPClient(client *http.Client) *RepoSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repo search params
func (o *RepoSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclusive adds the exclusive to the repo search params
func (o *RepoSearchParams) WithExclusive(exclusive *bool) *RepoSearchParams {
	o.SetExclusive(exclusive)
	return o
}

// SetExclusive adds the exclusive to the repo search params
func (o *RepoSearchParams) SetExclusive(exclusive *bool) {
	o.Exclusive = exclusive
}

// WithIncludeDesc adds the includeDesc to the repo search params
func (o *RepoSearchParams) WithIncludeDesc(includeDesc *bool) *RepoSearchParams {
	o.SetIncludeDesc(includeDesc)
	return o
}

// SetIncludeDesc adds the includeDesc to the repo search params
func (o *RepoSearchParams) SetIncludeDesc(includeDesc *bool) {
	o.IncludeDesc = includeDesc
}

// WithLimit adds the limit to the repo search params
func (o *RepoSearchParams) WithLimit(limit *int64) *RepoSearchParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the repo search params
func (o *RepoSearchParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMode adds the mode to the repo search params
func (o *RepoSearchParams) WithMode(mode *string) *RepoSearchParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the repo search params
func (o *RepoSearchParams) SetMode(mode *string) {
	o.Mode = mode
}

// WithOrder adds the order to the repo search params
func (o *RepoSearchParams) WithOrder(order *string) *RepoSearchParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the repo search params
func (o *RepoSearchParams) SetOrder(order *string) {
	o.Order = order
}

// WithPage adds the page to the repo search params
func (o *RepoSearchParams) WithPage(page *int64) *RepoSearchParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the repo search params
func (o *RepoSearchParams) SetPage(page *int64) {
	o.Page = page
}

// WithPrivate adds the private to the repo search params
func (o *RepoSearchParams) WithPrivate(private *bool) *RepoSearchParams {
	o.SetPrivate(private)
	return o
}

// SetPrivate adds the private to the repo search params
func (o *RepoSearchParams) SetPrivate(private *bool) {
	o.Private = private
}

// WithQ adds the q to the repo search params
func (o *RepoSearchParams) WithQ(q *string) *RepoSearchParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the repo search params
func (o *RepoSearchParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the repo search params
func (o *RepoSearchParams) WithSort(sort *string) *RepoSearchParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the repo search params
func (o *RepoSearchParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStarredBy adds the starredBy to the repo search params
func (o *RepoSearchParams) WithStarredBy(starredBy *int64) *RepoSearchParams {
	o.SetStarredBy(starredBy)
	return o
}

// SetStarredBy adds the starredBy to the repo search params
func (o *RepoSearchParams) SetStarredBy(starredBy *int64) {
	o.StarredBy = starredBy
}

// WithTopic adds the topic to the repo search params
func (o *RepoSearchParams) WithTopic(topic *bool) *RepoSearchParams {
	o.SetTopic(topic)
	return o
}

// SetTopic adds the topic to the repo search params
func (o *RepoSearchParams) SetTopic(topic *bool) {
	o.Topic = topic
}

// WithUID adds the uid to the repo search params
func (o *RepoSearchParams) WithUID(uid *int64) *RepoSearchParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the repo search params
func (o *RepoSearchParams) SetUID(uid *int64) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *RepoSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Exclusive != nil {

		// query param exclusive
		var qrExclusive bool
		if o.Exclusive != nil {
			qrExclusive = *o.Exclusive
		}
		qExclusive := swag.FormatBool(qrExclusive)
		if qExclusive != "" {
			if err := r.SetQueryParam("exclusive", qExclusive); err != nil {
				return err
			}
		}

	}

	if o.IncludeDesc != nil {

		// query param includeDesc
		var qrIncludeDesc bool
		if o.IncludeDesc != nil {
			qrIncludeDesc = *o.IncludeDesc
		}
		qIncludeDesc := swag.FormatBool(qrIncludeDesc)
		if qIncludeDesc != "" {
			if err := r.SetQueryParam("includeDesc", qIncludeDesc); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Mode != nil {

		// query param mode
		var qrMode string
		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {
			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}

	}

	if o.Order != nil {

		// query param order
		var qrOrder string
		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {
			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Private != nil {

		// query param private
		var qrPrivate bool
		if o.Private != nil {
			qrPrivate = *o.Private
		}
		qPrivate := swag.FormatBool(qrPrivate)
		if qPrivate != "" {
			if err := r.SetQueryParam("private", qPrivate); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if o.StarredBy != nil {

		// query param starredBy
		var qrStarredBy int64
		if o.StarredBy != nil {
			qrStarredBy = *o.StarredBy
		}
		qStarredBy := swag.FormatInt64(qrStarredBy)
		if qStarredBy != "" {
			if err := r.SetQueryParam("starredBy", qStarredBy); err != nil {
				return err
			}
		}

	}

	if o.Topic != nil {

		// query param topic
		var qrTopic bool
		if o.Topic != nil {
			qrTopic = *o.Topic
		}
		qTopic := swag.FormatBool(qrTopic)
		if qTopic != "" {
			if err := r.SetQueryParam("topic", qTopic); err != nil {
				return err
			}
		}

	}

	if o.UID != nil {

		// query param uid
		var qrUID int64
		if o.UID != nil {
			qrUID = *o.UID
		}
		qUID := swag.FormatInt64(qrUID)
		if qUID != "" {
			if err := r.SetQueryParam("uid", qUID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

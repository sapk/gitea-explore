// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TopicResponse TopicResponse for returning topics
// swagger:model TopicResponse
type TopicResponse struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"topic_name,omitempty"`

	// repo count
	RepoCount int64 `json:"repo_count,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this topic response
func (m *TopicResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopicResponse) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TopicResponse) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopicResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopicResponse) UnmarshalBinary(b []byte) error {
	var res TopicResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

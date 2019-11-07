// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InternalTracker InternalTracker represents settings for internal tracker
// swagger:model InternalTracker
type InternalTracker struct {

	// Let only contributors track time (Built-in issue tracker)
	AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time,omitempty"`

	// Enable dependencies for issues and pull requests (Built-in issue tracker)
	EnableIssueDependencies bool `json:"enable_issue_dependencies,omitempty"`

	// Enable time tracking (Built-in issue tracker)
	EnableTimeTracker bool `json:"enable_time_tracker,omitempty"`
}

// Validate validates this internal tracker
func (m *InternalTracker) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InternalTracker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalTracker) UnmarshalBinary(b []byte) error {
	var res InternalTracker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
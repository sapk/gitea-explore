// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserHeatmapData UserHeatmapData represents the data needed to create a heatmap
// swagger:model UserHeatmapData
type UserHeatmapData struct {

	// contributions
	Contributions int64 `json:"contributions,omitempty"`

	// timestamp
	Timestamp TimeStamp `json:"timestamp,omitempty"`
}

// Validate validates this user heatmap data
func (m *UserHeatmapData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserHeatmapData) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := m.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserHeatmapData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserHeatmapData) UnmarshalBinary(b []byte) error {
	var res UserHeatmapData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

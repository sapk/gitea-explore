// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Branch Branch represents a repository branch
// swagger:model Branch
type Branch struct {

	// name
	Name string `json:"name,omitempty"`

	// commit
	Commit *PayloadCommit `json:"commit,omitempty"`
}

// Validate validates this branch
func (m *Branch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Branch) validateCommit(formats strfmt.Registry) error {

	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {
		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Branch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Branch) UnmarshalBinary(b []byte) error {
	var res Branch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
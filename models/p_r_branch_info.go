// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PRBranchInfo PRBranchInfo information about a branch
// swagger:model PRBranchInfo
type PRBranchInfo struct {

	// name
	Name string `json:"label,omitempty"`

	// ref
	Ref string `json:"ref,omitempty"`

	// repo ID
	RepoID int64 `json:"repo_id,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// repo
	Repo *Repository `json:"repo,omitempty"`
}

// Validate validates this p r branch info
func (m *PRBranchInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PRBranchInfo) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {
		if err := m.Repo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PRBranchInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PRBranchInfo) UnmarshalBinary(b []byte) error {
	var res PRBranchInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// DeleteFileOptions DeleteFileOptions options for deleting files (used for other File structs below)
// Note: `author` and `committer` are optional (if only one is given, it will be used for the other, otherwise the authenticated user will be used)
// swagger:model DeleteFileOptions
type DeleteFileOptions struct {

	// branch (optional) to base this file from. if not given, the default branch is used
	BranchName string `json:"branch,omitempty"`

	// message (optional) for the commit of this file. if not supplied, a default message will be used
	Message string `json:"message,omitempty"`

	// new_branch (optional) will make a new branch from `branch` before creating the file
	NewBranchName string `json:"new_branch,omitempty"`

	// sha is the SHA for the file that already exists
	// Required: true
	SHA *string `json:"sha"`

	// author
	Author *Identity `json:"author,omitempty"`

	// committer
	Committer *Identity `json:"committer,omitempty"`
}

// Validate validates this delete file options
func (m *DeleteFileOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSHA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteFileOptions) validateSHA(formats strfmt.Registry) error {

	if err := validate.Required("sha", "body", m.SHA); err != nil {
		return err
	}

	return nil
}

func (m *DeleteFileOptions) validateAuthor(formats strfmt.Registry) error {

	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {
		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

func (m *DeleteFileOptions) validateCommitter(formats strfmt.Registry) error {

	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {
		if err := m.Committer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("committer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFileOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFileOptions) UnmarshalBinary(b []byte) error {
	var res DeleteFileOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

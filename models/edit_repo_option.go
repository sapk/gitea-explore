// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EditRepoOption EditRepoOption options when editing a repository's properties
// swagger:model EditRepoOption
type EditRepoOption struct {

	// either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits. `has_pull_requests` must be `true`.
	AllowMerge bool `json:"allow_merge_commits,omitempty"`

	// either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging. `has_pull_requests` must be `true`.
	AllowRebase bool `json:"allow_rebase,omitempty"`

	// either `true` to allow rebase with explicit merge commits (--no-ff), or `false` to prevent rebase with explicit merge commits. `has_pull_requests` must be `true`.
	AllowRebaseMerge bool `json:"allow_rebase_explicit,omitempty"`

	// either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging. `has_pull_requests` must be `true`.
	AllowSquash bool `json:"allow_squash_merge,omitempty"`

	// set to `true` to archive this repository.
	Archived bool `json:"archived,omitempty"`

	// sets the default branch for this repository.
	DefaultBranch string `json:"default_branch,omitempty"`

	// a short description of the repository.
	Description string `json:"description,omitempty"`

	// either `true` to enable issues for this repository or `false` to disable them.
	HasIssues bool `json:"has_issues,omitempty"`

	// either `true` to allow pull requests, or `false` to prevent pull request.
	HasPullRequests bool `json:"has_pull_requests,omitempty"`

	// either `true` to enable the wiki for this repository or `false` to disable it.
	HasWiki bool `json:"has_wiki,omitempty"`

	// either `true` to ignore whitespace for conflicts, or `false` to not ignore whitespace. `has_pull_requests` must be `true`.
	IgnoreWhitespaceConflicts bool `json:"ignore_whitespace_conflicts,omitempty"`

	// name of the repository
	// Unique: true
	Name string `json:"name,omitempty"`

	// either `true` to make the repository private or `false` to make it public.
	// Note: you will get a 422 error if the organization restricts changing repository visibility to organization
	// owners and a non-owner tries to change the value of private.
	Private bool `json:"private,omitempty"`

	// a URL with more information about the repository.
	Website string `json:"website,omitempty"`

	// external tracker
	ExternalTracker *ExternalTracker `json:"external_tracker,omitempty"`

	// external wiki
	ExternalWiki *ExternalWiki `json:"external_wiki,omitempty"`

	// internal tracker
	InternalTracker *InternalTracker `json:"internal_tracker,omitempty"`
}

// Validate validates this edit repo option
func (m *EditRepoOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalTracker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalWiki(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalTracker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditRepoOption) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	return nil
}

func (m *EditRepoOption) validateExternalTracker(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalTracker) { // not required
		return nil
	}

	if m.ExternalTracker != nil {
		if err := m.ExternalTracker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_tracker")
			}
			return err
		}
	}

	return nil
}

func (m *EditRepoOption) validateExternalWiki(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalWiki) { // not required
		return nil
	}

	if m.ExternalWiki != nil {
		if err := m.ExternalWiki.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_wiki")
			}
			return err
		}
	}

	return nil
}

func (m *EditRepoOption) validateInternalTracker(formats strfmt.Registry) error {

	if swag.IsZero(m.InternalTracker) { // not required
		return nil
	}

	if m.InternalTracker != nil {
		if err := m.InternalTracker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("internal_tracker")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditRepoOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditRepoOption) UnmarshalBinary(b []byte) error {
	var res EditRepoOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

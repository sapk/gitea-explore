// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContentsResponse ContentsResponse contains information about a repo's entry's (dir, file, symlink, submodule) metadata and content
// swagger:model ContentsResponse
type ContentsResponse struct {

	// `content` is populated when `type` is `file`, otherwise null
	Content string `json:"content,omitempty"`

	// download URL
	DownloadURL string `json:"download_url,omitempty"`

	// `encoding` is populated when `type` is `file`, otherwise null
	Encoding string `json:"encoding,omitempty"`

	// git URL
	GitURL string `json:"git_url,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// s h a
	SHA string `json:"sha,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// `submodule_git_url` is populated when `type` is `submodule`, otherwise null
	SubmoduleGitURL string `json:"submodule_git_url,omitempty"`

	// `target` is populated when `type` is `symlink`, otherwise null
	Target string `json:"target,omitempty"`

	// `type` will be `file`, `dir`, `symlink`, or `submodule`
	Type string `json:"type,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// links
	Links *FileLinksResponse `json:"_links,omitempty"`
}

// Validate validates this contents response
func (m *ContentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentsResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentsResponse) UnmarshalBinary(b []byte) error {
	var res ContentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
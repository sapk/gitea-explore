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

// Attachment Attachment a generic attachment
// swagger:model Attachment
type Attachment struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// download count
	DownloadCount int64 `json:"download_count,omitempty"`

	// download URL
	DownloadURL string `json:"browser_download_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// UUID
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this attachment
func (m *Attachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Attachment) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Attachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attachment) UnmarshalBinary(b []byte) error {
	var res Attachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
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

// PublicKey PublicKey publickey is a user key to push code to repository
// swagger:model PublicKey
type PublicKey struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// key type
	KeyType string `json:"key_type,omitempty"`

	// read only
	ReadOnly bool `json:"read_only,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this public key
func (m *PublicKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicKey) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicKey) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicKey) UnmarshalBinary(b []byte) error {
	var res PublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

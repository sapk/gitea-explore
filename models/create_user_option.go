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

// CreateUserOption CreateUserOption create user options
// swagger:model CreateUserOption
type CreateUserOption struct {

	// email
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// login name
	LoginName string `json:"login_name,omitempty"`

	// must change password
	MustChangePassword bool `json:"must_change_password,omitempty"`

	// password
	// Required: true
	Password *string `json:"password"`

	// send notify
	SendNotify bool `json:"send_notify,omitempty"`

	// source ID
	SourceID int64 `json:"source_id,omitempty"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this create user option
func (m *CreateUserOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateUserOption) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateUserOption) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CreateUserOption) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateUserOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUserOption) UnmarshalBinary(b []byte) error {
	var res CreateUserOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

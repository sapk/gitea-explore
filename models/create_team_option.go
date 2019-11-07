// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateTeamOption CreateTeamOption options for creating a team
// swagger:model CreateTeamOption
type CreateTeamOption struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// permission
	// Enum: [read write admin]
	Permission string `json:"permission,omitempty"`

	// units
	Units []string `json:"units"`
}

// Validate validates this create team option
func (m *CreateTeamOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTeamOption) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var createTeamOptionTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createTeamOptionTypePermissionPropEnum = append(createTeamOptionTypePermissionPropEnum, v)
	}
}

const (

	// CreateTeamOptionPermissionRead captures enum value "read"
	CreateTeamOptionPermissionRead string = "read"

	// CreateTeamOptionPermissionWrite captures enum value "write"
	CreateTeamOptionPermissionWrite string = "write"

	// CreateTeamOptionPermissionAdmin captures enum value "admin"
	CreateTeamOptionPermissionAdmin string = "admin"
)

// prop value enum
func (m *CreateTeamOption) validatePermissionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createTeamOptionTypePermissionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateTeamOption) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionEnum("permission", "body", m.Permission); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTeamOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTeamOption) UnmarshalBinary(b []byte) error {
	var res CreateTeamOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

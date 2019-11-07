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

// EditDeadlineOption EditDeadlineOption options for creating a deadline
// swagger:model EditDeadlineOption
type EditDeadlineOption struct {

	// deadline
	// Required: true
	// Format: date-time
	Deadline *strfmt.DateTime `json:"due_date"`
}

// Validate validates this edit deadline option
func (m *EditDeadlineOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditDeadlineOption) validateDeadline(formats strfmt.Registry) error {

	if err := validate.Required("due_date", "body", m.Deadline); err != nil {
		return err
	}

	if err := validate.FormatOf("due_date", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditDeadlineOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditDeadlineOption) UnmarshalBinary(b []byte) error {
	var res EditDeadlineOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

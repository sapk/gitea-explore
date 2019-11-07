// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AnnotatedTag AnnotatedTag represents an annotated tag
// swagger:model AnnotatedTag
type AnnotatedTag struct {

	// message
	Message string `json:"message,omitempty"`

	// s h a
	SHA string `json:"sha,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// object
	Object *AnnotatedTagObject `json:"object,omitempty"`

	// tagger
	Tagger *CommitUser `json:"tagger,omitempty"`

	// verification
	Verification *PayloadCommitVerification `json:"verification,omitempty"`
}

// Validate validates this annotated tag
func (m *AnnotatedTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnnotatedTag) validateObject(formats strfmt.Registry) error {

	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *AnnotatedTag) validateTagger(formats strfmt.Registry) error {

	if swag.IsZero(m.Tagger) { // not required
		return nil
	}

	if m.Tagger != nil {
		if err := m.Tagger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tagger")
			}
			return err
		}
	}

	return nil
}

func (m *AnnotatedTag) validateVerification(formats strfmt.Registry) error {

	if swag.IsZero(m.Verification) { // not required
		return nil
	}

	if m.Verification != nil {
		if err := m.Verification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnnotatedTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnnotatedTag) UnmarshalBinary(b []byte) error {
	var res AnnotatedTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
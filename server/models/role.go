// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Role role
// swagger:model Role
type Role struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// users
	Users []*User `json:"users"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {

		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {

			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

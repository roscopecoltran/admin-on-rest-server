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

// Field field
// swagger:model Field
type Field struct {

	// choices
	Choices []IChoiceItem `json:"choices"`

	// component
	Component string `json:"component,omitempty"`

	// data source Id
	DataSourceID string `json:"dataSourceId,omitempty"`

	// default value
	DefaultValue string `json:"defaultValue,omitempty"`

	// eid
	Eid string `json:"eid,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// input type
	InputType string `json:"inputType,omitempty"`

	// is auto incremented
	IsAutoIncremented bool `json:"isAutoIncremented,omitempty"`

	// is part of primary key
	IsPartOfPrimaryKey bool `json:"isPartOfPrimaryKey,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// max length
	MaxLength int32 `json:"maxLength,omitempty"`

	// max value
	MaxValue string `json:"maxValue,omitempty"`

	// min value
	MinValue string `json:"minValue,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`

	// reference option text
	ReferenceOptionText string `json:"referenceOptionText,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// show in create
	ShowInCreate bool `json:"showInCreate,omitempty"`

	// show in edit
	ShowInEdit bool `json:"showInEdit,omitempty"`

	// show in filter
	ShowInFilter bool `json:"showInFilter,omitempty"`

	// show in list
	ShowInList bool `json:"showInList,omitempty"`

	// show in show
	ShowInShow bool `json:"showInShow,omitempty"`
}

// Validate validates this field
func (m *Field) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChoices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInputType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Field) validateChoices(formats strfmt.Registry) error {

	if swag.IsZero(m.Choices) { // not required
		return nil
	}

	for i := 0; i < len(m.Choices); i++ {

	}

	return nil
}

var fieldTypeComponentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Autocomplete","Boolean","NullableBoolean","CheckboxGroup","Date","Disabled","File","Image","LongText","Number","RadioButtonGroup","ReferenceArray","Reference","RichText","SelectArray","Select","Text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldTypeComponentPropEnum = append(fieldTypeComponentPropEnum, v)
	}
}

const (
	// FieldComponentAutocomplete captures enum value "Autocomplete"
	FieldComponentAutocomplete string = "Autocomplete"
	// FieldComponentBoolean captures enum value "Boolean"
	FieldComponentBoolean string = "Boolean"
	// FieldComponentNullableBoolean captures enum value "NullableBoolean"
	FieldComponentNullableBoolean string = "NullableBoolean"
	// FieldComponentCheckboxGroup captures enum value "CheckboxGroup"
	FieldComponentCheckboxGroup string = "CheckboxGroup"
	// FieldComponentDate captures enum value "Date"
	FieldComponentDate string = "Date"
	// FieldComponentDisabled captures enum value "Disabled"
	FieldComponentDisabled string = "Disabled"
	// FieldComponentFile captures enum value "File"
	FieldComponentFile string = "File"
	// FieldComponentImage captures enum value "Image"
	FieldComponentImage string = "Image"
	// FieldComponentLongText captures enum value "LongText"
	FieldComponentLongText string = "LongText"
	// FieldComponentNumber captures enum value "Number"
	FieldComponentNumber string = "Number"
	// FieldComponentRadioButtonGroup captures enum value "RadioButtonGroup"
	FieldComponentRadioButtonGroup string = "RadioButtonGroup"
	// FieldComponentReferenceArray captures enum value "ReferenceArray"
	FieldComponentReferenceArray string = "ReferenceArray"
	// FieldComponentReference captures enum value "Reference"
	FieldComponentReference string = "Reference"
	// FieldComponentRichText captures enum value "RichText"
	FieldComponentRichText string = "RichText"
	// FieldComponentSelectArray captures enum value "SelectArray"
	FieldComponentSelectArray string = "SelectArray"
	// FieldComponentSelect captures enum value "Select"
	FieldComponentSelect string = "Select"
	// FieldComponentText captures enum value "Text"
	FieldComponentText string = "Text"
)

// prop value enum
func (m *Field) validateComponentEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fieldTypeComponentPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Field) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	// value enum
	if err := m.validateComponentEnum("component", "body", m.Component); err != nil {
		return err
	}

	return nil
}

var fieldTypeInputTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text","email","password","url"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldTypeInputTypePropEnum = append(fieldTypeInputTypePropEnum, v)
	}
}

const (
	// FieldInputTypeText captures enum value "text"
	FieldInputTypeText string = "text"
	// FieldInputTypeEmail captures enum value "email"
	FieldInputTypeEmail string = "email"
	// FieldInputTypePassword captures enum value "password"
	FieldInputTypePassword string = "password"
	// FieldInputTypeURL captures enum value "url"
	FieldInputTypeURL string = "url"
)

// prop value enum
func (m *Field) validateInputTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fieldTypeInputTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Field) validateInputType(formats strfmt.Registry) error {

	if swag.IsZero(m.InputType) { // not required
		return nil
	}

	// value enum
	if err := m.validateInputTypeEnum("inputType", "body", m.InputType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Field) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Field) UnmarshalBinary(b []byte) error {
	var res Field
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

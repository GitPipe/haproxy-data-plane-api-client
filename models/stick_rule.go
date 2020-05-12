// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// StickRule Stick Rule
//
// Define a pattern used to create an entry in a stickiness table or matching condition or associate a user to a server.
// swagger:model stick_rule
type StickRule struct {

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// pattern
	// Required: true
	// Pattern: ^[^\s]+$
	Pattern string `json:"pattern"`

	// table
	// Pattern: ^[^\s]+$
	Table string `json:"table,omitempty"`

	// type
	// Required: true
	// Enum: [match on store-request store-response]
	Type string `json:"type"`
}

// Validate validates this stick rule
func (m *StickRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var stickRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stickRuleTypeCondPropEnum = append(stickRuleTypeCondPropEnum, v)
	}
}

const (

	// StickRuleCondIf captures enum value "if"
	StickRuleCondIf string = "if"

	// StickRuleCondUnless captures enum value "unless"
	StickRuleCondUnless string = "unless"
)

// prop value enum
func (m *StickRule) validateCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stickRuleTypeCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StickRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *StickRule) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *StickRule) validatePattern(formats strfmt.Registry) error {

	if err := validate.RequiredString("pattern", "body", string(m.Pattern)); err != nil {
		return err
	}

	if err := validate.Pattern("pattern", "body", string(m.Pattern), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *StickRule) validateTable(formats strfmt.Registry) error {

	if swag.IsZero(m.Table) { // not required
		return nil
	}

	if err := validate.Pattern("table", "body", string(m.Table), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var stickRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["match","on","store-request","store-response"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stickRuleTypeTypePropEnum = append(stickRuleTypeTypePropEnum, v)
	}
}

const (

	// StickRuleTypeMatch captures enum value "match"
	StickRuleTypeMatch string = "match"

	// StickRuleTypeOn captures enum value "on"
	StickRuleTypeOn string = "on"

	// StickRuleTypeStoreRequest captures enum value "store-request"
	StickRuleTypeStoreRequest string = "store-request"

	// StickRuleTypeStoreResponse captures enum value "store-response"
	StickRuleTypeStoreResponse string = "store-response"
)

// prop value enum
func (m *StickRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stickRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StickRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StickRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StickRule) UnmarshalBinary(b []byte) error {
	var res StickRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
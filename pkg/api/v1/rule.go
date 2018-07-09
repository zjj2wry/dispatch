///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package v1

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NO TESTS

// Rule rule
// swagger:model Rule
type Rule struct {

	// actions
	// Required: true
	Actions []string `json:"actions"`

	// resources
	// Required: true
	Resources []string `json:"resources"`

	// subjects
	// Required: true
	Subjects []string `json:"subjects"`
}

// Validate validates this rule
func (m *Rule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubjects(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ruleActionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["get","create","update","delete","*"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ruleActionsItemsEnum = append(ruleActionsItemsEnum, v)
	}
}

func (m *Rule) validateActionsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ruleActionsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *Rule) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	for i := 0; i < len(m.Actions); i++ {

		// value enum
		if err := m.validateActionsItemsEnum("actions"+"."+strconv.Itoa(i), "body", m.Actions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Rule) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	for i := 0; i < len(m.Resources); i++ {

		if err := validate.Pattern("resources"+"."+strconv.Itoa(i), "body", string(m.Resources[i]), `^[\w\d\-\*]+$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *Rule) validateSubjects(formats strfmt.Registry) error {

	if err := validate.Required("subjects", "body", m.Subjects); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule) UnmarshalBinary(b []byte) error {
	var res Rule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
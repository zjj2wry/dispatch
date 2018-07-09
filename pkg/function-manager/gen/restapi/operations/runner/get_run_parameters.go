///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRunParams creates a new GetRunParams object
// no default values defined in spec.
func NewGetRunParams() GetRunParams {

	return GetRunParams{}
}

// GetRunParams contains all the bound params for the get run operation
// typically these are obtained from a http.Request
//
// swagger:parameters getRun
type GetRunParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XDispatchOrg string
	/*Name of function to retreive a run for
	  Pattern: ^[\w\d\-]+$
	  In: query
	*/
	FunctionName *string
	/*name of run to retrieve
	  Required: true
	  In: path
	*/
	RunName strfmt.UUID
	/*Retreive runs modified since given Unix time
	  In: query
	*/
	Since *int64
	/*Filter based on tags
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRunParams() beforehand.
func (o *GetRunParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXDispatchOrg(r.Header[http.CanonicalHeaderKey("X-Dispatch-Org")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qFunctionName, qhkFunctionName, _ := qs.GetOK("functionName")
	if err := o.bindFunctionName(qFunctionName, qhkFunctionName, route.Formats); err != nil {
		res = append(res, err)
	}

	rRunName, rhkRunName, _ := route.Params.GetOK("runName")
	if err := o.bindRunName(rRunName, rhkRunName, route.Formats); err != nil {
		res = append(res, err)
	}

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRunParams) bindXDispatchOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Dispatch-Org", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Dispatch-Org", "header", raw); err != nil {
		return err
	}

	o.XDispatchOrg = raw

	return nil
}

func (o *GetRunParams) bindFunctionName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FunctionName = &raw

	if err := o.validateFunctionName(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetRunParams) validateFunctionName(formats strfmt.Registry) error {

	if err := validate.Pattern("functionName", "query", (*o.FunctionName), `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (o *GetRunParams) bindRunName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("runName", "path", "strfmt.UUID", raw)
	}
	o.RunName = *(value.(*strfmt.UUID))

	if err := o.validateRunName(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetRunParams) validateRunName(formats strfmt.Registry) error {

	if err := validate.FormatOf("runName", "path", "uuid", o.RunName.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetRunParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("since", "query", "int64", raw)
	}
	o.Since = &value

	return nil
}

func (o *GetRunParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	tagsIC := rawData

	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetResultParams creates a new GetResultParams object
// no default values defined in spec.
func NewGetResultParams() GetResultParams {

	return GetResultParams{}
}

// GetResultParams contains all the bound params for the get result operation
// typically these are obtained from a http.Request
//
// swagger:parameters getResult
type GetResultParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: header
	*/
	XSPLATHONAPITOKEN *string
	/*
	  Required: true
	  In: path
	*/
	EventID int64
	/*team id
	  In: query
	*/
	TeamID *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetResultParams() beforehand.
func (o *GetResultParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXSPLATHONAPITOKEN(r.Header[http.CanonicalHeaderKey("X-SPLATHON-API-TOKEN")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rEventID, rhkEventID, _ := route.Params.GetOK("eventId")
	if err := o.bindEventID(rEventID, rhkEventID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTeamID, qhkTeamID, _ := qs.GetOK("team_id")
	if err := o.bindTeamID(qTeamID, qhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXSPLATHONAPITOKEN binds and validates parameter XSPLATHONAPITOKEN from header.
func (o *GetResultParams) bindXSPLATHONAPITOKEN(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XSPLATHONAPITOKEN = &raw

	return nil
}

// bindEventID binds and validates parameter EventID from path.
func (o *GetResultParams) bindEventID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("eventId", "path", "int64", raw)
	}
	o.EventID = value

	return nil
}

// bindTeamID binds and validates parameter TeamID from query.
func (o *GetResultParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("team_id", "query", "int64", raw)
	}
	o.TeamID = &value

	return nil
}

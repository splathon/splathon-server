// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCompleteReceptionParams creates a new CompleteReceptionParams object
// no default values defined in spec.
func NewCompleteReceptionParams() CompleteReceptionParams {

	return CompleteReceptionParams{}
}

// CompleteReceptionParams contains all the bound params for the complete reception operation
// typically these are obtained from a http.Request
//
// swagger:parameters completeReception
type CompleteReceptionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XSPLATHONAPITOKEN string
	/*
	  Required: true
	  In: path
	*/
	EventID int64
	/*ReceptionResponse.splathon.code と同じもの(たぶん内部SlackID).
	  Required: true
	  In: path
	*/
	SplathonReceptionCode string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCompleteReceptionParams() beforehand.
func (o *CompleteReceptionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXSPLATHONAPITOKEN(r.Header[http.CanonicalHeaderKey("X-SPLATHON-API-TOKEN")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rEventID, rhkEventID, _ := route.Params.GetOK("eventId")
	if err := o.bindEventID(rEventID, rhkEventID, route.Formats); err != nil {
		res = append(res, err)
	}

	rSplathonReceptionCode, rhkSplathonReceptionCode, _ := route.Params.GetOK("splathonReceptionCode")
	if err := o.bindSplathonReceptionCode(rSplathonReceptionCode, rhkSplathonReceptionCode, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXSPLATHONAPITOKEN binds and validates parameter XSPLATHONAPITOKEN from header.
func (o *CompleteReceptionParams) bindXSPLATHONAPITOKEN(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-SPLATHON-API-TOKEN", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-SPLATHON-API-TOKEN", "header", raw); err != nil {
		return err
	}

	o.XSPLATHONAPITOKEN = raw

	return nil
}

// bindEventID binds and validates parameter EventID from path.
func (o *CompleteReceptionParams) bindEventID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindSplathonReceptionCode binds and validates parameter SplathonReceptionCode from path.
func (o *CompleteReceptionParams) bindSplathonReceptionCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.SplathonReceptionCode = raw

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/splathon/splathon-server/swagger/models"
)

// NewUpdateMatchParams creates a new UpdateMatchParams object
// no default values defined in spec.
func NewUpdateMatchParams() UpdateMatchParams {

	return UpdateMatchParams{}
}

// UpdateMatchParams contains all the bound params for the update match operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateMatch
type UpdateMatchParams struct {

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
	/*
	  Required: true
	  In: body
	*/
	Match *models.NewMatchRequest
	/*match id
	  Required: true
	  In: path
	*/
	MatchID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateMatchParams() beforehand.
func (o *UpdateMatchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXSPLATHONAPITOKEN(r.Header[http.CanonicalHeaderKey("X-SPLATHON-API-TOKEN")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rEventID, rhkEventID, _ := route.Params.GetOK("eventId")
	if err := o.bindEventID(rEventID, rhkEventID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.NewMatchRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("match", "body"))
			} else {
				res = append(res, errors.NewParseError("match", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Match = &body
			}
		}
	} else {
		res = append(res, errors.Required("match", "body"))
	}
	rMatchID, rhkMatchID, _ := route.Params.GetOK("matchId")
	if err := o.bindMatchID(rMatchID, rhkMatchID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXSPLATHONAPITOKEN binds and validates parameter XSPLATHONAPITOKEN from header.
func (o *UpdateMatchParams) bindXSPLATHONAPITOKEN(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *UpdateMatchParams) bindEventID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindMatchID binds and validates parameter MatchID from path.
func (o *UpdateMatchParams) bindMatchID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("matchId", "path", "int64", raw)
	}
	o.MatchID = value

	return nil
}

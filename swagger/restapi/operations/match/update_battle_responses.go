// Code generated by go-swagger; DO NOT EDIT.

package match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/splathon/splathon-server/swagger/models"
)

// UpdateBattleOKCode is the HTTP code returned for type UpdateBattleOK
const UpdateBattleOKCode int = 200

/*UpdateBattleOK Success

swagger:response updateBattleOK
*/
type UpdateBattleOK struct {

	/*
	  In: Body
	*/
	Payload *models.Match `json:"body,omitempty"`
}

// NewUpdateBattleOK creates UpdateBattleOK with default headers values
func NewUpdateBattleOK() *UpdateBattleOK {

	return &UpdateBattleOK{}
}

// WithPayload adds the payload to the update battle o k response
func (o *UpdateBattleOK) WithPayload(payload *models.Match) *UpdateBattleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update battle o k response
func (o *UpdateBattleOK) SetPayload(payload *models.Match) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBattleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateBattleDefault Generic error

swagger:response updateBattleDefault
*/
type UpdateBattleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateBattleDefault creates UpdateBattleDefault with default headers values
func NewUpdateBattleDefault(code int) *UpdateBattleDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateBattleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update battle default response
func (o *UpdateBattleDefault) WithStatusCode(code int) *UpdateBattleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update battle default response
func (o *UpdateBattleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update battle default response
func (o *UpdateBattleDefault) WithPayload(payload *models.Error) *UpdateBattleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update battle default response
func (o *UpdateBattleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBattleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

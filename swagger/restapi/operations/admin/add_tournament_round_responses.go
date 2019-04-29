// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/splathon/splathon-server/swagger/models"
)

// AddTournamentRoundOKCode is the HTTP code returned for type AddTournamentRoundOK
const AddTournamentRoundOKCode int = 200

/*AddTournamentRoundOK Success

swagger:response addTournamentRoundOK
*/
type AddTournamentRoundOK struct {
}

// NewAddTournamentRoundOK creates AddTournamentRoundOK with default headers values
func NewAddTournamentRoundOK() *AddTournamentRoundOK {

	return &AddTournamentRoundOK{}
}

// WriteResponse to the client
func (o *AddTournamentRoundOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*AddTournamentRoundDefault Generic error

swagger:response addTournamentRoundDefault
*/
type AddTournamentRoundDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddTournamentRoundDefault creates AddTournamentRoundDefault with default headers values
func NewAddTournamentRoundDefault(code int) *AddTournamentRoundDefault {
	if code <= 0 {
		code = 500
	}

	return &AddTournamentRoundDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add tournament round default response
func (o *AddTournamentRoundDefault) WithStatusCode(code int) *AddTournamentRoundDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add tournament round default response
func (o *AddTournamentRoundDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add tournament round default response
func (o *AddTournamentRoundDefault) WithPayload(payload *models.Error) *AddTournamentRoundDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add tournament round default response
func (o *AddTournamentRoundDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddTournamentRoundDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/splathon/splathon-server/swagger/models"
)

// ListNoticesOKCode is the HTTP code returned for type ListNoticesOK
const ListNoticesOKCode int = 200

/*ListNoticesOK Success

swagger:response listNoticesOK
*/
type ListNoticesOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListNoticesResponse `json:"body,omitempty"`
}

// NewListNoticesOK creates ListNoticesOK with default headers values
func NewListNoticesOK() *ListNoticesOK {

	return &ListNoticesOK{}
}

// WithPayload adds the payload to the list notices o k response
func (o *ListNoticesOK) WithPayload(payload *models.ListNoticesResponse) *ListNoticesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list notices o k response
func (o *ListNoticesOK) SetPayload(payload *models.ListNoticesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNoticesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListNoticesDefault Generic error

swagger:response listNoticesDefault
*/
type ListNoticesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListNoticesDefault creates ListNoticesDefault with default headers values
func NewListNoticesDefault(code int) *ListNoticesDefault {
	if code <= 0 {
		code = 500
	}

	return &ListNoticesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list notices default response
func (o *ListNoticesDefault) WithStatusCode(code int) *ListNoticesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list notices default response
func (o *ListNoticesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list notices default response
func (o *ListNoticesDefault) WithPayload(payload *models.Error) *ListNoticesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list notices default response
func (o *ListNoticesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNoticesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

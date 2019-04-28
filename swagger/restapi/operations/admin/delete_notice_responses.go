// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/splathon/splathon-server/swagger/models"
)

// DeleteNoticeOKCode is the HTTP code returned for type DeleteNoticeOK
const DeleteNoticeOKCode int = 200

/*DeleteNoticeOK Success

swagger:response deleteNoticeOK
*/
type DeleteNoticeOK struct {
}

// NewDeleteNoticeOK creates DeleteNoticeOK with default headers values
func NewDeleteNoticeOK() *DeleteNoticeOK {

	return &DeleteNoticeOK{}
}

// WriteResponse to the client
func (o *DeleteNoticeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteNoticeDefault Generic error

swagger:response deleteNoticeDefault
*/
type DeleteNoticeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteNoticeDefault creates DeleteNoticeDefault with default headers values
func NewDeleteNoticeDefault(code int) *DeleteNoticeDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteNoticeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete notice default response
func (o *DeleteNoticeDefault) WithStatusCode(code int) *DeleteNoticeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete notice default response
func (o *DeleteNoticeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete notice default response
func (o *DeleteNoticeDefault) WithPayload(payload *models.Error) *DeleteNoticeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete notice default response
func (o *DeleteNoticeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNoticeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

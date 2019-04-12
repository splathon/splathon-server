// Code generated by go-swagger; DO NOT EDIT.

package match

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/splathon/splathon-server/swagger/models"
)

// GetMatchOKCode is the HTTP code returned for type GetMatchOK
const GetMatchOKCode int = 200

/*GetMatchOK Success

swagger:response getMatchOK
*/
type GetMatchOK struct {

	/*
	  In: Body
	*/
	Payload *models.Match `json:"body,omitempty"`
}

// NewGetMatchOK creates GetMatchOK with default headers values
func NewGetMatchOK() *GetMatchOK {

	return &GetMatchOK{}
}

// WithPayload adds the payload to the get match o k response
func (o *GetMatchOK) WithPayload(payload *models.Match) *GetMatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get match o k response
func (o *GetMatchOK) SetPayload(payload *models.Match) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
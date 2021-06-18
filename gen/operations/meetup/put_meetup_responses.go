// Code generated by go-swagger; DO NOT EDIT.

package meetup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudnative-id/community-system/models"
)

// PutMeetupOKCode is the HTTP code returned for type PutMeetupOK
const PutMeetupOKCode int = 200

/*PutMeetupOK object uuid & status

swagger:response putMeetupOK
*/
type PutMeetupOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CreateObject `json:"body,omitempty"`
}

// NewPutMeetupOK creates PutMeetupOK with default headers values
func NewPutMeetupOK() *PutMeetupOK {

	return &PutMeetupOK{}
}

// WithPayload adds the payload to the put meetup o k response
func (o *PutMeetupOK) WithPayload(payload []*models.CreateObject) *PutMeetupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put meetup o k response
func (o *PutMeetupOK) SetPayload(payload []*models.CreateObject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMeetupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CreateObject, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PutMeetupDefault Unexpected error

swagger:response putMeetupDefault
*/
type PutMeetupDefault struct {
	_statusCode int
}

// NewPutMeetupDefault creates PutMeetupDefault with default headers values
func NewPutMeetupDefault(code int) *PutMeetupDefault {
	if code <= 0 {
		code = 500
	}

	return &PutMeetupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put meetup default response
func (o *PutMeetupDefault) WithStatusCode(code int) *PutMeetupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put meetup default response
func (o *PutMeetupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PutMeetupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}

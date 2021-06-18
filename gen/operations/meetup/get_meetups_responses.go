// Code generated by go-swagger; DO NOT EDIT.

package meetup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudnative-id/community-system/models"
)

// GetMeetupsOKCode is the HTTP code returned for type GetMeetupsOK
const GetMeetupsOKCode int = 200

/*GetMeetupsOK A JSON array of meetups

swagger:response getMeetupsOK
*/
type GetMeetupsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Meetup `json:"body,omitempty"`
}

// NewGetMeetupsOK creates GetMeetupsOK with default headers values
func NewGetMeetupsOK() *GetMeetupsOK {

	return &GetMeetupsOK{}
}

// WithPayload adds the payload to the get meetups o k response
func (o *GetMeetupsOK) WithPayload(payload []*models.Meetup) *GetMeetupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get meetups o k response
func (o *GetMeetupsOK) SetPayload(payload []*models.Meetup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeetupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Meetup, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetMeetupsDefault Unexpected error

swagger:response getMeetupsDefault
*/
type GetMeetupsDefault struct {
	_statusCode int
}

// NewGetMeetupsDefault creates GetMeetupsDefault with default headers values
func NewGetMeetupsDefault(code int) *GetMeetupsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMeetupsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get meetups default response
func (o *GetMeetupsDefault) WithStatusCode(code int) *GetMeetupsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get meetups default response
func (o *GetMeetupsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *GetMeetupsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}

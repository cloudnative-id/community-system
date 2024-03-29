// Code generated by go-swagger; DO NOT EDIT.

package sponsor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloudnative-id/community-system/gen/models"
)

// GetSponsorsOKCode is the HTTP code returned for type GetSponsorsOK
const GetSponsorsOKCode int = 200

/*GetSponsorsOK Success

swagger:response getSponsorsOK
*/
type GetSponsorsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Sponsor `json:"body,omitempty"`
}

// NewGetSponsorsOK creates GetSponsorsOK with default headers values
func NewGetSponsorsOK() *GetSponsorsOK {

	return &GetSponsorsOK{}
}

// WithPayload adds the payload to the get sponsors o k response
func (o *GetSponsorsOK) WithPayload(payload []*models.Sponsor) *GetSponsorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sponsors o k response
func (o *GetSponsorsOK) SetPayload(payload []*models.Sponsor) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSponsorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Sponsor, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSponsorsNotFoundCode is the HTTP code returned for type GetSponsorsNotFound
const GetSponsorsNotFoundCode int = 404

/*GetSponsorsNotFound Speaker not found

swagger:response getSponsorsNotFound
*/
type GetSponsorsNotFound struct {
}

// NewGetSponsorsNotFound creates GetSponsorsNotFound with default headers values
func NewGetSponsorsNotFound() *GetSponsorsNotFound {

	return &GetSponsorsNotFound{}
}

// WriteResponse to the client
func (o *GetSponsorsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetSponsorsDefault Unexpected error

swagger:response getSponsorsDefault
*/
type GetSponsorsDefault struct {
	_statusCode int
}

// NewGetSponsorsDefault creates GetSponsorsDefault with default headers values
func NewGetSponsorsDefault(code int) *GetSponsorsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSponsorsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sponsors default response
func (o *GetSponsorsDefault) WithStatusCode(code int) *GetSponsorsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sponsors default response
func (o *GetSponsorsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *GetSponsorsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}

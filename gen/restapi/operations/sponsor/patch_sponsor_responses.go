// Code generated by go-swagger; DO NOT EDIT.

package sponsor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PatchSponsorNoContentCode is the HTTP code returned for type PatchSponsorNoContent
const PatchSponsorNoContentCode int = 204

/*PatchSponsorNoContent updated

swagger:response patchSponsorNoContent
*/
type PatchSponsorNoContent struct {
}

// NewPatchSponsorNoContent creates PatchSponsorNoContent with default headers values
func NewPatchSponsorNoContent() *PatchSponsorNoContent {

	return &PatchSponsorNoContent{}
}

// WriteResponse to the client
func (o *PatchSponsorNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*PatchSponsorDefault Unexpected error

swagger:response patchSponsorDefault
*/
type PatchSponsorDefault struct {
	_statusCode int
}

// NewPatchSponsorDefault creates PatchSponsorDefault with default headers values
func NewPatchSponsorDefault(code int) *PatchSponsorDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchSponsorDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch sponsor default response
func (o *PatchSponsorDefault) WithStatusCode(code int) *PatchSponsorDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch sponsor default response
func (o *PatchSponsorDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PatchSponsorDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
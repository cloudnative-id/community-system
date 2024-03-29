// Code generated by go-swagger; DO NOT EDIT.

package speaker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchSpeakerHandlerFunc turns a function with the right signature into a patch speaker handler
type PatchSpeakerHandlerFunc func(PatchSpeakerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchSpeakerHandlerFunc) Handle(params PatchSpeakerParams) middleware.Responder {
	return fn(params)
}

// PatchSpeakerHandler interface for that can handle valid patch speaker params
type PatchSpeakerHandler interface {
	Handle(PatchSpeakerParams) middleware.Responder
}

// NewPatchSpeaker creates a new http.Handler for the patch speaker operation
func NewPatchSpeaker(ctx *middleware.Context, handler PatchSpeakerHandler) *PatchSpeaker {
	return &PatchSpeaker{Context: ctx, Handler: handler}
}

/* PatchSpeaker swagger:route PATCH /meetups/{meetup_id}/speakers/{speaker_id} speaker patchSpeaker

Patch speaker data.

update speaker

*/
type PatchSpeaker struct {
	Context *middleware.Context
	Handler PatchSpeakerHandler
}

func (o *PatchSpeaker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchSpeakerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

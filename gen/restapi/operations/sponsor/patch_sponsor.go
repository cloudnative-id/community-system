// Code generated by go-swagger; DO NOT EDIT.

package sponsor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchSponsorHandlerFunc turns a function with the right signature into a patch sponsor handler
type PatchSponsorHandlerFunc func(PatchSponsorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchSponsorHandlerFunc) Handle(params PatchSponsorParams) middleware.Responder {
	return fn(params)
}

// PatchSponsorHandler interface for that can handle valid patch sponsor params
type PatchSponsorHandler interface {
	Handle(PatchSponsorParams) middleware.Responder
}

// NewPatchSponsor creates a new http.Handler for the patch sponsor operation
func NewPatchSponsor(ctx *middleware.Context, handler PatchSponsorHandler) *PatchSponsor {
	return &PatchSponsor{Context: ctx, Handler: handler}
}

/* PatchSponsor swagger:route PATCH /meetups/sponsors/{sponsor_id} sponsor patchSponsor

Patch sponsor data.

update sponsor

*/
type PatchSponsor struct {
	Context *middleware.Context
	Handler PatchSponsorHandler
}

func (o *PatchSponsor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchSponsorParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

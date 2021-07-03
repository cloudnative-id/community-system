// Code generated by go-swagger; DO NOT EDIT.

package speaker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSpeakersHandlerFunc turns a function with the right signature into a get speakers handler
type GetSpeakersHandlerFunc func(GetSpeakersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSpeakersHandlerFunc) Handle(params GetSpeakersParams) middleware.Responder {
	return fn(params)
}

// GetSpeakersHandler interface for that can handle valid get speakers params
type GetSpeakersHandler interface {
	Handle(GetSpeakersParams) middleware.Responder
}

// NewGetSpeakers creates a new http.Handler for the get speakers operation
func NewGetSpeakers(ctx *middleware.Context, handler GetSpeakersHandler) *GetSpeakers {
	return &GetSpeakers{Context: ctx, Handler: handler}
}

/* GetSpeakers swagger:route GET /meetups/{id}/speakers speaker getSpeakers

Returns a list of speakers in meetup.

A JSON array of speakers

*/
type GetSpeakers struct {
	Context *middleware.Context
	Handler GetSpeakersHandler
}

func (o *GetSpeakers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSpeakersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
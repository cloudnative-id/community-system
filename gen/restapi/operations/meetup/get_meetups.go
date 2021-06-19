// Code generated by go-swagger; DO NOT EDIT.

package meetup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMeetupsHandlerFunc turns a function with the right signature into a get meetups handler
type GetMeetupsHandlerFunc func(GetMeetupsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMeetupsHandlerFunc) Handle(params GetMeetupsParams) middleware.Responder {
	return fn(params)
}

// GetMeetupsHandler interface for that can handle valid get meetups params
type GetMeetupsHandler interface {
	Handle(GetMeetupsParams) middleware.Responder
}

// NewGetMeetups creates a new http.Handler for the get meetups operation
func NewGetMeetups(ctx *middleware.Context, handler GetMeetupsHandler) *GetMeetups {
	return &GetMeetups{Context: ctx, Handler: handler}
}

/* GetMeetups swagger:route GET /meetups meetup getMeetups

Returns a list of meetups.

A JSON array of meetups

*/
type GetMeetups struct {
	Context *middleware.Context
	Handler GetMeetupsHandler
}

func (o *GetMeetups) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMeetupsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
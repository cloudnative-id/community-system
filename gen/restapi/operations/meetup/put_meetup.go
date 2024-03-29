// Code generated by go-swagger; DO NOT EDIT.

package meetup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutMeetupHandlerFunc turns a function with the right signature into a put meetup handler
type PutMeetupHandlerFunc func(PutMeetupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutMeetupHandlerFunc) Handle(params PutMeetupParams) middleware.Responder {
	return fn(params)
}

// PutMeetupHandler interface for that can handle valid put meetup params
type PutMeetupHandler interface {
	Handle(PutMeetupParams) middleware.Responder
}

// NewPutMeetup creates a new http.Handler for the put meetup operation
func NewPutMeetup(ctx *middleware.Context, handler PutMeetupHandler) *PutMeetup {
	return &PutMeetup{Context: ctx, Handler: handler}
}

/* PutMeetup swagger:route PUT /meetups meetup putMeetup

Put meetup data.

create meetup

*/
type PutMeetup struct {
	Context *middleware.Context
	Handler PutMeetupHandler
}

func (o *PutMeetup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutMeetupParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

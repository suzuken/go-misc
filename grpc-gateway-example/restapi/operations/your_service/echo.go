package your_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// EchoHandlerFunc turns a function with the right signature into a echo handler
type EchoHandlerFunc func(EchoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn EchoHandlerFunc) Handle(params EchoParams) middleware.Responder {
	return fn(params)
}

// EchoHandler interface for that can handle valid echo params
type EchoHandler interface {
	Handle(EchoParams) middleware.Responder
}

// NewEcho creates a new http.Handler for the echo operation
func NewEcho(ctx *middleware.Context, handler EchoHandler) *Echo {
	return &Echo{Context: ctx, Handler: handler}
}

/*Echo swagger:route POST /v1/example/echo YourService echo

Echo echo API

*/
type Echo struct {
	Context *middleware.Context
	Handler EchoHandler
}

func (o *Echo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewEchoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

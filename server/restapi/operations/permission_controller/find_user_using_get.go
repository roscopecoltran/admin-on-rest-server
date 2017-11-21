// Code generated by go-swagger; DO NOT EDIT.

package permission_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindUserUsingGETHandlerFunc turns a function with the right signature into a find user using g e t handler
type FindUserUsingGETHandlerFunc func(FindUserUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindUserUsingGETHandlerFunc) Handle(params FindUserUsingGETParams) middleware.Responder {
	return fn(params)
}

// FindUserUsingGETHandler interface for that can handle valid find user using g e t params
type FindUserUsingGETHandler interface {
	Handle(FindUserUsingGETParams) middleware.Responder
}

// NewFindUserUsingGET creates a new http.Handler for the find user using g e t operation
func NewFindUserUsingGET(ctx *middleware.Context, handler FindUserUsingGETHandler) *FindUserUsingGET {
	return &FindUserUsingGET{Context: ctx, Handler: handler}
}

/*FindUserUsingGET swagger:route GET /permission/_permission/{id} permission-controller findUserUsingGET

findUser

*/
type FindUserUsingGET struct {
	Context *middleware.Context
	Handler FindUserUsingGETHandler
}

func (o *FindUserUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindUserUsingGETParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

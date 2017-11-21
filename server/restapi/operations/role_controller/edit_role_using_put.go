// Code generated by go-swagger; DO NOT EDIT.

package role_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// EditRoleUsingPUTHandlerFunc turns a function with the right signature into a edit role using p u t handler
type EditRoleUsingPUTHandlerFunc func(EditRoleUsingPUTParams) middleware.Responder

// Handle executing the request and returning a response
func (fn EditRoleUsingPUTHandlerFunc) Handle(params EditRoleUsingPUTParams) middleware.Responder {
	return fn(params)
}

// EditRoleUsingPUTHandler interface for that can handle valid edit role using p u t params
type EditRoleUsingPUTHandler interface {
	Handle(EditRoleUsingPUTParams) middleware.Responder
}

// NewEditRoleUsingPUT creates a new http.Handler for the edit role using p u t operation
func NewEditRoleUsingPUT(ctx *middleware.Context, handler EditRoleUsingPUTHandler) *EditRoleUsingPUT {
	return &EditRoleUsingPUT{Context: ctx, Handler: handler}
}

/*EditRoleUsingPUT swagger:route PUT /role/_roles/put/{id} role-controller editRoleUsingPUT

editRole

*/
type EditRoleUsingPUT struct {
	Context *middleware.Context
	Handler EditRoleUsingPUTHandler
}

func (o *EditRoleUsingPUT) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewEditRoleUsingPUTParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

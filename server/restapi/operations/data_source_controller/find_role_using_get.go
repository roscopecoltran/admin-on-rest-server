// Code generated by go-swagger; DO NOT EDIT.

package data_source_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindRoleUsingGETHandlerFunc turns a function with the right signature into a find role using g e t handler
type FindRoleUsingGETHandlerFunc func(FindRoleUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindRoleUsingGETHandlerFunc) Handle(params FindRoleUsingGETParams) middleware.Responder {
	return fn(params)
}

// FindRoleUsingGETHandler interface for that can handle valid find role using g e t params
type FindRoleUsingGETHandler interface {
	Handle(FindRoleUsingGETParams) middleware.Responder
}

// NewFindRoleUsingGET creates a new http.Handler for the find role using g e t operation
func NewFindRoleUsingGET(ctx *middleware.Context, handler FindRoleUsingGETHandler) *FindRoleUsingGET {
	return &FindRoleUsingGET{Context: ctx, Handler: handler}
}

/*FindRoleUsingGET swagger:route GET /datasource/_datasource/{datasourceId} data-source-controller findRoleUsingGET

findRole

*/
type FindRoleUsingGET struct {
	Context *middleware.Context
	Handler FindRoleUsingGETHandler
}

func (o *FindRoleUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindRoleUsingGETParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

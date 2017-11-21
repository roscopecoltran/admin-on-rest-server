// Code generated by go-swagger; DO NOT EDIT.

package schema_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindAllFieldsUsingGETHandlerFunc turns a function with the right signature into a find all fields using g e t handler
type FindAllFieldsUsingGETHandlerFunc func(FindAllFieldsUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindAllFieldsUsingGETHandlerFunc) Handle(params FindAllFieldsUsingGETParams) middleware.Responder {
	return fn(params)
}

// FindAllFieldsUsingGETHandler interface for that can handle valid find all fields using g e t params
type FindAllFieldsUsingGETHandler interface {
	Handle(FindAllFieldsUsingGETParams) middleware.Responder
}

// NewFindAllFieldsUsingGET creates a new http.Handler for the find all fields using g e t operation
func NewFindAllFieldsUsingGET(ctx *middleware.Context, handler FindAllFieldsUsingGETHandler) *FindAllFieldsUsingGET {
	return &FindAllFieldsUsingGET{Context: ctx, Handler: handler}
}

/*FindAllFieldsUsingGET swagger:route GET /schemas/_fields schema-controller findAllFieldsUsingGET

findAllFields

*/
type FindAllFieldsUsingGET struct {
	Context *middleware.Context
	Handler FindAllFieldsUsingGETHandler
}

func (o *FindAllFieldsUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindAllFieldsUsingGETParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

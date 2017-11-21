// Code generated by go-swagger; DO NOT EDIT.

package authentication_rest_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RefreshAndGetAuthenticationTokenUsingGETHandlerFunc turns a function with the right signature into a refresh and get authentication token using g e t handler
type RefreshAndGetAuthenticationTokenUsingGETHandlerFunc func(RefreshAndGetAuthenticationTokenUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RefreshAndGetAuthenticationTokenUsingGETHandlerFunc) Handle(params RefreshAndGetAuthenticationTokenUsingGETParams) middleware.Responder {
	return fn(params)
}

// RefreshAndGetAuthenticationTokenUsingGETHandler interface for that can handle valid refresh and get authentication token using g e t params
type RefreshAndGetAuthenticationTokenUsingGETHandler interface {
	Handle(RefreshAndGetAuthenticationTokenUsingGETParams) middleware.Responder
}

// NewRefreshAndGetAuthenticationTokenUsingGET creates a new http.Handler for the refresh and get authentication token using g e t operation
func NewRefreshAndGetAuthenticationTokenUsingGET(ctx *middleware.Context, handler RefreshAndGetAuthenticationTokenUsingGETHandler) *RefreshAndGetAuthenticationTokenUsingGET {
	return &RefreshAndGetAuthenticationTokenUsingGET{Context: ctx, Handler: handler}
}

/*RefreshAndGetAuthenticationTokenUsingGET swagger:route GET /refresh authentication-rest-controller refreshAndGetAuthenticationTokenUsingGET

refreshAndGetAuthenticationToken

*/
type RefreshAndGetAuthenticationTokenUsingGET struct {
	Context *middleware.Context
	Handler RefreshAndGetAuthenticationTokenUsingGETHandler
}

func (o *RefreshAndGetAuthenticationTokenUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRefreshAndGetAuthenticationTokenUsingGETParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RefreshAndGetAuthenticationTokenUsingGETOKBody refresh and get authentication token using g e t o k body
// swagger:model RefreshAndGetAuthenticationTokenUsingGETOKBody
type RefreshAndGetAuthenticationTokenUsingGETOKBody interface{}

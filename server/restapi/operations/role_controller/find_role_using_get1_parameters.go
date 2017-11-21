// Code generated by go-swagger; DO NOT EDIT.

package role_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindRoleUsingGET1Params creates a new FindRoleUsingGET1Params object
// with the default values initialized.
func NewFindRoleUsingGET1Params() FindRoleUsingGET1Params {
	var ()
	return FindRoleUsingGET1Params{}
}

// FindRoleUsingGET1Params contains all the bound params for the find role using g e t 1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters findRoleUsingGET_1
type FindRoleUsingGET1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*roleId
	  Required: true
	  In: path
	*/
	RoleID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindRoleUsingGET1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rRoleID, rhkRoleID, _ := route.Params.GetOK("roleId")
	if err := o.bindRoleID(rRoleID, rhkRoleID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindRoleUsingGET1Params) bindRoleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.RoleID = raw

	return nil
}
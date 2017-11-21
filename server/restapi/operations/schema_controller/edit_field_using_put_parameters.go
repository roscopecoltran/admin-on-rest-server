// Code generated by go-swagger; DO NOT EDIT.

package schema_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/roscopecoltran/admin-on-rest-server/server/models"
)

// NewEditFieldUsingPUTParams creates a new EditFieldUsingPUTParams object
// with the default values initialized.
func NewEditFieldUsingPUTParams() EditFieldUsingPUTParams {
	var ()
	return EditFieldUsingPUTParams{}
}

// EditFieldUsingPUTParams contains all the bound params for the edit field using p u t operation
// typically these are obtained from a http.Request
//
// swagger:parameters editFieldUsingPUT
type EditFieldUsingPUTParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*field
	  Required: true
	  In: body
	*/
	Field *models.Field
	/*id
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *EditFieldUsingPUTParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Field
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("field", "body"))
			} else {
				res = append(res, errors.NewParseError("field", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Field = &body
			}
		}

	} else {
		res = append(res, errors.Required("field", "body"))
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditFieldUsingPUTParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}

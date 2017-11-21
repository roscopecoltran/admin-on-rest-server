// Code generated by go-swagger; DO NOT EDIT.

package data_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDataMutationUsingPUTParams creates a new DataMutationUsingPUTParams object
// with the default values initialized.
func NewDataMutationUsingPUTParams() DataMutationUsingPUTParams {
	var ()
	return DataMutationUsingPUTParams{}
}

// DataMutationUsingPUTParams contains all the bound params for the data mutation using p u t operation
// typically these are obtained from a http.Request
//
// swagger:parameters dataMutationUsingPUT
type DataMutationUsingPUTParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*allRequestParams
	  In: body
	*/
	AllRequestParams interface{}
	/*entity
	  Required: true
	  In: path
	*/
	Entity string
	/*id
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DataMutationUsingPUTParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body interface{}
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("allRequestParams", "body", "", err))
		} else {

			if len(res) == 0 {
				o.AllRequestParams = body
			}
		}

	}

	rEntity, rhkEntity, _ := route.Params.GetOK("entity")
	if err := o.bindEntity(rEntity, rhkEntity, route.Formats); err != nil {
		res = append(res, err)
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

func (o *DataMutationUsingPUTParams) bindEntity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Entity = raw

	return nil
}

func (o *DataMutationUsingPUTParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}

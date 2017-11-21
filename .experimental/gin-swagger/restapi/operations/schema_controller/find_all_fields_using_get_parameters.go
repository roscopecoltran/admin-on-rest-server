package schema_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
	"github.com/mikkeloscar/gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"
)

// BusinessLogicFindAllFieldsUsingGET executes the core logic of the related
// route endpoint.
func BusinessLogicFindAllFieldsUsingGET(f func(ctx *gin.Context, params *FindAllFieldsUsingGETParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &FindAllFieldsUsingGETParams{}
		err := params.bindRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}
			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := f(ctx, params)
		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// FindAllFieldsUsingGETParams contains all the bound params for the find all fields using g e t operation
// typically these are obtained from a http.Request
//
// swagger:parameters findAllFieldsUsingGET
type FindAllFieldsUsingGETParams struct {

	/*eid
	  Required: true
	  In: query
	*/
	Eid string
}

// FindAllFieldsUsingGETParamsFromCtx gets the params struct from the gin context.
func FindAllFieldsUsingGETParamsFromCtx(ctx *gin.Context) *FindAllFieldsUsingGETParams {
	params, _ := ctx.Get("params")
	return params.(*FindAllFieldsUsingGETParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindAllFieldsUsingGETParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	qs := runtime.Values(ctx.Request.URL.Query())

	qEid, qhkEid, _ := qs.GetOK("eid")
	if err := o.bindEid(qEid, qhkEid, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindAllFieldsUsingGETParams) bindEid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("eid", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("eid", "query", raw); err != nil {
		return err
	}

	o.Eid = raw

	return nil
}

// vim: ft=go

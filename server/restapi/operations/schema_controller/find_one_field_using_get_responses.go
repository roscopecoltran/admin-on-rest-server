// Code generated by go-swagger; DO NOT EDIT.

package schema_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/roscopecoltran/admin-on-rest-server/server/models"
)

// FindOneFieldUsingGETOKCode is the HTTP code returned for type FindOneFieldUsingGETOK
const FindOneFieldUsingGETOKCode int = 200

/*FindOneFieldUsingGETOK OK

swagger:response findOneFieldUsingGETOK
*/
type FindOneFieldUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.Field `json:"body,omitempty"`
}

// NewFindOneFieldUsingGETOK creates FindOneFieldUsingGETOK with default headers values
func NewFindOneFieldUsingGETOK() *FindOneFieldUsingGETOK {
	return &FindOneFieldUsingGETOK{}
}

// WithPayload adds the payload to the find one field using g e t o k response
func (o *FindOneFieldUsingGETOK) WithPayload(payload *models.Field) *FindOneFieldUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find one field using g e t o k response
func (o *FindOneFieldUsingGETOK) SetPayload(payload *models.Field) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOneFieldUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindOneFieldUsingGETUnauthorizedCode is the HTTP code returned for type FindOneFieldUsingGETUnauthorized
const FindOneFieldUsingGETUnauthorizedCode int = 401

/*FindOneFieldUsingGETUnauthorized Unauthorized

swagger:response findOneFieldUsingGETUnauthorized
*/
type FindOneFieldUsingGETUnauthorized struct {
}

// NewFindOneFieldUsingGETUnauthorized creates FindOneFieldUsingGETUnauthorized with default headers values
func NewFindOneFieldUsingGETUnauthorized() *FindOneFieldUsingGETUnauthorized {
	return &FindOneFieldUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *FindOneFieldUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// FindOneFieldUsingGETForbiddenCode is the HTTP code returned for type FindOneFieldUsingGETForbidden
const FindOneFieldUsingGETForbiddenCode int = 403

/*FindOneFieldUsingGETForbidden Forbidden

swagger:response findOneFieldUsingGETForbidden
*/
type FindOneFieldUsingGETForbidden struct {
}

// NewFindOneFieldUsingGETForbidden creates FindOneFieldUsingGETForbidden with default headers values
func NewFindOneFieldUsingGETForbidden() *FindOneFieldUsingGETForbidden {
	return &FindOneFieldUsingGETForbidden{}
}

// WriteResponse to the client
func (o *FindOneFieldUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// FindOneFieldUsingGETNotFoundCode is the HTTP code returned for type FindOneFieldUsingGETNotFound
const FindOneFieldUsingGETNotFoundCode int = 404

/*FindOneFieldUsingGETNotFound Not Found

swagger:response findOneFieldUsingGETNotFound
*/
type FindOneFieldUsingGETNotFound struct {
}

// NewFindOneFieldUsingGETNotFound creates FindOneFieldUsingGETNotFound with default headers values
func NewFindOneFieldUsingGETNotFound() *FindOneFieldUsingGETNotFound {
	return &FindOneFieldUsingGETNotFound{}
}

// WriteResponse to the client
func (o *FindOneFieldUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
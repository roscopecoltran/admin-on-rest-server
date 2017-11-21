// Code generated by go-swagger; DO NOT EDIT.

package data_source_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/roscopecoltran/admin-on-rest-server/server/models"
)

// ListUsingGETOKCode is the HTTP code returned for type ListUsingGETOK
const ListUsingGETOKCode int = 200

/*ListUsingGETOK OK

swagger:response listUsingGETOK
*/
type ListUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload []*models.DataSource `json:"body,omitempty"`
}

// NewListUsingGETOK creates ListUsingGETOK with default headers values
func NewListUsingGETOK() *ListUsingGETOK {
	return &ListUsingGETOK{}
}

// WithPayload adds the payload to the list using g e t o k response
func (o *ListUsingGETOK) WithPayload(payload []*models.DataSource) *ListUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list using g e t o k response
func (o *ListUsingGETOK) SetPayload(payload []*models.DataSource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.DataSource, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ListUsingGETUnauthorizedCode is the HTTP code returned for type ListUsingGETUnauthorized
const ListUsingGETUnauthorizedCode int = 401

/*ListUsingGETUnauthorized Unauthorized

swagger:response listUsingGETUnauthorized
*/
type ListUsingGETUnauthorized struct {
}

// NewListUsingGETUnauthorized creates ListUsingGETUnauthorized with default headers values
func NewListUsingGETUnauthorized() *ListUsingGETUnauthorized {
	return &ListUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *ListUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ListUsingGETForbiddenCode is the HTTP code returned for type ListUsingGETForbidden
const ListUsingGETForbiddenCode int = 403

/*ListUsingGETForbidden Forbidden

swagger:response listUsingGETForbidden
*/
type ListUsingGETForbidden struct {
}

// NewListUsingGETForbidden creates ListUsingGETForbidden with default headers values
func NewListUsingGETForbidden() *ListUsingGETForbidden {
	return &ListUsingGETForbidden{}
}

// WriteResponse to the client
func (o *ListUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ListUsingGETNotFoundCode is the HTTP code returned for type ListUsingGETNotFound
const ListUsingGETNotFoundCode int = 404

/*ListUsingGETNotFound Not Found

swagger:response listUsingGETNotFound
*/
type ListUsingGETNotFound struct {
}

// NewListUsingGETNotFound creates ListUsingGETNotFound with default headers values
func NewListUsingGETNotFound() *ListUsingGETNotFound {
	return &ListUsingGETNotFound{}
}

// WriteResponse to the client
func (o *ListUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
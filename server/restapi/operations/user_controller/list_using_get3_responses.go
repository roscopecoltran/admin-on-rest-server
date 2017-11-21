// Code generated by go-swagger; DO NOT EDIT.

package user_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/roscopecoltran/admin-on-rest-server/server/models"
)

// ListUsingGET3OKCode is the HTTP code returned for type ListUsingGET3OK
const ListUsingGET3OKCode int = 200

/*ListUsingGET3OK OK

swagger:response listUsingGET3OK
*/
type ListUsingGET3OK struct {

	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewListUsingGET3OK creates ListUsingGET3OK with default headers values
func NewListUsingGET3OK() *ListUsingGET3OK {
	return &ListUsingGET3OK{}
}

// WithPayload adds the payload to the list using g e t3 o k response
func (o *ListUsingGET3OK) WithPayload(payload []*models.User) *ListUsingGET3OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list using g e t3 o k response
func (o *ListUsingGET3OK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUsingGET3OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ListUsingGET3UnauthorizedCode is the HTTP code returned for type ListUsingGET3Unauthorized
const ListUsingGET3UnauthorizedCode int = 401

/*ListUsingGET3Unauthorized Unauthorized

swagger:response listUsingGET3Unauthorized
*/
type ListUsingGET3Unauthorized struct {
}

// NewListUsingGET3Unauthorized creates ListUsingGET3Unauthorized with default headers values
func NewListUsingGET3Unauthorized() *ListUsingGET3Unauthorized {
	return &ListUsingGET3Unauthorized{}
}

// WriteResponse to the client
func (o *ListUsingGET3Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ListUsingGET3ForbiddenCode is the HTTP code returned for type ListUsingGET3Forbidden
const ListUsingGET3ForbiddenCode int = 403

/*ListUsingGET3Forbidden Forbidden

swagger:response listUsingGET3Forbidden
*/
type ListUsingGET3Forbidden struct {
}

// NewListUsingGET3Forbidden creates ListUsingGET3Forbidden with default headers values
func NewListUsingGET3Forbidden() *ListUsingGET3Forbidden {
	return &ListUsingGET3Forbidden{}
}

// WriteResponse to the client
func (o *ListUsingGET3Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ListUsingGET3NotFoundCode is the HTTP code returned for type ListUsingGET3NotFound
const ListUsingGET3NotFoundCode int = 404

/*ListUsingGET3NotFound Not Found

swagger:response listUsingGET3NotFound
*/
type ListUsingGET3NotFound struct {
}

// NewListUsingGET3NotFound creates ListUsingGET3NotFound with default headers values
func NewListUsingGET3NotFound() *ListUsingGET3NotFound {
	return &ListUsingGET3NotFound{}
}

// WriteResponse to the client
func (o *ListUsingGET3NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

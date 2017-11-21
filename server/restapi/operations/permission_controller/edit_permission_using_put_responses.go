// Code generated by go-swagger; DO NOT EDIT.

package permission_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/roscopecoltran/admin-on-rest-server/server/models"
)

// EditPermissionUsingPUTOKCode is the HTTP code returned for type EditPermissionUsingPUTOK
const EditPermissionUsingPUTOKCode int = 200

/*EditPermissionUsingPUTOK OK

swagger:response editPermissionUsingPUTOK
*/
type EditPermissionUsingPUTOK struct {

	/*
	  In: Body
	*/
	Payload *models.Permission `json:"body,omitempty"`
}

// NewEditPermissionUsingPUTOK creates EditPermissionUsingPUTOK with default headers values
func NewEditPermissionUsingPUTOK() *EditPermissionUsingPUTOK {
	return &EditPermissionUsingPUTOK{}
}

// WithPayload adds the payload to the edit permission using p u t o k response
func (o *EditPermissionUsingPUTOK) WithPayload(payload *models.Permission) *EditPermissionUsingPUTOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the edit permission using p u t o k response
func (o *EditPermissionUsingPUTOK) SetPayload(payload *models.Permission) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EditPermissionUsingPUTOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EditPermissionUsingPUTCreatedCode is the HTTP code returned for type EditPermissionUsingPUTCreated
const EditPermissionUsingPUTCreatedCode int = 201

/*EditPermissionUsingPUTCreated Created

swagger:response editPermissionUsingPUTCreated
*/
type EditPermissionUsingPUTCreated struct {
}

// NewEditPermissionUsingPUTCreated creates EditPermissionUsingPUTCreated with default headers values
func NewEditPermissionUsingPUTCreated() *EditPermissionUsingPUTCreated {
	return &EditPermissionUsingPUTCreated{}
}

// WriteResponse to the client
func (o *EditPermissionUsingPUTCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// EditPermissionUsingPUTUnauthorizedCode is the HTTP code returned for type EditPermissionUsingPUTUnauthorized
const EditPermissionUsingPUTUnauthorizedCode int = 401

/*EditPermissionUsingPUTUnauthorized Unauthorized

swagger:response editPermissionUsingPUTUnauthorized
*/
type EditPermissionUsingPUTUnauthorized struct {
}

// NewEditPermissionUsingPUTUnauthorized creates EditPermissionUsingPUTUnauthorized with default headers values
func NewEditPermissionUsingPUTUnauthorized() *EditPermissionUsingPUTUnauthorized {
	return &EditPermissionUsingPUTUnauthorized{}
}

// WriteResponse to the client
func (o *EditPermissionUsingPUTUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// EditPermissionUsingPUTForbiddenCode is the HTTP code returned for type EditPermissionUsingPUTForbidden
const EditPermissionUsingPUTForbiddenCode int = 403

/*EditPermissionUsingPUTForbidden Forbidden

swagger:response editPermissionUsingPUTForbidden
*/
type EditPermissionUsingPUTForbidden struct {
}

// NewEditPermissionUsingPUTForbidden creates EditPermissionUsingPUTForbidden with default headers values
func NewEditPermissionUsingPUTForbidden() *EditPermissionUsingPUTForbidden {
	return &EditPermissionUsingPUTForbidden{}
}

// WriteResponse to the client
func (o *EditPermissionUsingPUTForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// EditPermissionUsingPUTNotFoundCode is the HTTP code returned for type EditPermissionUsingPUTNotFound
const EditPermissionUsingPUTNotFoundCode int = 404

/*EditPermissionUsingPUTNotFound Not Found

swagger:response editPermissionUsingPUTNotFound
*/
type EditPermissionUsingPUTNotFound struct {
}

// NewEditPermissionUsingPUTNotFound creates EditPermissionUsingPUTNotFound with default headers values
func NewEditPermissionUsingPUTNotFound() *EditPermissionUsingPUTNotFound {
	return &EditPermissionUsingPUTNotFound{}
}

// WriteResponse to the client
func (o *EditPermissionUsingPUTNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

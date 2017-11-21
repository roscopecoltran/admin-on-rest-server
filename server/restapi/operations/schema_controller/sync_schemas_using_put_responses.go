// Code generated by go-swagger; DO NOT EDIT.

package schema_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/roscopecoltran/admin-on-rest-server/server/models"
)

// SyncSchemasUsingPUTOKCode is the HTTP code returned for type SyncSchemasUsingPUTOK
const SyncSchemasUsingPUTOKCode int = 200

/*SyncSchemasUsingPUTOK OK

swagger:response syncSchemasUsingPUTOK
*/
type SyncSchemasUsingPUTOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseEntity `json:"body,omitempty"`
}

// NewSyncSchemasUsingPUTOK creates SyncSchemasUsingPUTOK with default headers values
func NewSyncSchemasUsingPUTOK() *SyncSchemasUsingPUTOK {
	return &SyncSchemasUsingPUTOK{}
}

// WithPayload adds the payload to the sync schemas using p u t o k response
func (o *SyncSchemasUsingPUTOK) WithPayload(payload *models.ResponseEntity) *SyncSchemasUsingPUTOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sync schemas using p u t o k response
func (o *SyncSchemasUsingPUTOK) SetPayload(payload *models.ResponseEntity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SyncSchemasUsingPUTOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SyncSchemasUsingPUTCreatedCode is the HTTP code returned for type SyncSchemasUsingPUTCreated
const SyncSchemasUsingPUTCreatedCode int = 201

/*SyncSchemasUsingPUTCreated Created

swagger:response syncSchemasUsingPUTCreated
*/
type SyncSchemasUsingPUTCreated struct {
}

// NewSyncSchemasUsingPUTCreated creates SyncSchemasUsingPUTCreated with default headers values
func NewSyncSchemasUsingPUTCreated() *SyncSchemasUsingPUTCreated {
	return &SyncSchemasUsingPUTCreated{}
}

// WriteResponse to the client
func (o *SyncSchemasUsingPUTCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// SyncSchemasUsingPUTUnauthorizedCode is the HTTP code returned for type SyncSchemasUsingPUTUnauthorized
const SyncSchemasUsingPUTUnauthorizedCode int = 401

/*SyncSchemasUsingPUTUnauthorized Unauthorized

swagger:response syncSchemasUsingPUTUnauthorized
*/
type SyncSchemasUsingPUTUnauthorized struct {
}

// NewSyncSchemasUsingPUTUnauthorized creates SyncSchemasUsingPUTUnauthorized with default headers values
func NewSyncSchemasUsingPUTUnauthorized() *SyncSchemasUsingPUTUnauthorized {
	return &SyncSchemasUsingPUTUnauthorized{}
}

// WriteResponse to the client
func (o *SyncSchemasUsingPUTUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SyncSchemasUsingPUTForbiddenCode is the HTTP code returned for type SyncSchemasUsingPUTForbidden
const SyncSchemasUsingPUTForbiddenCode int = 403

/*SyncSchemasUsingPUTForbidden Forbidden

swagger:response syncSchemasUsingPUTForbidden
*/
type SyncSchemasUsingPUTForbidden struct {
}

// NewSyncSchemasUsingPUTForbidden creates SyncSchemasUsingPUTForbidden with default headers values
func NewSyncSchemasUsingPUTForbidden() *SyncSchemasUsingPUTForbidden {
	return &SyncSchemasUsingPUTForbidden{}
}

// WriteResponse to the client
func (o *SyncSchemasUsingPUTForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// SyncSchemasUsingPUTNotFoundCode is the HTTP code returned for type SyncSchemasUsingPUTNotFound
const SyncSchemasUsingPUTNotFoundCode int = 404

/*SyncSchemasUsingPUTNotFound Not Found

swagger:response syncSchemasUsingPUTNotFound
*/
type SyncSchemasUsingPUTNotFound struct {
}

// NewSyncSchemasUsingPUTNotFound creates SyncSchemasUsingPUTNotFound with default headers values
func NewSyncSchemasUsingPUTNotFound() *SyncSchemasUsingPUTNotFound {
	return &SyncSchemasUsingPUTNotFound{}
}

// WriteResponse to the client
func (o *SyncSchemasUsingPUTNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
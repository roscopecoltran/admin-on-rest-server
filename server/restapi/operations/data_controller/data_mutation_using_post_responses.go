// Code generated by go-swagger; DO NOT EDIT.

package data_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DataMutationUsingPOSTOKCode is the HTTP code returned for type DataMutationUsingPOSTOK
const DataMutationUsingPOSTOKCode int = 200

/*DataMutationUsingPOSTOK OK

swagger:response dataMutationUsingPOSTOK
*/
type DataMutationUsingPOSTOK struct {

	/*
	  In: Body
	*/
	Payload DataMutationUsingPOSTOKBody `json:"body,omitempty"`
}

// NewDataMutationUsingPOSTOK creates DataMutationUsingPOSTOK with default headers values
func NewDataMutationUsingPOSTOK() *DataMutationUsingPOSTOK {
	return &DataMutationUsingPOSTOK{}
}

// WithPayload adds the payload to the data mutation using p o s t o k response
func (o *DataMutationUsingPOSTOK) WithPayload(payload DataMutationUsingPOSTOKBody) *DataMutationUsingPOSTOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the data mutation using p o s t o k response
func (o *DataMutationUsingPOSTOK) SetPayload(payload DataMutationUsingPOSTOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DataMutationUsingPOSTOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// DataMutationUsingPOSTCreatedCode is the HTTP code returned for type DataMutationUsingPOSTCreated
const DataMutationUsingPOSTCreatedCode int = 201

/*DataMutationUsingPOSTCreated Created

swagger:response dataMutationUsingPOSTCreated
*/
type DataMutationUsingPOSTCreated struct {
}

// NewDataMutationUsingPOSTCreated creates DataMutationUsingPOSTCreated with default headers values
func NewDataMutationUsingPOSTCreated() *DataMutationUsingPOSTCreated {
	return &DataMutationUsingPOSTCreated{}
}

// WriteResponse to the client
func (o *DataMutationUsingPOSTCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// DataMutationUsingPOSTUnauthorizedCode is the HTTP code returned for type DataMutationUsingPOSTUnauthorized
const DataMutationUsingPOSTUnauthorizedCode int = 401

/*DataMutationUsingPOSTUnauthorized Unauthorized

swagger:response dataMutationUsingPOSTUnauthorized
*/
type DataMutationUsingPOSTUnauthorized struct {
}

// NewDataMutationUsingPOSTUnauthorized creates DataMutationUsingPOSTUnauthorized with default headers values
func NewDataMutationUsingPOSTUnauthorized() *DataMutationUsingPOSTUnauthorized {
	return &DataMutationUsingPOSTUnauthorized{}
}

// WriteResponse to the client
func (o *DataMutationUsingPOSTUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DataMutationUsingPOSTForbiddenCode is the HTTP code returned for type DataMutationUsingPOSTForbidden
const DataMutationUsingPOSTForbiddenCode int = 403

/*DataMutationUsingPOSTForbidden Forbidden

swagger:response dataMutationUsingPOSTForbidden
*/
type DataMutationUsingPOSTForbidden struct {
}

// NewDataMutationUsingPOSTForbidden creates DataMutationUsingPOSTForbidden with default headers values
func NewDataMutationUsingPOSTForbidden() *DataMutationUsingPOSTForbidden {
	return &DataMutationUsingPOSTForbidden{}
}

// WriteResponse to the client
func (o *DataMutationUsingPOSTForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DataMutationUsingPOSTNotFoundCode is the HTTP code returned for type DataMutationUsingPOSTNotFound
const DataMutationUsingPOSTNotFoundCode int = 404

/*DataMutationUsingPOSTNotFound Not Found

swagger:response dataMutationUsingPOSTNotFound
*/
type DataMutationUsingPOSTNotFound struct {
}

// NewDataMutationUsingPOSTNotFound creates DataMutationUsingPOSTNotFound with default headers values
func NewDataMutationUsingPOSTNotFound() *DataMutationUsingPOSTNotFound {
	return &DataMutationUsingPOSTNotFound{}
}

// WriteResponse to the client
func (o *DataMutationUsingPOSTNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// PostEnrollmentsOKCode is the HTTP code returned for type PostEnrollmentsOK
const PostEnrollmentsOKCode int = 200

/*PostEnrollmentsOK OK

swagger:response postEnrollmentsOK
*/
type PostEnrollmentsOK struct {
}

// NewPostEnrollmentsOK creates PostEnrollmentsOK with default headers values
func NewPostEnrollmentsOK() *PostEnrollmentsOK {

	return &PostEnrollmentsOK{}
}

// WriteResponse to the client
func (o *PostEnrollmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostEnrollmentsBadRequestCode is the HTTP code returned for type PostEnrollmentsBadRequest
const PostEnrollmentsBadRequestCode int = 400

/*PostEnrollmentsBadRequest Invalid input

swagger:response postEnrollmentsBadRequest
*/
type PostEnrollmentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostEnrollmentsBadRequest creates PostEnrollmentsBadRequest with default headers values
func NewPostEnrollmentsBadRequest() *PostEnrollmentsBadRequest {

	return &PostEnrollmentsBadRequest{}
}

// WithPayload adds the payload to the post enrollments bad request response
func (o *PostEnrollmentsBadRequest) WithPayload(payload *models.Error) *PostEnrollmentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post enrollments bad request response
func (o *PostEnrollmentsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostEnrollmentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostEnrollmentsUnauthorizedCode is the HTTP code returned for type PostEnrollmentsUnauthorized
const PostEnrollmentsUnauthorizedCode int = 401

/*PostEnrollmentsUnauthorized Unauthorized

swagger:response postEnrollmentsUnauthorized
*/
type PostEnrollmentsUnauthorized struct {
}

// NewPostEnrollmentsUnauthorized creates PostEnrollmentsUnauthorized with default headers values
func NewPostEnrollmentsUnauthorized() *PostEnrollmentsUnauthorized {

	return &PostEnrollmentsUnauthorized{}
}

// WriteResponse to the client
func (o *PostEnrollmentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

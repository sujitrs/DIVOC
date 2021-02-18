// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// VerifyOTPOKCode is the HTTP code returned for type VerifyOTPOK
const VerifyOTPOKCode int = 200

/*VerifyOTPOK OK

swagger:response verifyOTPOK
*/
type VerifyOTPOK struct {

	/*
	  In: Body
	*/
	Payload *VerifyOTPOKBody `json:"body,omitempty"`
}

// NewVerifyOTPOK creates VerifyOTPOK with default headers values
func NewVerifyOTPOK() *VerifyOTPOK {

	return &VerifyOTPOK{}
}

// WithPayload adds the payload to the verify o t p o k response
func (o *VerifyOTPOK) WithPayload(payload *VerifyOTPOKBody) *VerifyOTPOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify o t p o k response
func (o *VerifyOTPOK) SetPayload(payload *VerifyOTPOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyOTPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VerifyOTPBadRequestCode is the HTTP code returned for type VerifyOTPBadRequest
const VerifyOTPBadRequestCode int = 400

/*VerifyOTPBadRequest Bad request

swagger:response verifyOTPBadRequest
*/
type VerifyOTPBadRequest struct {
}

// NewVerifyOTPBadRequest creates VerifyOTPBadRequest with default headers values
func NewVerifyOTPBadRequest() *VerifyOTPBadRequest {

	return &VerifyOTPBadRequest{}
}

// WriteResponse to the client
func (o *VerifyOTPBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// VerifyOTPUnauthorizedCode is the HTTP code returned for type VerifyOTPUnauthorized
const VerifyOTPUnauthorizedCode int = 401

/*VerifyOTPUnauthorized Invalid OTP

swagger:response verifyOTPUnauthorized
*/
type VerifyOTPUnauthorized struct {
}

// NewVerifyOTPUnauthorized creates VerifyOTPUnauthorized with default headers values
func NewVerifyOTPUnauthorized() *VerifyOTPUnauthorized {

	return &VerifyOTPUnauthorized{}
}

// WriteResponse to the client
func (o *VerifyOTPUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// VerifyOTPTooManyRequestsCode is the HTTP code returned for type VerifyOTPTooManyRequests
const VerifyOTPTooManyRequestsCode int = 429

/*VerifyOTPTooManyRequests Verify otp attempts exceeded, generate new OTP

swagger:response verifyOTPTooManyRequests
*/
type VerifyOTPTooManyRequests struct {
}

// NewVerifyOTPTooManyRequests creates VerifyOTPTooManyRequests with default headers values
func NewVerifyOTPTooManyRequests() *VerifyOTPTooManyRequests {

	return &VerifyOTPTooManyRequests{}
}

// WriteResponse to the client
func (o *VerifyOTPTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}

// VerifyOTPInternalServerErrorCode is the HTTP code returned for type VerifyOTPInternalServerError
const VerifyOTPInternalServerErrorCode int = 500

/*VerifyOTPInternalServerError Internal error

swagger:response verifyOTPInternalServerError
*/
type VerifyOTPInternalServerError struct {
}

// NewVerifyOTPInternalServerError creates VerifyOTPInternalServerError with default headers values
func NewVerifyOTPInternalServerError() *VerifyOTPInternalServerError {

	return &VerifyOTPInternalServerError{}
}

// WriteResponse to the client
func (o *VerifyOTPInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

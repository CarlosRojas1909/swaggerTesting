// Code generated by go-swagger; DO NOT EDIT.

package findsummary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/JCbayfisher/swaggerTesting/models"
)

// PostFindsummaryOKCode is the HTTP code returned for type PostFindsummaryOK
const PostFindsummaryOKCode int = 200

/*PostFindsummaryOK Successful Request

swagger:response postFindsummaryOK
*/
type PostFindsummaryOK struct {

	/*
	  In: Body
	*/
	Payload *models.FindResult `json:"body,omitempty"`
}

// NewPostFindsummaryOK creates PostFindsummaryOK with default headers values
func NewPostFindsummaryOK() *PostFindsummaryOK {

	return &PostFindsummaryOK{}
}

// WithPayload adds the payload to the post findsummary o k response
func (o *PostFindsummaryOK) WithPayload(payload *models.FindResult) *PostFindsummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post findsummary o k response
func (o *PostFindsummaryOK) SetPayload(payload *models.FindResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFindsummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostFindsummaryBadRequestCode is the HTTP code returned for type PostFindsummaryBadRequest
const PostFindsummaryBadRequestCode int = 400

/*PostFindsummaryBadRequest Bad Request

swagger:response postFindsummaryBadRequest
*/
type PostFindsummaryBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostFindsummaryBadRequest creates PostFindsummaryBadRequest with default headers values
func NewPostFindsummaryBadRequest() *PostFindsummaryBadRequest {

	return &PostFindsummaryBadRequest{}
}

// WithPayload adds the payload to the post findsummary bad request response
func (o *PostFindsummaryBadRequest) WithPayload(payload *models.ReturnCode) *PostFindsummaryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post findsummary bad request response
func (o *PostFindsummaryBadRequest) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFindsummaryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostFindsummaryNotFoundCode is the HTTP code returned for type PostFindsummaryNotFound
const PostFindsummaryNotFoundCode int = 404

/*PostFindsummaryNotFound Not Found

swagger:response postFindsummaryNotFound
*/
type PostFindsummaryNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostFindsummaryNotFound creates PostFindsummaryNotFound with default headers values
func NewPostFindsummaryNotFound() *PostFindsummaryNotFound {

	return &PostFindsummaryNotFound{}
}

// WithPayload adds the payload to the post findsummary not found response
func (o *PostFindsummaryNotFound) WithPayload(payload *models.ReturnCode) *PostFindsummaryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post findsummary not found response
func (o *PostFindsummaryNotFound) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFindsummaryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostFindsummaryInternalServerErrorCode is the HTTP code returned for type PostFindsummaryInternalServerError
const PostFindsummaryInternalServerErrorCode int = 500

/*PostFindsummaryInternalServerError Internal Server Error

swagger:response postFindsummaryInternalServerError
*/
type PostFindsummaryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostFindsummaryInternalServerError creates PostFindsummaryInternalServerError with default headers values
func NewPostFindsummaryInternalServerError() *PostFindsummaryInternalServerError {

	return &PostFindsummaryInternalServerError{}
}

// WithPayload adds the payload to the post findsummary internal server error response
func (o *PostFindsummaryInternalServerError) WithPayload(payload *models.ReturnCode) *PostFindsummaryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post findsummary internal server error response
func (o *PostFindsummaryInternalServerError) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFindsummaryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostFindsummaryDefault Unexpected Error

swagger:response postFindsummaryDefault
*/
type PostFindsummaryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostFindsummaryDefault creates PostFindsummaryDefault with default headers values
func NewPostFindsummaryDefault(code int) *PostFindsummaryDefault {
	if code <= 0 {
		code = 500
	}

	return &PostFindsummaryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post findsummary default response
func (o *PostFindsummaryDefault) WithStatusCode(code int) *PostFindsummaryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post findsummary default response
func (o *PostFindsummaryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post findsummary default response
func (o *PostFindsummaryDefault) WithPayload(payload *models.ReturnCode) *PostFindsummaryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post findsummary default response
func (o *PostFindsummaryDefault) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFindsummaryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

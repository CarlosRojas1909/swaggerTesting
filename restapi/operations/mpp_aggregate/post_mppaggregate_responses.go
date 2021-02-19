// Code generated by go-swagger; DO NOT EDIT.

package mpp_aggregate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/JCbayfisher/swaggerTesting/models"
)

// PostMppaggregateOKCode is the HTTP code returned for type PostMppaggregateOK
const PostMppaggregateOKCode int = 200

/*PostMppaggregateOK Successful Request

swagger:response postMppaggregateOK
*/
type PostMppaggregateOK struct {

	/*
	  In: Body
	*/
	Payload *models.MppAggregateResult `json:"body,omitempty"`
}

// NewPostMppaggregateOK creates PostMppaggregateOK with default headers values
func NewPostMppaggregateOK() *PostMppaggregateOK {

	return &PostMppaggregateOK{}
}

// WithPayload adds the payload to the post mppaggregate o k response
func (o *PostMppaggregateOK) WithPayload(payload *models.MppAggregateResult) *PostMppaggregateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post mppaggregate o k response
func (o *PostMppaggregateOK) SetPayload(payload *models.MppAggregateResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMppaggregateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMppaggregateBadRequestCode is the HTTP code returned for type PostMppaggregateBadRequest
const PostMppaggregateBadRequestCode int = 400

/*PostMppaggregateBadRequest Bad Request

swagger:response postMppaggregateBadRequest
*/
type PostMppaggregateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostMppaggregateBadRequest creates PostMppaggregateBadRequest with default headers values
func NewPostMppaggregateBadRequest() *PostMppaggregateBadRequest {

	return &PostMppaggregateBadRequest{}
}

// WithPayload adds the payload to the post mppaggregate bad request response
func (o *PostMppaggregateBadRequest) WithPayload(payload *models.ReturnCode) *PostMppaggregateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post mppaggregate bad request response
func (o *PostMppaggregateBadRequest) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMppaggregateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMppaggregateNotFoundCode is the HTTP code returned for type PostMppaggregateNotFound
const PostMppaggregateNotFoundCode int = 404

/*PostMppaggregateNotFound Not Found

swagger:response postMppaggregateNotFound
*/
type PostMppaggregateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostMppaggregateNotFound creates PostMppaggregateNotFound with default headers values
func NewPostMppaggregateNotFound() *PostMppaggregateNotFound {

	return &PostMppaggregateNotFound{}
}

// WithPayload adds the payload to the post mppaggregate not found response
func (o *PostMppaggregateNotFound) WithPayload(payload *models.ReturnCode) *PostMppaggregateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post mppaggregate not found response
func (o *PostMppaggregateNotFound) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMppaggregateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMppaggregateInternalServerErrorCode is the HTTP code returned for type PostMppaggregateInternalServerError
const PostMppaggregateInternalServerErrorCode int = 500

/*PostMppaggregateInternalServerError Internal Server Error

swagger:response postMppaggregateInternalServerError
*/
type PostMppaggregateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostMppaggregateInternalServerError creates PostMppaggregateInternalServerError with default headers values
func NewPostMppaggregateInternalServerError() *PostMppaggregateInternalServerError {

	return &PostMppaggregateInternalServerError{}
}

// WithPayload adds the payload to the post mppaggregate internal server error response
func (o *PostMppaggregateInternalServerError) WithPayload(payload *models.ReturnCode) *PostMppaggregateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post mppaggregate internal server error response
func (o *PostMppaggregateInternalServerError) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMppaggregateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostMppaggregateDefault Unexpected Error

swagger:response postMppaggregateDefault
*/
type PostMppaggregateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewPostMppaggregateDefault creates PostMppaggregateDefault with default headers values
func NewPostMppaggregateDefault(code int) *PostMppaggregateDefault {
	if code <= 0 {
		code = 500
	}

	return &PostMppaggregateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post mppaggregate default response
func (o *PostMppaggregateDefault) WithStatusCode(code int) *PostMppaggregateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post mppaggregate default response
func (o *PostMppaggregateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post mppaggregate default response
func (o *PostMppaggregateDefault) WithPayload(payload *models.ReturnCode) *PostMppaggregateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post mppaggregate default response
func (o *PostMppaggregateDefault) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMppaggregateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
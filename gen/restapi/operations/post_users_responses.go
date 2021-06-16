// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shellingford330/lambda-go/gen/models"
)

// PostUsersOKCode is the HTTP code returned for type PostUsersOK
const PostUsersOKCode int = 200

/*PostUsersOK Success

swagger:response postUsersOK
*/
type PostUsersOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUsersOK creates PostUsersOK with default headers values
func NewPostUsersOK() *PostUsersOK {

	return &PostUsersOK{}
}

// WithPayload adds the payload to the post users o k response
func (o *PostUsersOK) WithPayload(payload *models.User) *PostUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users o k response
func (o *PostUsersOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUsersBadRequestCode is the HTTP code returned for type PostUsersBadRequest
const PostUsersBadRequestCode int = 400

/*PostUsersBadRequest Bad Request

swagger:response postUsersBadRequest
*/
type PostUsersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUsersBadRequest creates PostUsersBadRequest with default headers values
func NewPostUsersBadRequest() *PostUsersBadRequest {

	return &PostUsersBadRequest{}
}

// WithPayload adds the payload to the post users bad request response
func (o *PostUsersBadRequest) WithPayload(payload *models.Error) *PostUsersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users bad request response
func (o *PostUsersBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUsersInternalServerErrorCode is the HTTP code returned for type PostUsersInternalServerError
const PostUsersInternalServerErrorCode int = 500

/*PostUsersInternalServerError Internal Server Error

swagger:response postUsersInternalServerError
*/
type PostUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUsersInternalServerError creates PostUsersInternalServerError with default headers values
func NewPostUsersInternalServerError() *PostUsersInternalServerError {

	return &PostUsersInternalServerError{}
}

// WithPayload adds the payload to the post users internal server error response
func (o *PostUsersInternalServerError) WithPayload(payload *models.Error) *PostUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users internal server error response
func (o *PostUsersInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

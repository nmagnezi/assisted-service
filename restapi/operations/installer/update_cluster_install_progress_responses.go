// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// UpdateClusterInstallProgressNoContentCode is the HTTP code returned for type UpdateClusterInstallProgressNoContent
const UpdateClusterInstallProgressNoContentCode int = 204

/*UpdateClusterInstallProgressNoContent Update cluster install progress.

swagger:response updateClusterInstallProgressNoContent
*/
type UpdateClusterInstallProgressNoContent struct {
}

// NewUpdateClusterInstallProgressNoContent creates UpdateClusterInstallProgressNoContent with default headers values
func NewUpdateClusterInstallProgressNoContent() *UpdateClusterInstallProgressNoContent {

	return &UpdateClusterInstallProgressNoContent{}
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// UpdateClusterInstallProgressUnauthorizedCode is the HTTP code returned for type UpdateClusterInstallProgressUnauthorized
const UpdateClusterInstallProgressUnauthorizedCode int = 401

/*UpdateClusterInstallProgressUnauthorized Unauthorized.

swagger:response updateClusterInstallProgressUnauthorized
*/
type UpdateClusterInstallProgressUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUpdateClusterInstallProgressUnauthorized creates UpdateClusterInstallProgressUnauthorized with default headers values
func NewUpdateClusterInstallProgressUnauthorized() *UpdateClusterInstallProgressUnauthorized {

	return &UpdateClusterInstallProgressUnauthorized{}
}

// WithPayload adds the payload to the update cluster install progress unauthorized response
func (o *UpdateClusterInstallProgressUnauthorized) WithPayload(payload *models.InfraError) *UpdateClusterInstallProgressUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install progress unauthorized response
func (o *UpdateClusterInstallProgressUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallProgressForbiddenCode is the HTTP code returned for type UpdateClusterInstallProgressForbidden
const UpdateClusterInstallProgressForbiddenCode int = 403

/*UpdateClusterInstallProgressForbidden Forbidden.

swagger:response updateClusterInstallProgressForbidden
*/
type UpdateClusterInstallProgressForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUpdateClusterInstallProgressForbidden creates UpdateClusterInstallProgressForbidden with default headers values
func NewUpdateClusterInstallProgressForbidden() *UpdateClusterInstallProgressForbidden {

	return &UpdateClusterInstallProgressForbidden{}
}

// WithPayload adds the payload to the update cluster install progress forbidden response
func (o *UpdateClusterInstallProgressForbidden) WithPayload(payload *models.InfraError) *UpdateClusterInstallProgressForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install progress forbidden response
func (o *UpdateClusterInstallProgressForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallProgressNotFoundCode is the HTTP code returned for type UpdateClusterInstallProgressNotFound
const UpdateClusterInstallProgressNotFoundCode int = 404

/*UpdateClusterInstallProgressNotFound Error.

swagger:response updateClusterInstallProgressNotFound
*/
type UpdateClusterInstallProgressNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallProgressNotFound creates UpdateClusterInstallProgressNotFound with default headers values
func NewUpdateClusterInstallProgressNotFound() *UpdateClusterInstallProgressNotFound {

	return &UpdateClusterInstallProgressNotFound{}
}

// WithPayload adds the payload to the update cluster install progress not found response
func (o *UpdateClusterInstallProgressNotFound) WithPayload(payload *models.Error) *UpdateClusterInstallProgressNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install progress not found response
func (o *UpdateClusterInstallProgressNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallProgressMethodNotAllowedCode is the HTTP code returned for type UpdateClusterInstallProgressMethodNotAllowed
const UpdateClusterInstallProgressMethodNotAllowedCode int = 405

/*UpdateClusterInstallProgressMethodNotAllowed Method Not Allowed.

swagger:response updateClusterInstallProgressMethodNotAllowed
*/
type UpdateClusterInstallProgressMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallProgressMethodNotAllowed creates UpdateClusterInstallProgressMethodNotAllowed with default headers values
func NewUpdateClusterInstallProgressMethodNotAllowed() *UpdateClusterInstallProgressMethodNotAllowed {

	return &UpdateClusterInstallProgressMethodNotAllowed{}
}

// WithPayload adds the payload to the update cluster install progress method not allowed response
func (o *UpdateClusterInstallProgressMethodNotAllowed) WithPayload(payload *models.Error) *UpdateClusterInstallProgressMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install progress method not allowed response
func (o *UpdateClusterInstallProgressMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallProgressConflictCode is the HTTP code returned for type UpdateClusterInstallProgressConflict
const UpdateClusterInstallProgressConflictCode int = 409

/*UpdateClusterInstallProgressConflict Error.

swagger:response updateClusterInstallProgressConflict
*/
type UpdateClusterInstallProgressConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallProgressConflict creates UpdateClusterInstallProgressConflict with default headers values
func NewUpdateClusterInstallProgressConflict() *UpdateClusterInstallProgressConflict {

	return &UpdateClusterInstallProgressConflict{}
}

// WithPayload adds the payload to the update cluster install progress conflict response
func (o *UpdateClusterInstallProgressConflict) WithPayload(payload *models.Error) *UpdateClusterInstallProgressConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install progress conflict response
func (o *UpdateClusterInstallProgressConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallProgressInternalServerErrorCode is the HTTP code returned for type UpdateClusterInstallProgressInternalServerError
const UpdateClusterInstallProgressInternalServerErrorCode int = 500

/*UpdateClusterInstallProgressInternalServerError Error.

swagger:response updateClusterInstallProgressInternalServerError
*/
type UpdateClusterInstallProgressInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallProgressInternalServerError creates UpdateClusterInstallProgressInternalServerError with default headers values
func NewUpdateClusterInstallProgressInternalServerError() *UpdateClusterInstallProgressInternalServerError {

	return &UpdateClusterInstallProgressInternalServerError{}
}

// WithPayload adds the payload to the update cluster install progress internal server error response
func (o *UpdateClusterInstallProgressInternalServerError) WithPayload(payload *models.Error) *UpdateClusterInstallProgressInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install progress internal server error response
func (o *UpdateClusterInstallProgressInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateClusterInstallProgressServiceUnavailableCode is the HTTP code returned for type UpdateClusterInstallProgressServiceUnavailable
const UpdateClusterInstallProgressServiceUnavailableCode int = 503

/*UpdateClusterInstallProgressServiceUnavailable Unavailable.

swagger:response updateClusterInstallProgressServiceUnavailable
*/
type UpdateClusterInstallProgressServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterInstallProgressServiceUnavailable creates UpdateClusterInstallProgressServiceUnavailable with default headers values
func NewUpdateClusterInstallProgressServiceUnavailable() *UpdateClusterInstallProgressServiceUnavailable {

	return &UpdateClusterInstallProgressServiceUnavailable{}
}

// WithPayload adds the payload to the update cluster install progress service unavailable response
func (o *UpdateClusterInstallProgressServiceUnavailable) WithPayload(payload *models.Error) *UpdateClusterInstallProgressServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster install progress service unavailable response
func (o *UpdateClusterInstallProgressServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterInstallProgressServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/modockey/pi-home/photo-award/gen/models"
)

// GetSystemVersionOKCode is the HTTP code returned for type GetSystemVersionOK
const GetSystemVersionOKCode int = 200

/*
GetSystemVersionOK successful pet response

swagger:response getSystemVersionOK
*/
type GetSystemVersionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Version `json:"body,omitempty"`
}

// NewGetSystemVersionOK creates GetSystemVersionOK with default headers values
func NewGetSystemVersionOK() *GetSystemVersionOK {

	return &GetSystemVersionOK{}
}

// WithPayload adds the payload to the get system version o k response
func (o *GetSystemVersionOK) WithPayload(payload *models.Version) *GetSystemVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get system version o k response
func (o *GetSystemVersionOK) SetPayload(payload *models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSystemVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

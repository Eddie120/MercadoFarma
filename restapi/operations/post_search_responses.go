// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mercadofarma/services/models"
)

// PostSearchOKCode is the HTTP code returned for type PostSearchOK
const PostSearchOKCode int = 200

/*
PostSearchOK A search input created

swagger:response postSearchOK
*/
type PostSearchOK struct {

	/*
	  In: Body
	*/
	Payload *models.SearchInputResponse `json:"body,omitempty"`
}

// NewPostSearchOK creates PostSearchOK with default headers values
func NewPostSearchOK() *PostSearchOK {

	return &PostSearchOK{}
}

// WithPayload adds the payload to the post search o k response
func (o *PostSearchOK) WithPayload(payload *models.SearchInputResponse) *PostSearchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post search o k response
func (o *PostSearchOK) SetPayload(payload *models.SearchInputResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSearchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
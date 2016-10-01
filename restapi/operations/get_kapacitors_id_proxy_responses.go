package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*GetKapacitorsIDProxyNoContent Kapacitor returned no content

swagger:response getKapacitorsIdProxyNoContent
*/
type GetKapacitorsIDProxyNoContent struct {
}

// NewGetKapacitorsIDProxyNoContent creates GetKapacitorsIDProxyNoContent with default headers values
func NewGetKapacitorsIDProxyNoContent() *GetKapacitorsIDProxyNoContent {
	return &GetKapacitorsIDProxyNoContent{}
}

// WriteResponse to the client
func (o *GetKapacitorsIDProxyNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*GetKapacitorsIDProxyNotFound Kapacitor ID does not exist.

swagger:response getKapacitorsIdProxyNotFound
*/
type GetKapacitorsIDProxyNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetKapacitorsIDProxyNotFound creates GetKapacitorsIDProxyNotFound with default headers values
func NewGetKapacitorsIDProxyNotFound() *GetKapacitorsIDProxyNotFound {
	return &GetKapacitorsIDProxyNotFound{}
}

// WithPayload adds the payload to the get kapacitors Id proxy not found response
func (o *GetKapacitorsIDProxyNotFound) WithPayload(payload *models.Error) *GetKapacitorsIDProxyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get kapacitors Id proxy not found response
func (o *GetKapacitorsIDProxyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetKapacitorsIDProxyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetKapacitorsIDProxyDefault Response directly from kapacitor

swagger:response getKapacitorsIdProxyDefault
*/
type GetKapacitorsIDProxyDefault struct {
	_statusCode int

	// In: body
	Payload models.KapacitorProxyResponse `json:"body,omitempty"`
}

// NewGetKapacitorsIDProxyDefault creates GetKapacitorsIDProxyDefault with default headers values
func NewGetKapacitorsIDProxyDefault(code int) *GetKapacitorsIDProxyDefault {
	if code <= 0 {
		code = 500
	}

	return &GetKapacitorsIDProxyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get kapacitors ID proxy default response
func (o *GetKapacitorsIDProxyDefault) WithStatusCode(code int) *GetKapacitorsIDProxyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get kapacitors ID proxy default response
func (o *GetKapacitorsIDProxyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get kapacitors ID proxy default response
func (o *GetKapacitorsIDProxyDefault) WithPayload(payload models.KapacitorProxyResponse) *GetKapacitorsIDProxyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get kapacitors ID proxy default response
func (o *GetKapacitorsIDProxyDefault) SetPayload(payload models.KapacitorProxyResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetKapacitorsIDProxyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

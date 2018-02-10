// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DownloadModelDefinitionHandlerFunc turns a function with the right signature into a download model definition handler
type DownloadModelDefinitionHandlerFunc func(DownloadModelDefinitionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadModelDefinitionHandlerFunc) Handle(params DownloadModelDefinitionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DownloadModelDefinitionHandler interface for that can handle valid download model definition params
type DownloadModelDefinitionHandler interface {
	Handle(DownloadModelDefinitionParams, interface{}) middleware.Responder
}

// NewDownloadModelDefinition creates a new http.Handler for the download model definition operation
func NewDownloadModelDefinition(ctx *middleware.Context, handler DownloadModelDefinitionHandler) *DownloadModelDefinition {
	return &DownloadModelDefinition{Context: ctx, Handler: handler}
}

/*DownloadModelDefinition swagger:route GET /v1/models/{model_id}/definition Models downloadModelDefinition

Downloads the model definition.

Downloads the model definition that was initial used for training as ZIP archive.

*/
type DownloadModelDefinition struct {
	Context *middleware.Context
	Handler DownloadModelDefinitionHandler
}

func (o *DownloadModelDefinition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDownloadModelDefinitionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
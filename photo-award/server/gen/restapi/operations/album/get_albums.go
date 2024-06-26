// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAlbumsHandlerFunc turns a function with the right signature into a get albums handler
type GetAlbumsHandlerFunc func(GetAlbumsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAlbumsHandlerFunc) Handle(params GetAlbumsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAlbumsHandler interface for that can handle valid get albums params
type GetAlbumsHandler interface {
	Handle(GetAlbumsParams, interface{}) middleware.Responder
}

// NewGetAlbums creates a new http.Handler for the get albums operation
func NewGetAlbums(ctx *middleware.Context, handler GetAlbumsHandler) *GetAlbums {
	return &GetAlbums{Context: ctx, Handler: handler}
}

/*
	GetAlbums swagger:route GET /albums album getAlbums

get user's albums
*/
type GetAlbums struct {
	Context *middleware.Context
	Handler GetAlbumsHandler
}

func (o *GetAlbums) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAlbumsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

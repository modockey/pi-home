// Code generated by go-swagger; DO NOT EDIT.

package album

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostAlbumsHandlerFunc turns a function with the right signature into a post albums handler
type PostAlbumsHandlerFunc func(PostAlbumsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAlbumsHandlerFunc) Handle(params PostAlbumsParams) middleware.Responder {
	return fn(params)
}

// PostAlbumsHandler interface for that can handle valid post albums params
type PostAlbumsHandler interface {
	Handle(PostAlbumsParams) middleware.Responder
}

// NewPostAlbums creates a new http.Handler for the post albums operation
func NewPostAlbums(ctx *middleware.Context, handler PostAlbumsHandler) *PostAlbums {
	return &PostAlbums{Context: ctx, Handler: handler}
}

/*
	PostAlbums swagger:route POST /albums album postAlbums

create new album
*/
type PostAlbums struct {
	Context *middleware.Context
	Handler PostAlbumsHandler
}

func (o *PostAlbums) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAlbumsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAlbumsOKBody post albums o k body
//
// swagger:model PostAlbumsOKBody
type PostAlbumsOKBody struct {

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this post albums o k body
func (o *PostAlbumsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post albums o k body based on context it is used
func (o *PostAlbumsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAlbumsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAlbumsOKBody) UnmarshalBinary(b []byte) error {
	var res PostAlbumsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

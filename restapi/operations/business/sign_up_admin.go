// Code generated by go-swagger; DO NOT EDIT.

package business

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SignUpAdminHandlerFunc turns a function with the right signature into a sign up admin handler
type SignUpAdminHandlerFunc func(SignUpAdminParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SignUpAdminHandlerFunc) Handle(params SignUpAdminParams) middleware.Responder {
	return fn(params)
}

// SignUpAdminHandler interface for that can handle valid sign up admin params
type SignUpAdminHandler interface {
	Handle(SignUpAdminParams) middleware.Responder
}

// NewSignUpAdmin creates a new http.Handler for the sign up admin operation
func NewSignUpAdmin(ctx *middleware.Context, handler SignUpAdminHandler) *SignUpAdmin {
	return &SignUpAdmin{Context: ctx, Handler: handler}
}

/*
	SignUpAdmin swagger:route POST /v1/admin/signup business signUpAdmin

Sign up for business.

Sign up for business.
*/
type SignUpAdmin struct {
	Context *middleware.Context
	Handler SignUpAdminHandler
}

func (o *SignUpAdmin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSignUpAdminParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/mercadofarma/services/controllers"
	"go.uber.org/dig"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/mercadofarma/services/restapi/operations"
	"github.com/mercadofarma/services/restapi/operations/business"
)

//go:generate swagger generate server --target ../../mercadofarma --name Mercadofarma --spec ../swagger.json --principal interface{}

func configureFlags(api *operations.MercadofarmaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MercadofarmaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	container := buildContainer()
	invoker := func(dependency interface{}, opt ...dig.InvokeOption) {
		err := container.Invoke(dependency, opt...)
		if err != nil {
			panic(fmt.Errorf("error calling dependency %s", err))
		}
	}

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	invoker(func(controller *controllers.BusinessController) {
		api.BusinessSignUpAdminHandler = business.SignUpAdminHandlerFunc(func(params business.SignUpAdminParams) middleware.Responder {
			return controller.SignUp(params)
		})
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

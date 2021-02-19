// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/JCbayfisher/swaggerTesting/restapi/operations"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/aggregate"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/authenticate"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/change_password"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/findsummary"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/login"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/logout"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/mpp_aggregate"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/ready"
	"github.com/JCbayfisher/swaggerTesting/restapi/operations/register"
)

//go:generate swagger generate server --target ../../swaggerTesting --name PartDAPI --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.PartDAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PartDAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.RoleCustomerAuth == nil {
		api.RoleCustomerAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (role_customer) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.ReadyGetReadyHandler == nil {
		api.ReadyGetReadyHandler = ready.GetReadyHandlerFunc(func(params ready.GetReadyParams) middleware.Responder {
			return middleware.NotImplemented("operation ready.GetReady has not yet been implemented")
		})
	}
	if api.AggregatePostAggregateHandler == nil {
		api.AggregatePostAggregateHandler = aggregate.PostAggregateHandlerFunc(func(params aggregate.PostAggregateParams) middleware.Responder {
			return middleware.NotImplemented("operation aggregate.PostAggregate has not yet been implemented")
		})
	}
	if api.AuthenticatePostAuthenticateHandler == nil {
		api.AuthenticatePostAuthenticateHandler = authenticate.PostAuthenticateHandlerFunc(func(params authenticate.PostAuthenticateParams) middleware.Responder {
			return middleware.NotImplemented("operation authenticate.PostAuthenticate has not yet been implemented")
		})
	}
	if api.ChangePasswordPostChangePasswordHandler == nil {
		api.ChangePasswordPostChangePasswordHandler = change_password.PostChangePasswordHandlerFunc(func(params change_password.PostChangePasswordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation change_password.PostChangePassword has not yet been implemented")
		})
	}
	if api.FindsummaryPostFindsummaryHandler == nil {
		api.FindsummaryPostFindsummaryHandler = findsummary.PostFindsummaryHandlerFunc(func(params findsummary.PostFindsummaryParams) middleware.Responder {
			return middleware.NotImplemented("operation findsummary.PostFindsummary has not yet been implemented")
		})
	}
	if api.LoginPostLoginHandler == nil {
		api.LoginPostLoginHandler = login.PostLoginHandlerFunc(func(params login.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation login.PostLogin has not yet been implemented")
		})
	}
	if api.LogoutPostLogoutHandler == nil {
		api.LogoutPostLogoutHandler = logout.PostLogoutHandlerFunc(func(params logout.PostLogoutParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation logout.PostLogout has not yet been implemented")
		})
	}
	if api.MppAggregatePostMppaggregateHandler == nil {
		api.MppAggregatePostMppaggregateHandler = mpp_aggregate.PostMppaggregateHandlerFunc(func(params mpp_aggregate.PostMppaggregateParams) middleware.Responder {
			return middleware.NotImplemented("operation mpp_aggregate.PostMppaggregate has not yet been implemented")
		})
	}
	if api.RegisterPostRegisterHandler == nil {
		api.RegisterPostRegisterHandler = register.PostRegisterHandlerFunc(func(params register.PostRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation register.PostRegister has not yet been implemented")
		})
	}

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

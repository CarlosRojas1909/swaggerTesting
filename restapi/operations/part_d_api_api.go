// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

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

// NewPartDAPIAPI creates a new PartDAPI instance
func NewPartDAPIAPI(spec *loads.Document) *PartDAPIAPI {
	return &PartDAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ReadyGetReadyHandler: ready.GetReadyHandlerFunc(func(params ready.GetReadyParams) middleware.Responder {
			return middleware.NotImplemented("operation ready.GetReady has not yet been implemented")
		}),
		AggregatePostAggregateHandler: aggregate.PostAggregateHandlerFunc(func(params aggregate.PostAggregateParams) middleware.Responder {
			return middleware.NotImplemented("operation aggregate.PostAggregate has not yet been implemented")
		}),
		AuthenticatePostAuthenticateHandler: authenticate.PostAuthenticateHandlerFunc(func(params authenticate.PostAuthenticateParams) middleware.Responder {
			return middleware.NotImplemented("operation authenticate.PostAuthenticate has not yet been implemented")
		}),
		ChangePasswordPostChangePasswordHandler: change_password.PostChangePasswordHandlerFunc(func(params change_password.PostChangePasswordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation change_password.PostChangePassword has not yet been implemented")
		}),
		FindsummaryPostFindsummaryHandler: findsummary.PostFindsummaryHandlerFunc(func(params findsummary.PostFindsummaryParams) middleware.Responder {
			return middleware.NotImplemented("operation findsummary.PostFindsummary has not yet been implemented")
		}),
		LoginPostLoginHandler: login.PostLoginHandlerFunc(func(params login.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation login.PostLogin has not yet been implemented")
		}),
		LogoutPostLogoutHandler: logout.PostLogoutHandlerFunc(func(params logout.PostLogoutParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation logout.PostLogout has not yet been implemented")
		}),
		MppAggregatePostMppaggregateHandler: mpp_aggregate.PostMppaggregateHandlerFunc(func(params mpp_aggregate.PostMppaggregateParams) middleware.Responder {
			return middleware.NotImplemented("operation mpp_aggregate.PostMppaggregate has not yet been implemented")
		}),
		RegisterPostRegisterHandler: register.PostRegisterHandlerFunc(func(params register.PostRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation register.PostRegister has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		RoleCustomerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (role_customer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*PartDAPIAPI Medicare Provider Project partD micro service */
type PartDAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// RoleCustomerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	RoleCustomerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ReadyGetReadyHandler sets the operation handler for the get ready operation
	ReadyGetReadyHandler ready.GetReadyHandler
	// AggregatePostAggregateHandler sets the operation handler for the post aggregate operation
	AggregatePostAggregateHandler aggregate.PostAggregateHandler
	// AuthenticatePostAuthenticateHandler sets the operation handler for the post authenticate operation
	AuthenticatePostAuthenticateHandler authenticate.PostAuthenticateHandler
	// ChangePasswordPostChangePasswordHandler sets the operation handler for the post change password operation
	ChangePasswordPostChangePasswordHandler change_password.PostChangePasswordHandler
	// FindsummaryPostFindsummaryHandler sets the operation handler for the post findsummary operation
	FindsummaryPostFindsummaryHandler findsummary.PostFindsummaryHandler
	// LoginPostLoginHandler sets the operation handler for the post login operation
	LoginPostLoginHandler login.PostLoginHandler
	// LogoutPostLogoutHandler sets the operation handler for the post logout operation
	LogoutPostLogoutHandler logout.PostLogoutHandler
	// MppAggregatePostMppaggregateHandler sets the operation handler for the post mppaggregate operation
	MppAggregatePostMppaggregateHandler mpp_aggregate.PostMppaggregateHandler
	// RegisterPostRegisterHandler sets the operation handler for the post register operation
	RegisterPostRegisterHandler register.PostRegisterHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *PartDAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *PartDAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *PartDAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *PartDAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *PartDAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *PartDAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *PartDAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *PartDAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *PartDAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the PartDAPIAPI
func (o *PartDAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.RoleCustomerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.ReadyGetReadyHandler == nil {
		unregistered = append(unregistered, "ready.GetReadyHandler")
	}
	if o.AggregatePostAggregateHandler == nil {
		unregistered = append(unregistered, "aggregate.PostAggregateHandler")
	}
	if o.AuthenticatePostAuthenticateHandler == nil {
		unregistered = append(unregistered, "authenticate.PostAuthenticateHandler")
	}
	if o.ChangePasswordPostChangePasswordHandler == nil {
		unregistered = append(unregistered, "change_password.PostChangePasswordHandler")
	}
	if o.FindsummaryPostFindsummaryHandler == nil {
		unregistered = append(unregistered, "findsummary.PostFindsummaryHandler")
	}
	if o.LoginPostLoginHandler == nil {
		unregistered = append(unregistered, "login.PostLoginHandler")
	}
	if o.LogoutPostLogoutHandler == nil {
		unregistered = append(unregistered, "logout.PostLogoutHandler")
	}
	if o.MppAggregatePostMppaggregateHandler == nil {
		unregistered = append(unregistered, "mpp_aggregate.PostMppaggregateHandler")
	}
	if o.RegisterPostRegisterHandler == nil {
		unregistered = append(unregistered, "register.PostRegisterHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *PartDAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *PartDAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "role_customer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.RoleCustomerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *PartDAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *PartDAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *PartDAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *PartDAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the part d API API
func (o *PartDAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *PartDAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ready"] = ready.NewGetReady(o.context, o.ReadyGetReadyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/aggregate"] = aggregate.NewPostAggregate(o.context, o.AggregatePostAggregateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/authenticate"] = authenticate.NewPostAuthenticate(o.context, o.AuthenticatePostAuthenticateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/change_password"] = change_password.NewPostChangePassword(o.context, o.ChangePasswordPostChangePasswordHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/findsummary"] = findsummary.NewPostFindsummary(o.context, o.FindsummaryPostFindsummaryHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = login.NewPostLogin(o.context, o.LoginPostLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/logout"] = logout.NewPostLogout(o.context, o.LogoutPostLogoutHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mppaggregate"] = mpp_aggregate.NewPostMppaggregate(o.context, o.MppAggregatePostMppaggregateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/register"] = register.NewPostRegister(o.context, o.RegisterPostRegisterHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *PartDAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *PartDAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *PartDAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *PartDAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *PartDAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mahakamcloud/mahakam/pkg/api/v1/restapi/operations/apps"
	"github.com/mahakamcloud/mahakam/pkg/api/v1/restapi/operations/bare_metal_hosts"
	"github.com/mahakamcloud/mahakam/pkg/api/v1/restapi/operations/clusters"
	"github.com/mahakamcloud/mahakam/pkg/api/v1/restapi/operations/networks"
)

// NewMahakamAPI creates a new Mahakam instance
func NewMahakamAPI(spec *loads.Document) *MahakamAPI {
	return &MahakamAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		NetworksAllocateOrReleaseFromIPPoolHandler: networks.AllocateOrReleaseFromIPPoolHandlerFunc(func(params networks.AllocateOrReleaseFromIPPoolParams) middleware.Responder {
			return middleware.NotImplemented("operation NetworksAllocateOrReleaseFromIPPool has not yet been implemented")
		}),
		AppsCreateAppHandler: apps.CreateAppHandlerFunc(func(params apps.CreateAppParams) middleware.Responder {
			return middleware.NotImplemented("operation AppsCreateApp has not yet been implemented")
		}),
		ClustersCreateClusterHandler: clusters.CreateClusterHandlerFunc(func(params clusters.CreateClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersCreateCluster has not yet been implemented")
		}),
		NetworksCreateIPPoolHandler: networks.CreateIPPoolHandlerFunc(func(params networks.CreateIPPoolParams) middleware.Responder {
			return middleware.NotImplemented("operation NetworksCreateIPPool has not yet been implemented")
		}),
		NetworksCreateNetworkHandler: networks.CreateNetworkHandlerFunc(func(params networks.CreateNetworkParams) middleware.Responder {
			return middleware.NotImplemented("operation NetworksCreateNetwork has not yet been implemented")
		}),
		ClustersDescribeClustersHandler: clusters.DescribeClustersHandlerFunc(func(params clusters.DescribeClustersParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersDescribeClusters has not yet been implemented")
		}),
		AppsGetAppsHandler: apps.GetAppsHandlerFunc(func(params apps.GetAppsParams) middleware.Responder {
			return middleware.NotImplemented("operation AppsGetApps has not yet been implemented")
		}),
		BareMetalHostsGetBareMetalHostsHandler: bare_metal_hosts.GetBareMetalHostsHandlerFunc(func(params bare_metal_hosts.GetBareMetalHostsParams) middleware.Responder {
			return middleware.NotImplemented("operation BareMetalHostsGetBareMetalHosts has not yet been implemented")
		}),
		ClustersGetClustersHandler: clusters.GetClustersHandlerFunc(func(params clusters.GetClustersParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersGetClusters has not yet been implemented")
		}),
		NetworksGetNetworksHandler: networks.GetNetworksHandlerFunc(func(params networks.GetNetworksParams) middleware.Responder {
			return middleware.NotImplemented("operation NetworksGetNetworks has not yet been implemented")
		}),
		BareMetalHostsRegisterBareMetalHostHandler: bare_metal_hosts.RegisterBareMetalHostHandlerFunc(func(params bare_metal_hosts.RegisterBareMetalHostParams) middleware.Responder {
			return middleware.NotImplemented("operation BareMetalHostsRegisterBareMetalHost has not yet been implemented")
		}),
		AppsUploadAppValuesHandler: apps.UploadAppValuesHandlerFunc(func(params apps.UploadAppValuesParams) middleware.Responder {
			return middleware.NotImplemented("operation AppsUploadAppValues has not yet been implemented")
		}),
		ClustersValidateClusterHandler: clusters.ValidateClusterHandlerFunc(func(params clusters.ValidateClusterParams) middleware.Responder {
			return middleware.NotImplemented("operation ClustersValidateCluster has not yet been implemented")
		}),
	}
}

/*MahakamAPI API definitions for Mahakam Cloud Platform */
type MahakamAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// NetworksAllocateOrReleaseFromIPPoolHandler sets the operation handler for the allocate or release from Ip pool operation
	NetworksAllocateOrReleaseFromIPPoolHandler networks.AllocateOrReleaseFromIPPoolHandler
	// AppsCreateAppHandler sets the operation handler for the create app operation
	AppsCreateAppHandler apps.CreateAppHandler
	// ClustersCreateClusterHandler sets the operation handler for the create cluster operation
	ClustersCreateClusterHandler clusters.CreateClusterHandler
	// NetworksCreateIPPoolHandler sets the operation handler for the create Ip pool operation
	NetworksCreateIPPoolHandler networks.CreateIPPoolHandler
	// NetworksCreateNetworkHandler sets the operation handler for the create network operation
	NetworksCreateNetworkHandler networks.CreateNetworkHandler
	// ClustersDescribeClustersHandler sets the operation handler for the describe clusters operation
	ClustersDescribeClustersHandler clusters.DescribeClustersHandler
	// AppsGetAppsHandler sets the operation handler for the get apps operation
	AppsGetAppsHandler apps.GetAppsHandler
	// BareMetalHostsGetBareMetalHostsHandler sets the operation handler for the get bare metal hosts operation
	BareMetalHostsGetBareMetalHostsHandler bare_metal_hosts.GetBareMetalHostsHandler
	// ClustersGetClustersHandler sets the operation handler for the get clusters operation
	ClustersGetClustersHandler clusters.GetClustersHandler
	// NetworksGetNetworksHandler sets the operation handler for the get networks operation
	NetworksGetNetworksHandler networks.GetNetworksHandler
	// BareMetalHostsRegisterBareMetalHostHandler sets the operation handler for the register bare metal host operation
	BareMetalHostsRegisterBareMetalHostHandler bare_metal_hosts.RegisterBareMetalHostHandler
	// AppsUploadAppValuesHandler sets the operation handler for the upload app values operation
	AppsUploadAppValuesHandler apps.UploadAppValuesHandler
	// ClustersValidateClusterHandler sets the operation handler for the validate cluster operation
	ClustersValidateClusterHandler clusters.ValidateClusterHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *MahakamAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MahakamAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MahakamAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MahakamAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MahakamAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MahakamAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MahakamAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MahakamAPI
func (o *MahakamAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.NetworksAllocateOrReleaseFromIPPoolHandler == nil {
		unregistered = append(unregistered, "networks.AllocateOrReleaseFromIPPoolHandler")
	}

	if o.AppsCreateAppHandler == nil {
		unregistered = append(unregistered, "apps.CreateAppHandler")
	}

	if o.ClustersCreateClusterHandler == nil {
		unregistered = append(unregistered, "clusters.CreateClusterHandler")
	}

	if o.NetworksCreateIPPoolHandler == nil {
		unregistered = append(unregistered, "networks.CreateIPPoolHandler")
	}

	if o.NetworksCreateNetworkHandler == nil {
		unregistered = append(unregistered, "networks.CreateNetworkHandler")
	}

	if o.ClustersDescribeClustersHandler == nil {
		unregistered = append(unregistered, "clusters.DescribeClustersHandler")
	}

	if o.AppsGetAppsHandler == nil {
		unregistered = append(unregistered, "apps.GetAppsHandler")
	}

	if o.BareMetalHostsGetBareMetalHostsHandler == nil {
		unregistered = append(unregistered, "bare_metal_hosts.GetBareMetalHostsHandler")
	}

	if o.ClustersGetClustersHandler == nil {
		unregistered = append(unregistered, "clusters.GetClustersHandler")
	}

	if o.NetworksGetNetworksHandler == nil {
		unregistered = append(unregistered, "networks.GetNetworksHandler")
	}

	if o.BareMetalHostsRegisterBareMetalHostHandler == nil {
		unregistered = append(unregistered, "bare_metal_hosts.RegisterBareMetalHostHandler")
	}

	if o.AppsUploadAppValuesHandler == nil {
		unregistered = append(unregistered, "apps.UploadAppValuesHandler")
	}

	if o.ClustersValidateClusterHandler == nil {
		unregistered = append(unregistered, "clusters.ValidateClusterHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MahakamAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MahakamAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *MahakamAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *MahakamAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *MahakamAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
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
func (o *MahakamAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the mahakam API
func (o *MahakamAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MahakamAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/networks/pools/ipPools/{poolId}"] = networks.NewAllocateOrReleaseFromIPPool(o.context, o.NetworksAllocateOrReleaseFromIPPoolHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/apps"] = apps.NewCreateApp(o.context, o.AppsCreateAppHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters"] = clusters.NewCreateCluster(o.context, o.ClustersCreateClusterHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/networks/pools/ipPools"] = networks.NewCreateIPPool(o.context, o.NetworksCreateIPPoolHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/networks"] = networks.NewCreateNetwork(o.context, o.NetworksCreateNetworkHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters/describe"] = clusters.NewDescribeClusters(o.context, o.ClustersDescribeClustersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/apps"] = apps.NewGetApps(o.context, o.AppsGetAppsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/bare-metal-hosts"] = bare_metal_hosts.NewGetBareMetalHosts(o.context, o.BareMetalHostsGetBareMetalHostsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/clusters"] = clusters.NewGetClusters(o.context, o.ClustersGetClustersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/networks"] = networks.NewGetNetworks(o.context, o.NetworksGetNetworksHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/bare-metal-hosts"] = bare_metal_hosts.NewRegisterBareMetalHost(o.context, o.BareMetalHostsRegisterBareMetalHostHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/apps/values"] = apps.NewUploadAppValues(o.context, o.AppsUploadAppValuesHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/clusters/validate"] = clusters.NewValidateCluster(o.context, o.ClustersValidateClusterHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MahakamAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MahakamAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MahakamAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MahakamAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

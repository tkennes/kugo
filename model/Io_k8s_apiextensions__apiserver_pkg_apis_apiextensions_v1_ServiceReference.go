package model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_WebhookClientConfig.go


// ServiceReference holds a reference to Service.legacy.k8s.io
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_ServiceReference struct {
	// name is the name of the service. Required
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name      *string `json:"name"`

	// namespace is the namespace of the service. Required
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace *string `json:"namespace"`

	// path is an optional URL path at which the webhook will be contacted.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path      *string `json:"path,omitempty"`

	// port is an optional service port at which the webhook will be contacted. `port` should be a valid port number (1-65535,
	// inclusive). Defaults to 443 for backward compatibility.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Port      *int    `json:"port,omitempty"`
}

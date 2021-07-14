package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_admissionregistration_v1beta1_WebhookClientConfig.go


// ServiceReference holds a reference to Service.legacy.k8s.io
type Io_k8s_api_admissionregistration_v1beta1_ServiceReference struct {
	// `name` is the name of the service. Required
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name      *string `json:"name"`

	// `namespace` is the namespace of the service. Required
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace *string `json:"namespace"`

	// `path` is an optional URL path which will be sent in any request to this service.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path      *string `json:"path,omitempty"`

	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be
	// a valid port number (1-65535, inclusive).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Port      *int    `json:"port,omitempty"`
}

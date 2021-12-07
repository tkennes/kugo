package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceSpec.go


// ServiceReference holds a reference to Service.legacy.k8s.io
type Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_ServiceReference struct {
	// Name is the name of the service
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name      *string `json:"name,omitempty"`

	// Namespace is the namespace of the service
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Namespace *string `json:"namespace,omitempty"`

	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be
	// a valid port number (1-65535, inclusive).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Port      *int    `json:"port,omitempty"`
}

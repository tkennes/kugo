package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_HTTPIngressPath.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_IngressSpec.go


// IngressBackend describes all endpoints for a given service and port.
type Io_k8s_api_networking_v1_IngressBackend struct {
	// Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is
	// specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_TypedLocalObjectReference.go
	Resource *Io_k8s_api_core_v1_TypedLocalObjectReference   `json:"resource,omitempty"`

	// Service references a Service as a Backend. This is a mutually exclusive setting with "Resource".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_IngressServiceBackend.go
	Service  *Io_k8s_api_networking_v1_IngressServiceBackend `json:"service,omitempty"`
}

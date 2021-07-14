package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_HTTPIngressPath.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressSpec.go


// IngressBackend describes all endpoints for a given service and port.
type Io_k8s_api_networking_v1beta1_IngressBackend struct {
	// Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is
	// specified, serviceName and servicePort must not be specified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_TypedLocalObjectReference.go
	Resource    *Io_k8s_api_core_v1_TypedLocalObjectReference `json:"resource,omitempty"`

	// Specifies the name of the referenced service.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ServiceName *string                                       `json:"serviceName,omitempty"`

	// Specifies the port of the referenced service.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ServicePort *int                                          `json:"servicePort,omitempty"`
}

package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_IngressBackend.go


// IngressServiceBackend references a Kubernetes Service as a Backend.
type Io_k8s_api_networking_v1_IngressServiceBackend struct {
	// Name is the referenced service. The service must exist in the same namespace as the Ingress object.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name *string                                      `json:"name"`

	// Port of the referenced service. A port name or port number is required for a IngressServiceBackend.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_ServiceBackendPort.go
	Port *Io_k8s_api_networking_v1_ServiceBackendPort `json:"port,omitempty"`
}

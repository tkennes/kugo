package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ServiceStatus.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_IngressStatus.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1beta1_IngressStatus.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_extensions_v1beta1_IngressStatus.go


// LoadBalancerStatus represents the status of a load-balancer.
type Io_k8s_api_core_v1_LoadBalancerStatus struct {
	// Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to
	// these ingress points.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_LoadBalancerIngress.go
	Ingress []Io_k8s_api_core_v1_LoadBalancerIngress `json:"ingress,omitempty"`
}

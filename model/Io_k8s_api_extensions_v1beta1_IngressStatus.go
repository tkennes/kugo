package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_extensions_v1beta1_Ingress.go


// IngressStatus describe the current state of the Ingress.
type Io_k8s_api_extensions_v1beta1_IngressStatus struct {
	// LoadBalancer contains the current status of the load-balancer.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LoadBalancerStatus.go
	LoadBalancer *Io_k8s_api_core_v1_LoadBalancerStatus `json:"loadBalancer,omitempty"`
}

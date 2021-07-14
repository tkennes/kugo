package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Service.go


// ServiceStatus represents the current status of a service.
type Io_k8s_api_core_v1_ServiceStatus struct {
	// Current service state
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_Condition.go
	Conditions   []Io_k8s_apimachinery_pkg_apis_meta_v1_Condition `json:"conditions,omitempty"`

	// LoadBalancer contains the current status of the load-balancer, if one is present.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LoadBalancerStatus.go
	LoadBalancer *Io_k8s_api_core_v1_LoadBalancerStatus           `json:"loadBalancer,omitempty"`
}

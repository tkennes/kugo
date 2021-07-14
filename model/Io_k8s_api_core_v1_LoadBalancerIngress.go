package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LoadBalancerStatus.go


// LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be
// sent to an ingress point.
type Io_k8s_api_core_v1_LoadBalancerIngress struct {
	// Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Hostname *string                         `json:"hostname,omitempty"`

	// IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Ip       *string                         `json:"ip,omitempty"`

	// Ports is a list of records of service ports If used, every port defined in the service should have an entry in it
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PortStatus.go
	Ports    []Io_k8s_api_core_v1_PortStatus `json:"ports,omitempty"`
}

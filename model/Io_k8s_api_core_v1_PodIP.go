package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodStatus.go


// IP address information for entries in the (plural) PodIPs field. Each entry includes:    IP: An IP address allocated to
// the pod. Routable at least within the cluster.
type Io_k8s_api_core_v1_PodIP struct {
	// ip is an IP address (IPv4 or IPv6) assigned to the pod
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Ip *string `json:"ip,omitempty"`
}

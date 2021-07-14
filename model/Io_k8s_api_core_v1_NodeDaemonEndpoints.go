package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeStatus.go


// NodeDaemonEndpoints lists ports opened by daemons running on the Node.
type Io_k8s_api_core_v1_NodeDaemonEndpoints struct {
	// Endpoint on which Kubelet is listening.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_DaemonEndpoint.go
	KubeletEndpoint *Io_k8s_api_core_v1_DaemonEndpoint `json:"kubeletEndpoint,omitempty"`
}

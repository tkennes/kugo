package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeDaemonEndpoints.go


// DaemonEndpoint contains information about a single Daemon endpoint.
type Io_k8s_api_core_v1_DaemonEndpoint struct {
	// Port number of the given endpoint.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Port *int `json:"Port"`
}

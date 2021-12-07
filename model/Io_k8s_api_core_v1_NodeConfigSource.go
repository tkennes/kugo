package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeConfigStatus.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeSpec.go


// NodeConfigSource specifies a source of node configuration. Exactly one subfield (excluding metadata) must be non-nil.
type Io_k8s_api_core_v1_NodeConfigSource struct {
	// ConfigMap is a reference to a Node's ConfigMap
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ConfigMapNodeConfigSource.go
	ConfigMap *Io_k8s_api_core_v1_ConfigMapNodeConfigSource `json:"configMap,omitempty"`
}

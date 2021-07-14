package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeSelector.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PreferredSchedulingTerm.go


// A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type
// implements a subset of the NodeSelectorTerm.
type Io_k8s_api_core_v1_NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeSelectorRequirement.go
	MatchExpressions []Io_k8s_api_core_v1_NodeSelectorRequirement `json:"matchExpressions,omitempty"`

	// A list of node selector requirements by node's fields.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeSelectorRequirement.go
	MatchFields      []Io_k8s_api_core_v1_NodeSelectorRequirement `json:"matchFields,omitempty"`
}

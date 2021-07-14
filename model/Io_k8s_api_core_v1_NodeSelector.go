package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeAffinity.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_VolumeNodeAffinity.go


// A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it
// represents the OR of the selectors represented by the node selector terms.
type Io_k8s_api_core_v1_NodeSelector struct {
	// Required. A list of node selector terms. The terms are ORed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeSelectorTerm.go
	NodeSelectorTerms []Io_k8s_api_core_v1_NodeSelectorTerm `json:"nodeSelectorTerms"`
}

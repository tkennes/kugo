package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeAffinity.go


// An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred
// scheduling term matches no objects (i.e. is also a no-op).
type Io_k8s_api_core_v1_PreferredSchedulingTerm struct {
	// A node selector term, associated with the corresponding weight.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeSelectorTerm.go
	Preference Io_k8s_api_core_v1_NodeSelectorTerm `json:"preference"`

	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Weight     *int                                `json:"weight"`
}

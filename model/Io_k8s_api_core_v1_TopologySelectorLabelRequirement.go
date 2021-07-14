package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_TopologySelectorTerm.go


// A topology selector requirement is a selector that matches given label. This is an alpha feature and may change in the
// future.
type Io_k8s_api_core_v1_TopologySelectorLabelRequirement struct {
	// The label key that the selector applies to.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Key    *string   `json:"key"`

	// An array of string values. One value must match the label to be selected. Each entry in Values is ORed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Values []*string `json:"values"`
}

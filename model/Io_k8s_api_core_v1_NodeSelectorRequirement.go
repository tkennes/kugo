package kugo_model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeSelectorTerm.go


// A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type Io_k8s_api_core_v1_NodeSelectorRequirement struct {
	// The label key that the selector applies to.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Key      *string   `json:"key"`

	// Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Operator *string   `json:"operator"`

	// An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists
	// or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single
	// element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Values   []*string `json:"values,omitempty"`
}

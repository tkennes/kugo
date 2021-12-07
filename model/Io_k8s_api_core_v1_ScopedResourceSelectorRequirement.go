package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ScopeSelector.go


// A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates
// the scope name and values.
type Io_k8s_api_core_v1_ScopedResourceSelectorRequirement struct {
	// Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Operator  *string   `json:"operator"`

	// The name of the scope that the selector applies to.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ScopeName *string   `json:"scopeName"`

	// An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists
	// or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Values    []*string `json:"values,omitempty"`
}

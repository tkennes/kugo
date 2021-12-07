package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go


// A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelectorRequirement struct {
	// key is the label key that the selector applies to.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Key      *string   `json:"key"`

	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Operator *string   `json:"operator"`

	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator
	// is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Values   []*string `json:"values,omitempty"`
}

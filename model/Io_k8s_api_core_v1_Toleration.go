package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_node_v1_Scheduling.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_node_v1alpha1_Scheduling.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_node_v1beta1_Scheduling.go


// The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching
// operator <operator>.
type Io_k8s_api_core_v1_Toleration struct {
	// Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are
	// NoSchedule, PreferNoSchedule and NoExecute.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Effect            *string `json:"effect,omitempty"`

	// Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator
	// must be Exists; this combination means to match all values and all keys.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Key               *string `json:"key,omitempty"`

	// Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists
	// is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Operator          *string `json:"operator,omitempty"`

	// TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field
	// is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero
	// and negative values will be treated as 0 (evict immediately) by the system.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	TolerationSeconds *int    `json:"tolerationSeconds,omitempty"`

	// Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just
	// a regular string.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Value             *string `json:"value,omitempty"`
}

package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_LimitRange.go


// LimitRangeSpec defines a min/max usage limit for resources that match on kind.
type Io_k8s_api_core_v1_LimitRangeSpec struct {
	// Limits is the list of LimitRangeItem objects that are enforced.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_LimitRangeItem.go
	Limits []Io_k8s_api_core_v1_LimitRangeItem `json:"limits"`
}

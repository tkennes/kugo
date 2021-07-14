package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LimitRangeSpec.go


// LimitRangeItem defines a min/max usage limit for any resource that matches on kind.
type Io_k8s_api_core_v1_LimitRangeItem struct {
	// Default resource requirement limit value by resource name if resource limit is omitted.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Default              *interface{} `json:"default,omitempty"`

	// DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	DefaultRequest       *interface{} `json:"defaultRequest,omitempty"`

	// Max usage constraints on this kind by resource name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Max                  *interface{} `json:"max,omitempty"`

	// MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit
	// divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	MaxLimitRequestRatio *interface{} `json:"maxLimitRequestRatio,omitempty"`

	// Min usage constraints on this kind by resource name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Min                  *interface{} `json:"min,omitempty"`

	// Type of resource that this limit applies to.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type                 *string      `json:"type"`
}

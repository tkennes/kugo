package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ResourceQuotaSpec.go


// A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.
type Io_k8s_api_core_v1_ScopeSelector struct {
	// A list of scope selector requirements by scope of the resources.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ScopedResourceSelectorRequirement.go
	MatchExpressions []Io_k8s_api_core_v1_ScopedResourceSelectorRequirement `json:"matchExpressions,omitempty"`
}

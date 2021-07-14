package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ResourceQuota.go


// ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
type Io_k8s_api_core_v1_ResourceQuotaSpec struct {
	// hard is the set of desired hard limits for each named resource. More info:
	// https://kubernetes.io/docs/concepts/policy/resource-quotas/
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Hard          *interface{}                      `json:"hard,omitempty"`

	// scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed
	// using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector
	// (if specified in spec), must be matched.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ScopeSelector.go
	ScopeSelector *Io_k8s_api_core_v1_ScopeSelector `json:"scopeSelector,omitempty"`

	// A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Scopes        []*string                         `json:"scopes,omitempty"`
}

package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ResourceQuota.go


// ResourceQuotaStatus defines the enforced hard limits and observed use.
type Io_k8s_api_core_v1_ResourceQuotaStatus struct {
	// Hard is the set of enforced hard limits for each named resource. More info:
	// https://kubernetes.io/docs/concepts/policy/resource-quotas/
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Hard *interface{} `json:"hard,omitempty"`

	// Used is the current observed total usage of the resource in the namespace.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Used *interface{} `json:"used,omitempty"`
}

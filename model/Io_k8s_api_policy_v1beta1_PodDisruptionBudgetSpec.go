package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodDisruptionBudget.go


// PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
type Io_k8s_api_policy_v1beta1_PodDisruptionBudgetSpec struct {
	// An eviction is allowed if at most "maxUnavailable" pods selected by "selector" are unavailable after the eviction, i.e.
	// even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a
	// mutually exclusive setting with "minAvailable".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	MaxUnavailable *int                                                `json:"maxUnavailable,omitempty"`

	// An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be available after the
	// eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by
	// specifying "100%".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	MinAvailable   *int                                                `json:"minAvailable,omitempty"`

	// Label query over pods whose evictions are managed by the disruption budget.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector       *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
}

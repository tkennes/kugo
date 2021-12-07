package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Affinity.go


// Pod affinity is a group of inter pod affinity scheduling rules.
type Io_k8s_api_core_v1_PodAffinity struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but
	// it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the
	// greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request,
	// requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and
	// adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the
	// highest sum are the most preferred.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_WeightedPodAffinityTerm.go
	PreferredDuringSchedulingIgnoredDuringExecution []Io_k8s_api_core_v1_WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

	// If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto
	// the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g.
	// due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are
	// multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be
	// satisfied.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodAffinityTerm.go
	RequiredDuringSchedulingIgnoredDuringExecution  []Io_k8s_api_core_v1_PodAffinityTerm         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

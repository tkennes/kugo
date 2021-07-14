package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodAffinity.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodAntiAffinity.go


// The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
type Io_k8s_api_core_v1_WeightedPodAffinityTerm struct {
	// Required. A pod affinity term, associated with the corresponding weight.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodAffinityTerm.go
	PodAffinityTerm Io_k8s_api_core_v1_PodAffinityTerm `json:"podAffinityTerm"`

	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Weight          *int                               `json:"weight"`
}

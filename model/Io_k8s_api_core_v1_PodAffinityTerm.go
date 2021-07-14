package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodAffinity.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodAntiAffinity.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_WeightedPodAffinityTerm.go


// Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should
// be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose
// value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
type Io_k8s_api_core_v1_PodAffinityTerm struct {
	// A label query over a set of resources, in this case pods.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	LabelSelector *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"labelSelector,omitempty"`

	// namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means "this
	// pod's namespace"
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespaces    []*string                                           `json:"namespaces,omitempty"`

	// This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in
	// the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey
	// matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	TopologyKey   *string                                             `json:"topologyKey"`
}

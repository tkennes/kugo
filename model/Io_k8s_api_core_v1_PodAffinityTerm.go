package kugo_model

type Io_k8s_api_core_v1_PodAffinityTerm struct {
	LabelSelector *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"labelSelector,omitempty"`
	Namespaces    []string                                            `json:"namespaces,omitempty"`
	TopologyKey   string                                              `json:"topologyKey"`
}


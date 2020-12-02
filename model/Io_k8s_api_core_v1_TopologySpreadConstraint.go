package kugo_model

type Io_k8s_api_core_v1_TopologySpreadConstraint struct {
	LabelSelector     *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"labelSelector,omitempty"`
	MaxSkew           int                                                 `json:"maxSkew"`
	TopologyKey       string                                              `json:"topologyKey"`
	WhenUnsatisfiable string                                              `json:"whenUnsatisfiable"`
}


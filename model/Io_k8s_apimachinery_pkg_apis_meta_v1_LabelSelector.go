package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector struct {
	MatchExpressions []Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchLabels      *interface{}                                                    `json:"matchLabels,omitempty"`
}


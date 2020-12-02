package kugo_model

type Io_k8s_api_networking_v1_NetworkPolicySpec struct {
	Egress      []Io_k8s_api_networking_v1_NetworkPolicyEgressRule  `json:"egress,omitempty"`
	Ingress     []Io_k8s_api_networking_v1_NetworkPolicyIngressRule `json:"ingress,omitempty"`
	PodSelector Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector  `json:"podSelector"`
	PolicyTypes []string                                            `json:"policyTypes,omitempty"`
}


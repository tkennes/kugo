package kugo_model

type Io_k8s_api_networking_v1_NetworkPolicyPeer struct {
	IpBlock           *Io_k8s_api_networking_v1_IPBlock                   `json:"ipBlock,omitempty"`
	NamespaceSelector *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"namespaceSelector,omitempty"`
	PodSelector       *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"podSelector,omitempty"`
}


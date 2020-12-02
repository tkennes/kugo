package kugo_model

type Io_k8s_api_networking_v1beta1_Ingress struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Spec       *Io_k8s_api_networking_v1beta1_IngressSpec       `json:"spec,omitempty"`
	Status     *Io_k8s_api_networking_v1beta1_IngressStatus     `json:"status,omitempty"`
}


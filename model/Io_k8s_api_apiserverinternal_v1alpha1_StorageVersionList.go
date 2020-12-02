package kugo_model

type Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionList struct {
	ApiVersion string                                                 `json:"apiVersion,omitempty"`
	Items      []Io_k8s_api_apiserverinternal_v1alpha1_StorageVersion `json:"items"`
	Kind       string                                                 `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta         `json:"metadata,omitempty"`
}


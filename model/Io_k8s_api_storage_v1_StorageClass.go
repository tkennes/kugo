package kugo_model

type Io_k8s_api_storage_v1_StorageClass struct {
	AllowVolumeExpansion bool                                             `json:"allowVolumeExpansion,omitempty"`
	AllowedTopologies    []Io_k8s_api_core_v1_TopologySelectorTerm        `json:"allowedTopologies,omitempty"`
	ApiVersion           string                                           `json:"apiVersion,omitempty"`
	Kind                 string                                           `json:"kind,omitempty"`
	Metadata             *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	MountOptions         []string                                         `json:"mountOptions,omitempty"`
	Parameters           *interface{}                                     `json:"parameters,omitempty"`
	Provisioner          string                                           `json:"provisioner"`
	ReclaimPolicy        string                                           `json:"reclaimPolicy,omitempty"`
	VolumeBindingMode    string                                           `json:"volumeBindingMode,omitempty"`
}


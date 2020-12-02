package kugo_model

type Io_k8s_api_core_v1_PersistentVolumeClaimSpec struct {
	AccessModes      []string                                            `json:"accessModes,omitempty"`
	DataSource       *Io_k8s_api_core_v1_TypedLocalObjectReference       `json:"dataSource,omitempty"`
	Resources        *Io_k8s_api_core_v1_ResourceRequirements            `json:"resources,omitempty"`
	Selector         *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
	StorageClassName string                                              `json:"storageClassName,omitempty"`
	VolumeMode       string                                              `json:"volumeMode,omitempty"`
	VolumeName       string                                              `json:"volumeName,omitempty"`
}


package kugo_model

type Io_k8s_api_core_v1_CSIPersistentVolumeSource struct {
	ControllerExpandSecretRef  *Io_k8s_api_core_v1_SecretReference `json:"controllerExpandSecretRef,omitempty"`
	ControllerPublishSecretRef *Io_k8s_api_core_v1_SecretReference `json:"controllerPublishSecretRef,omitempty"`
	Driver                     string                              `json:"driver"`
	FsType                     string                              `json:"fsType,omitempty"`
	NodePublishSecretRef       *Io_k8s_api_core_v1_SecretReference `json:"nodePublishSecretRef,omitempty"`
	NodeStageSecretRef         *Io_k8s_api_core_v1_SecretReference `json:"nodeStageSecretRef,omitempty"`
	ReadOnly                   bool                                `json:"readOnly,omitempty"`
	VolumeAttributes           *interface{}                        `json:"volumeAttributes,omitempty"`
	VolumeHandle               string                              `json:"volumeHandle"`
}


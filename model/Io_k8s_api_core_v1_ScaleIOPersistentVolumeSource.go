package kugo_model

type Io_k8s_api_core_v1_ScaleIOPersistentVolumeSource struct {
	FsType           string                             `json:"fsType,omitempty"`
	Gateway          string                             `json:"gateway"`
	ProtectionDomain string                             `json:"protectionDomain,omitempty"`
	ReadOnly         bool                               `json:"readOnly,omitempty"`
	SecretRef        Io_k8s_api_core_v1_SecretReference `json:"secretRef"`
	SslEnabled       bool                               `json:"sslEnabled,omitempty"`
	StorageMode      string                             `json:"storageMode,omitempty"`
	StoragePool      string                             `json:"storagePool,omitempty"`
	System           string                             `json:"system"`
	VolumeName       string                             `json:"volumeName,omitempty"`
}


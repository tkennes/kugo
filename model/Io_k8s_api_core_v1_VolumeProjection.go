package kugo_model

type Io_k8s_api_core_v1_VolumeProjection struct {
	ConfigMap           *Io_k8s_api_core_v1_ConfigMapProjection           `json:"configMap,omitempty"`
	DownwardAPI         *Io_k8s_api_core_v1_DownwardAPIProjection         `json:"downwardAPI,omitempty"`
	Secret              *Io_k8s_api_core_v1_SecretProjection              `json:"secret,omitempty"`
	ServiceAccountToken *Io_k8s_api_core_v1_ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty"`
}


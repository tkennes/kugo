package kugo_model

type Io_k8s_api_core_v1_AzureFileVolumeSource struct {
	ReadOnly   bool   `json:"readOnly,omitempty"`
	SecretName string `json:"secretName"`
	ShareName  string `json:"shareName"`
}


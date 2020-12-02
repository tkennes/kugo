package kugo_model

type Io_k8s_api_core_v1_AzureFilePersistentVolumeSource struct {
	ReadOnly        bool   `json:"readOnly,omitempty"`
	SecretName      string `json:"secretName"`
	SecretNamespace string `json:"secretNamespace,omitempty"`
	ShareName       string `json:"shareName"`
}


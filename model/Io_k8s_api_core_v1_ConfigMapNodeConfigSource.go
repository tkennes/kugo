package kugo_model

type Io_k8s_api_core_v1_ConfigMapNodeConfigSource struct {
	KubeletConfigKey string `json:"kubeletConfigKey"`
	Name             string `json:"name"`
	Namespace        string `json:"namespace"`
	ResourceVersion  string `json:"resourceVersion,omitempty"`
	Uid              string `json:"uid,omitempty"`
}


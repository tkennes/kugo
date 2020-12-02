package kugo_model

type Io_k8s_api_core_v1_EnvFromSource struct {
	ConfigMapRef *Io_k8s_api_core_v1_ConfigMapEnvSource `json:"configMapRef,omitempty"`
	Prefix       string                                 `json:"prefix,omitempty"`
	SecretRef    *Io_k8s_api_core_v1_SecretEnvSource    `json:"secretRef,omitempty"`
}


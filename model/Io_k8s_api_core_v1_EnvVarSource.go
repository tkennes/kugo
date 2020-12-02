package kugo_model

type Io_k8s_api_core_v1_EnvVarSource struct {
	ConfigMapKeyRef  *Io_k8s_api_core_v1_ConfigMapKeySelector  `json:"configMapKeyRef,omitempty"`
	FieldRef         *Io_k8s_api_core_v1_ObjectFieldSelector   `json:"fieldRef,omitempty"`
	ResourceFieldRef *Io_k8s_api_core_v1_ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	SecretKeyRef     *Io_k8s_api_core_v1_SecretKeySelector     `json:"secretKeyRef,omitempty"`
}


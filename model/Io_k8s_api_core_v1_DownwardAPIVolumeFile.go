package kugo_model

type Io_k8s_api_core_v1_DownwardAPIVolumeFile struct {
	FieldRef         *Io_k8s_api_core_v1_ObjectFieldSelector   `json:"fieldRef,omitempty"`
	Mode             int                                       `json:"mode,omitempty"`
	Path             string                                    `json:"path"`
	ResourceFieldRef *Io_k8s_api_core_v1_ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}


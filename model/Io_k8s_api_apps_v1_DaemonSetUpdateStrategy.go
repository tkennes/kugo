package kugo_model

type Io_k8s_api_apps_v1_DaemonSetUpdateStrategy struct {
	RollingUpdate *Io_k8s_api_apps_v1_RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"`
	Type          string                                     `json:"type,omitempty"`
}


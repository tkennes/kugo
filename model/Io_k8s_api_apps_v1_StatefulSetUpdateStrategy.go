package kugo_model

type Io_k8s_api_apps_v1_StatefulSetUpdateStrategy struct {
	RollingUpdate *Io_k8s_api_apps_v1_RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"`
	Type          string                                               `json:"type,omitempty"`
}


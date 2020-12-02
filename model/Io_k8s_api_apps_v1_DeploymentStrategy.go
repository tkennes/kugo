package kugo_model

type Io_k8s_api_apps_v1_DeploymentStrategy struct {
	RollingUpdate *Io_k8s_api_apps_v1_RollingUpdateDeployment `json:"rollingUpdate,omitempty"`
	Type          string                                      `json:"type,omitempty"`
}


package kugo_model

type Io_k8s_api_autoscaling_v1_ScaleStatus struct {
	Replicas int    `json:"replicas"`
	Selector string `json:"selector,omitempty"`
}


package kugo_model

type Io_k8s_api_autoscaling_v2beta2_HPAScalingRules struct {
	Policies                   []Io_k8s_api_autoscaling_v2beta2_HPAScalingPolicy `json:"policies,omitempty"`
	SelectPolicy               string                                            `json:"selectPolicy,omitempty"`
	StabilizationWindowSeconds int                                               `json:"stabilizationWindowSeconds,omitempty"`
}


package kugo_model

type Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerBehavior struct {
	ScaleDown *Io_k8s_api_autoscaling_v2beta2_HPAScalingRules `json:"scaleDown,omitempty"`
	ScaleUp   *Io_k8s_api_autoscaling_v2beta2_HPAScalingRules `json:"scaleUp,omitempty"`
}


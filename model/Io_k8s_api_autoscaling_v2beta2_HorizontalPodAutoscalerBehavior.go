package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerSpec.go


// HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp
// and scaleDown fields respectively).
type Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerBehavior struct {
	// scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas
	// pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_HPAScalingRules.go
	ScaleDown *Io_k8s_api_autoscaling_v2beta2_HPAScalingRules `json:"scaleDown,omitempty"`

	// scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of:   * increase no more than 4
	// pods per 60 seconds   * double the number of pods per 60 seconds No stabilization is used.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_HPAScalingRules.go
	ScaleUp   *Io_k8s_api_autoscaling_v2beta2_HPAScalingRules `json:"scaleUp,omitempty"`
}

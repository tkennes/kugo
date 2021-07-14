package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerBehavior.go


// HPAScalingRules configures the scaling behavior for one direction. These Rules are applied after calculating
// DesiredReplicas from metrics for the HPA. They can limit the scaling velocity by specifying scaling policies. They can
// prevent flapping by specifying the stabilization window, so that the number of replicas is not set instantly, instead,
// the safest value from the stabilization window is chosen.
type Io_k8s_api_autoscaling_v2beta2_HPAScalingRules struct {
	// policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified,
	// otherwise the HPAScalingRules will be discarded as invalid
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_HPAScalingPolicy.go
	Policies                   []Io_k8s_api_autoscaling_v2beta2_HPAScalingPolicy `json:"policies,omitempty"`

	// selectPolicy is used to specify which policy should be used. If not set, the default value MaxPolicySelect is used.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SelectPolicy               *string                                           `json:"selectPolicy,omitempty"`

	// StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up
	// or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one
	// hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300
	// (i.e. the stabilization window is 300 seconds long).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	StabilizationWindowSeconds *int                                              `json:"stabilizationWindowSeconds,omitempty"`
}

package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v1_HorizontalPodAutoscaler.go


// specification of a horizontal pod autoscaler.
type Io_k8s_api_autoscaling_v1_HorizontalPodAutoscalerSpec struct {
	// upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	MaxReplicas                    *int                                                  `json:"maxReplicas"`

	// minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.
	// minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External
	// metric is configured.  Scaling is active as long as at least one metric value is available.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	MinReplicas                    *int                                                  `json:"minReplicas,omitempty"`

	// reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the
	// desired number of pods by using its Scale subresource.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v1_CrossVersionObjectReference.go
	ScaleTargetRef                 Io_k8s_api_autoscaling_v1_CrossVersionObjectReference `json:"scaleTargetRef"`

	// target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the
	// default autoscaling policy will be used.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	TargetCPUUtilizationPercentage *int                                                  `json:"targetCPUUtilizationPercentage,omitempty"`
}

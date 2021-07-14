package kugo_model

import (
	"time"
)


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscaler.go


// HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
type Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerStatus struct {
	// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those
	// conditions are met.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerCondition.go
	Conditions         []Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerCondition `json:"conditions"`

	// currentMetrics is the last read state of the metrics used by this autoscaler.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_MetricStatus.go
	CurrentMetrics     []Io_k8s_api_autoscaling_v2beta2_MetricStatus                     `json:"currentMetrics,omitempty"`

	// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CurrentReplicas    *int                                                              `json:"currentReplicas"`

	// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the
	// autoscaler.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	DesiredReplicas    *int                                                              `json:"desiredReplicas"`

	// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control
	// how often the number of pods is changed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastScaleTime      *time.Time                                                        `json:"lastScaleTime,omitempty"`

	// observedGeneration is the most recent generation observed by this autoscaler.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration *int                                                              `json:"observedGeneration,omitempty"`
}

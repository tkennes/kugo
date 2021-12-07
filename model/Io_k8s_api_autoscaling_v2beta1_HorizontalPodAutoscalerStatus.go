package kugo_model

import (
	"time"
)


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscaler.go


// HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
type Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerStatus struct {
	// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those
	// conditions are met.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerCondition.go
	Conditions         []Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerCondition `json:"conditions"`

	// currentMetrics is the last read state of the metrics used by this autoscaler.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta1_MetricStatus.go
	CurrentMetrics     []Io_k8s_api_autoscaling_v2beta1_MetricStatus                     `json:"currentMetrics,omitempty"`

	// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	CurrentReplicas    *int                                                              `json:"currentReplicas"`

	// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the
	// autoscaler.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	DesiredReplicas    *int                                                              `json:"desiredReplicas"`

	// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control
	// how often the number of pods is changed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	LastScaleTime      *time.Time                                                        `json:"lastScaleTime,omitempty"`

	// observedGeneration is the most recent generation observed by this autoscaler.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	ObservedGeneration *int                                                              `json:"observedGeneration,omitempty"`
}

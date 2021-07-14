package model

import (
	"time"
)


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v1_HorizontalPodAutoscaler.go


// current status of a horizontal pod autoscaler
type Io_k8s_api_autoscaling_v1_HorizontalPodAutoscalerStatus struct {
	// current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an
	// average pod is using now 70% of its requested CPU.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CurrentCPUUtilizationPercentage *int       `json:"currentCPUUtilizationPercentage,omitempty"`

	// current number of replicas of pods managed by this autoscaler.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	CurrentReplicas                 *int       `json:"currentReplicas"`

	// desired number of replicas of pods managed by this autoscaler.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	DesiredReplicas                 *int       `json:"desiredReplicas"`

	// last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number
	// of pods is changed.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastScaleTime                   *time.Time `json:"lastScaleTime,omitempty"`

	// most recent generation observed by this autoscaler.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ObservedGeneration              *int       `json:"observedGeneration,omitempty"`
}

package kugo_model

import (
	"time"
)

type Io_k8s_api_autoscaling_v1_HorizontalPodAutoscalerStatus struct {
	CurrentCPUUtilizationPercentage int        `json:"currentCPUUtilizationPercentage,omitempty"`
	CurrentReplicas                 int        `json:"currentReplicas"`
	DesiredReplicas                 int        `json:"desiredReplicas"`
	LastScaleTime                   *time.Time `json:"lastScaleTime,omitempty"`
	ObservedGeneration              int        `json:"observedGeneration,omitempty"`
}


package kugo_model

import (
	"time"
)

type Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerStatus struct {
	Conditions         []Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerCondition `json:"conditions"`
	CurrentMetrics     []Io_k8s_api_autoscaling_v2beta1_MetricStatus                     `json:"currentMetrics,omitempty"`
	CurrentReplicas    int                                                               `json:"currentReplicas"`
	DesiredReplicas    int                                                               `json:"desiredReplicas"`
	LastScaleTime      *time.Time                                                        `json:"lastScaleTime,omitempty"`
	ObservedGeneration int                                                               `json:"observedGeneration,omitempty"`
}


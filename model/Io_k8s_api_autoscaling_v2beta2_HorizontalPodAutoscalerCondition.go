package kugo_model

import (
	"time"
)

type Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerCondition struct {
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	Reason             string     `json:"reason,omitempty"`
	Status             string     `json:"status"`
	Type               string     `json:"type"`
}


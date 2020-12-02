package kugo_model

import (
	"time"
)

type Io_k8s_api_core_v1_PodCondition struct {
	LastProbeTime      *time.Time `json:"lastProbeTime,omitempty"`
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	Reason             string     `json:"reason,omitempty"`
	Status             string     `json:"status"`
	Type               string     `json:"type"`
}


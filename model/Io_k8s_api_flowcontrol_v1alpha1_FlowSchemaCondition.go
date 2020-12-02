package kugo_model

import (
	"time"
)

type Io_k8s_api_flowcontrol_v1alpha1_FlowSchemaCondition struct {
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	Reason             string     `json:"reason,omitempty"`
	Status             string     `json:"status,omitempty"`
	Type               string     `json:"type,omitempty"`
}


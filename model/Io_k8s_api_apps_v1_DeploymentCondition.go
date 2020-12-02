package kugo_model

import (
	"time"
)

type Io_k8s_api_apps_v1_DeploymentCondition struct {
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	LastUpdateTime     *time.Time `json:"lastUpdateTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	Reason             string     `json:"reason,omitempty"`
	Status             string     `json:"status"`
	Type               string     `json:"type"`
}


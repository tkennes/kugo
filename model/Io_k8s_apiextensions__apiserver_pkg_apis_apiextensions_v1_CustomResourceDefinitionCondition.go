package kugo_model

import (
	"time"
)

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionCondition struct {
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	Reason             string     `json:"reason,omitempty"`
	Status             string     `json:"status"`
	Type               string     `json:"type"`
}


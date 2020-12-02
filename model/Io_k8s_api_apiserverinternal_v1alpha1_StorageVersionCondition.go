package kugo_model

import (
	"time"
)

type Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionCondition struct {
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	ObservedGeneration int        `json:"observedGeneration,omitempty"`
	Reason             string     `json:"reason"`
	Status             string     `json:"status"`
	Type               string     `json:"type"`
}


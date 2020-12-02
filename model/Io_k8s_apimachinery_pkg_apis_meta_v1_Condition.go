package kugo_model

import (
	"time"
)

type Io_k8s_apimachinery_pkg_apis_meta_v1_Condition struct {
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	Message            string    `json:"message"`
	ObservedGeneration int       `json:"observedGeneration,omitempty"`
	Reason             string    `json:"reason"`
	Status             string    `json:"status"`
	Type               string    `json:"type"`
}


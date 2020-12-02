package kugo_model

import (
	"time"
)

type Io_k8s_api_certificates_v1beta1_CertificateSigningRequestCondition struct {
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	LastUpdateTime     *time.Time `json:"lastUpdateTime,omitempty"`
	Message            string     `json:"message,omitempty"`
	Reason             string     `json:"reason,omitempty"`
	Status             string     `json:"status,omitempty"`
	Type               string     `json:"type"`
}


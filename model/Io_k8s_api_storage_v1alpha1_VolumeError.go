package kugo_model

import (
	"time"
)

type Io_k8s_api_storage_v1alpha1_VolumeError struct {
	Message string     `json:"message,omitempty"`
	Time    *time.Time `json:"time,omitempty"`
}


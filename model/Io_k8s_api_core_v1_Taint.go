package kugo_model

import (
	"time"
)

type Io_k8s_api_core_v1_Taint struct {
	Effect    string     `json:"effect"`
	Key       string     `json:"key"`
	TimeAdded *time.Time `json:"timeAdded,omitempty"`
	Value     string     `json:"value,omitempty"`
}


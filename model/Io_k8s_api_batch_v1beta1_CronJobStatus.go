package kugo_model

import (
	"time"
)

type Io_k8s_api_batch_v1beta1_CronJobStatus struct {
	Active           []Io_k8s_api_core_v1_ObjectReference `json:"active,omitempty"`
	LastScheduleTime *time.Time                           `json:"lastScheduleTime,omitempty"`
}


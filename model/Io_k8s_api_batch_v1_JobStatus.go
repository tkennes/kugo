package kugo_model

import (
	"time"
)

type Io_k8s_api_batch_v1_JobStatus struct {
	Active         int                                `json:"active,omitempty"`
	CompletionTime *time.Time                         `json:"completionTime,omitempty"`
	Conditions     []Io_k8s_api_batch_v1_JobCondition `json:"conditions,omitempty"`
	Failed         int                                `json:"failed,omitempty"`
	StartTime      *time.Time                         `json:"startTime,omitempty"`
	Succeeded      int                                `json:"succeeded,omitempty"`
}


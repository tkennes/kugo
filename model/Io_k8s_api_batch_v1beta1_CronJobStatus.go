package model

import (
	"time"
)


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v1beta1_CronJob.go


// CronJobStatus represents the current state of a cron job.
type Io_k8s_api_batch_v1beta1_CronJobStatus struct {
	// A list of pointers to currently running jobs.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectReference.go
	Active           []Io_k8s_api_core_v1_ObjectReference `json:"active,omitempty"`

	// Information when was the last time the job was successfully scheduled.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	LastScheduleTime *time.Time                           `json:"lastScheduleTime,omitempty"`
}

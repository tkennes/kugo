package kugo_model

import (
	"time"
)


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v1_Job.go


// JobStatus represents the current state of a Job.
type Io_k8s_api_batch_v1_JobStatus struct {
	// The number of actively running pods.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Active         *int                               `json:"active,omitempty"`

	// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate
	// operations. It is represented in RFC3339 form and is in UTC. The completion time is only set when the job finishes
	// successfully.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	CompletionTime *time.Time                         `json:"completionTime,omitempty"`

	// The latest available observations of an object's current state. When a job fails, one of the conditions will have type
	// == "Failed". More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v1_JobCondition.go
	Conditions     []Io_k8s_api_batch_v1_JobCondition `json:"conditions,omitempty"`

	// The number of pods which reached phase Failed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Failed         *int                               `json:"failed,omitempty"`

	// Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before
	// order across separate operations. It is represented in RFC3339 form and is in UTC.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/time.Time.go
	StartTime      *time.Time                         `json:"startTime,omitempty"`

	// The number of pods which reached phase Succeeded.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Succeeded      *int                               `json:"succeeded,omitempty"`
}

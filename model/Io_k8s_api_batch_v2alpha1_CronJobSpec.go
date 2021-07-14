package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v2alpha1_CronJob.go


// CronJobSpec describes how the job execution will look like and when it will actually run.
type Io_k8s_api_batch_v2alpha1_CronJobSpec struct {
	// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run
	// concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace":
	// cancels currently running job and replaces it with a new one
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ConcurrencyPolicy          *string                                   `json:"concurrencyPolicy,omitempty"`

	// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	FailedJobsHistoryLimit     *int                                      `json:"failedJobsHistoryLimit,omitempty"`

	// Specifies the job that will be created when executing a CronJob.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v2alpha1_JobTemplateSpec.go
	JobTemplate                Io_k8s_api_batch_v2alpha1_JobTemplateSpec `json:"jobTemplate"`

	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Schedule                   *string                                   `json:"schedule"`

	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions
	// will be counted as failed ones.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	StartingDeadlineSeconds    *int                                      `json:"startingDeadlineSeconds,omitempty"`

	// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not
	// specified.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	SuccessfulJobsHistoryLimit *int                                      `json:"successfulJobsHistoryLimit,omitempty"`

	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.
	// Defaults to false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Suspend                    *bool                                     `json:"suspend,omitempty"`
}

package kugo_model

type Io_k8s_api_batch_v1beta1_CronJobSpec struct {
	ConcurrencyPolicy          string                                   `json:"concurrencyPolicy,omitempty"`
	FailedJobsHistoryLimit     int                                      `json:"failedJobsHistoryLimit,omitempty"`
	JobTemplate                Io_k8s_api_batch_v1beta1_JobTemplateSpec `json:"jobTemplate"`
	Schedule                   string                                   `json:"schedule"`
	StartingDeadlineSeconds    int                                      `json:"startingDeadlineSeconds,omitempty"`
	SuccessfulJobsHistoryLimit int                                      `json:"successfulJobsHistoryLimit,omitempty"`
	Suspend                    bool                                     `json:"suspend,omitempty"`
}


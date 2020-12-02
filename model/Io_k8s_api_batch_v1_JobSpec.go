package kugo_model

type Io_k8s_api_batch_v1_JobSpec struct {
	ActiveDeadlineSeconds   int                                                 `json:"activeDeadlineSeconds,omitempty"`
	BackoffLimit            int                                                 `json:"backoffLimit,omitempty"`
	Completions             int                                                 `json:"completions,omitempty"`
	ManualSelector          bool                                                `json:"manualSelector,omitempty"`
	Parallelism             int                                                 `json:"parallelism,omitempty"`
	Selector                *Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector,omitempty"`
	Template                Io_k8s_api_core_v1_PodTemplateSpec                  `json:"template"`
	TtlSecondsAfterFinished int                                                 `json:"ttlSecondsAfterFinished,omitempty"`
}


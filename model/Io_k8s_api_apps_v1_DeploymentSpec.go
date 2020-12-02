package kugo_model

type Io_k8s_api_apps_v1_DeploymentSpec struct {
	MinReadySeconds         int                                                `json:"minReadySeconds,omitempty"`
	Paused                  bool                                               `json:"paused,omitempty"`
	ProgressDeadlineSeconds int                                                `json:"progressDeadlineSeconds,omitempty"`
	Replicas                int                                                `json:"replicas,omitempty"`
	RevisionHistoryLimit    int                                                `json:"revisionHistoryLimit,omitempty"`
	Selector                Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector"`
	Strategy                *Io_k8s_api_apps_v1_DeploymentStrategy             `json:"strategy,omitempty"`
	Template                Io_k8s_api_core_v1_PodTemplateSpec                 `json:"template"`
}


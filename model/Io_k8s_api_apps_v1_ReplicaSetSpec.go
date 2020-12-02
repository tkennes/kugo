package kugo_model

type Io_k8s_api_apps_v1_ReplicaSetSpec struct {
	MinReadySeconds int                                                `json:"minReadySeconds,omitempty"`
	Replicas        int                                                `json:"replicas,omitempty"`
	Selector        Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector"`
	Template        *Io_k8s_api_core_v1_PodTemplateSpec                `json:"template,omitempty"`
}


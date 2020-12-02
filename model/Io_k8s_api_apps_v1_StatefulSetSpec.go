package kugo_model

type Io_k8s_api_apps_v1_StatefulSetSpec struct {
	PodManagementPolicy  string                                             `json:"podManagementPolicy,omitempty"`
	Replicas             int                                                `json:"replicas,omitempty"`
	RevisionHistoryLimit int                                                `json:"revisionHistoryLimit,omitempty"`
	Selector             Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector"`
	ServiceName          string                                             `json:"serviceName"`
	Template             Io_k8s_api_core_v1_PodTemplateSpec                 `json:"template"`
	UpdateStrategy       *Io_k8s_api_apps_v1_StatefulSetUpdateStrategy      `json:"updateStrategy,omitempty"`
	VolumeClaimTemplates []Io_k8s_api_core_v1_PersistentVolumeClaim         `json:"volumeClaimTemplates,omitempty"`
}


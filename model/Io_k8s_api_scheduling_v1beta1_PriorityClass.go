package kugo_model

type Io_k8s_api_scheduling_v1beta1_PriorityClass struct {
	ApiVersion       string                                           `json:"apiVersion,omitempty"`
	Description      string                                           `json:"description,omitempty"`
	GlobalDefault    bool                                             `json:"globalDefault,omitempty"`
	Kind             string                                           `json:"kind,omitempty"`
	Metadata         *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	PreemptionPolicy string                                           `json:"preemptionPolicy,omitempty"`
	Value            int                                              `json:"value"`
}


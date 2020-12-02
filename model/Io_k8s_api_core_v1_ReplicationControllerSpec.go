package kugo_model

type Io_k8s_api_core_v1_ReplicationControllerSpec struct {
	MinReadySeconds int                                 `json:"minReadySeconds,omitempty"`
	Replicas        int                                 `json:"replicas,omitempty"`
	Selector        *interface{}                        `json:"selector,omitempty"`
	Template        *Io_k8s_api_core_v1_PodTemplateSpec `json:"template,omitempty"`
}


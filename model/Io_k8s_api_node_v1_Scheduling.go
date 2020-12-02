package kugo_model

type Io_k8s_api_node_v1_Scheduling struct {
	NodeSelector *interface{}                    `json:"nodeSelector,omitempty"`
	Tolerations  []Io_k8s_api_core_v1_Toleration `json:"tolerations,omitempty"`
}


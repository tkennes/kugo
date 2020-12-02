package kugo_model

type Io_k8s_api_core_v1_Toleration struct {
	Effect            string `json:"effect,omitempty"`
	Key               string `json:"key,omitempty"`
	Operator          string `json:"operator,omitempty"`
	TolerationSeconds int    `json:"tolerationSeconds,omitempty"`
	Value             string `json:"value,omitempty"`
}


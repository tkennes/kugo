package kugo_model

type Io_k8s_api_discovery_v1beta1_EndpointConditions struct {
	Ready       bool `json:"ready,omitempty"`
	Serving     bool `json:"serving,omitempty"`
	Terminating bool `json:"terminating,omitempty"`
}


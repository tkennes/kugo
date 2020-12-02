package kugo_model

type Io_k8s_api_core_v1_EndpointAddress struct {
	Hostname  string                              `json:"hostname,omitempty"`
	Ip        string                              `json:"ip"`
	NodeName  string                              `json:"nodeName,omitempty"`
	TargetRef *Io_k8s_api_core_v1_ObjectReference `json:"targetRef,omitempty"`
}


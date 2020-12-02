package kugo_model

type Io_k8s_api_discovery_v1beta1_Endpoint struct {
	Addresses  []string                                         `json:"addresses"`
	Conditions *Io_k8s_api_discovery_v1beta1_EndpointConditions `json:"conditions,omitempty"`
	Hostname   string                                           `json:"hostname,omitempty"`
	NodeName   string                                           `json:"nodeName,omitempty"`
	TargetRef  *Io_k8s_api_core_v1_ObjectReference              `json:"targetRef,omitempty"`
	Topology   *interface{}                                     `json:"topology,omitempty"`
}


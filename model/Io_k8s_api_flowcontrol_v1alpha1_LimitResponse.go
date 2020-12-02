package kugo_model

type Io_k8s_api_flowcontrol_v1alpha1_LimitResponse struct {
	Queuing *Io_k8s_api_flowcontrol_v1alpha1_QueuingConfiguration `json:"queuing,omitempty"`
	Type    string                                                `json:"type"`
}


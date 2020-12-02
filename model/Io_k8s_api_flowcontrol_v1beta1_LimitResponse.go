package kugo_model

type Io_k8s_api_flowcontrol_v1beta1_LimitResponse struct {
	Queuing *Io_k8s_api_flowcontrol_v1beta1_QueuingConfiguration `json:"queuing,omitempty"`
	Type    string                                               `json:"type"`
}


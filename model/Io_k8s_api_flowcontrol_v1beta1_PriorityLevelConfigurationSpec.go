package kugo_model

type Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfigurationSpec struct {
	Limited *Io_k8s_api_flowcontrol_v1beta1_LimitedPriorityLevelConfiguration `json:"limited,omitempty"`
	Type    string                                                            `json:"type"`
}


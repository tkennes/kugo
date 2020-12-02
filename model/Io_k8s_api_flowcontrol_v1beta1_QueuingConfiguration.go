package kugo_model

type Io_k8s_api_flowcontrol_v1beta1_QueuingConfiguration struct {
	HandSize         int `json:"handSize,omitempty"`
	QueueLengthLimit int `json:"queueLengthLimit,omitempty"`
	Queues           int `json:"queues,omitempty"`
}


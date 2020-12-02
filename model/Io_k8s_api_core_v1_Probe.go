package kugo_model

type Io_k8s_api_core_v1_Probe struct {
	Exec                *Io_k8s_api_core_v1_ExecAction      `json:"exec,omitempty"`
	FailureThreshold    int                                 `json:"failureThreshold,omitempty"`
	HttpGet             *Io_k8s_api_core_v1_HTTPGetAction   `json:"httpGet,omitempty"`
	InitialDelaySeconds int                                 `json:"initialDelaySeconds,omitempty"`
	PeriodSeconds       int                                 `json:"periodSeconds,omitempty"`
	SuccessThreshold    int                                 `json:"successThreshold,omitempty"`
	TcpSocket           *Io_k8s_api_core_v1_TCPSocketAction `json:"tcpSocket,omitempty"`
	TimeoutSeconds      int                                 `json:"timeoutSeconds,omitempty"`
}


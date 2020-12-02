package kugo_model

type Io_k8s_api_core_v1_Handler struct {
	Exec      *Io_k8s_api_core_v1_ExecAction      `json:"exec,omitempty"`
	HttpGet   *Io_k8s_api_core_v1_HTTPGetAction   `json:"httpGet,omitempty"`
	TcpSocket *Io_k8s_api_core_v1_TCPSocketAction `json:"tcpSocket,omitempty"`
}


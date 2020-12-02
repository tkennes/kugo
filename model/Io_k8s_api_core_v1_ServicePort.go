package kugo_model

type Io_k8s_api_core_v1_ServicePort struct {
	AppProtocol string `json:"appProtocol,omitempty"`
	Name        string `json:"name,omitempty"`
	NodePort    int    `json:"nodePort,omitempty"`
	Port        int    `json:"port"`
	Protocol    string `json:"protocol,omitempty"`
	TargetPort  int    `json:"targetPort,omitempty"`
}


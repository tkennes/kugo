package kugo_model

type Io_k8s_api_networking_v1_IngressServiceBackend struct {
	Name string                                       `json:"name"`
	Port *Io_k8s_api_networking_v1_ServiceBackendPort `json:"port,omitempty"`
}


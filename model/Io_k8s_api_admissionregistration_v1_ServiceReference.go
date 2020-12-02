package kugo_model

type Io_k8s_api_admissionregistration_v1_ServiceReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Path      string `json:"path,omitempty"`
	Port      int    `json:"port,omitempty"`
}


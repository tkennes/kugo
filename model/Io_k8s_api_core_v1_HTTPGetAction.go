package kugo_model

type Io_k8s_api_core_v1_HTTPGetAction struct {
	Host        string                          `json:"host,omitempty"`
	HttpHeaders []Io_k8s_api_core_v1_HTTPHeader `json:"httpHeaders,omitempty"`
	Path        string                          `json:"path,omitempty"`
	Port        int                             `json:"port"`
	Scheme      string                          `json:"scheme,omitempty"`
}


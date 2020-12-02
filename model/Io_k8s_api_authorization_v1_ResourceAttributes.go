package kugo_model

type Io_k8s_api_authorization_v1_ResourceAttributes struct {
	Group       string `json:"group,omitempty"`
	Name        string `json:"name,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
	Resource    string `json:"resource,omitempty"`
	Subresource string `json:"subresource,omitempty"`
	Verb        string `json:"verb,omitempty"`
	Version     string `json:"version,omitempty"`
}


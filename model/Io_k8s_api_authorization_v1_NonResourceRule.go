package kugo_model

type Io_k8s_api_authorization_v1_NonResourceRule struct {
	NonResourceURLs []string `json:"nonResourceURLs,omitempty"`
	Verbs           []string `json:"verbs"`
}


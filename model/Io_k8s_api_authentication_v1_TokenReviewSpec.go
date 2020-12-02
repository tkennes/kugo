package kugo_model

type Io_k8s_api_authentication_v1_TokenReviewSpec struct {
	Audiences []string `json:"audiences,omitempty"`
	Token     string   `json:"token,omitempty"`
}


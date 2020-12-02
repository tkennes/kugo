package kugo_model

type Io_k8s_api_authentication_v1_TokenReviewStatus struct {
	Audiences     []string                               `json:"audiences,omitempty"`
	Authenticated bool                                   `json:"authenticated,omitempty"`
	Error         string                                 `json:"error,omitempty"`
	User          *Io_k8s_api_authentication_v1_UserInfo `json:"user,omitempty"`
}


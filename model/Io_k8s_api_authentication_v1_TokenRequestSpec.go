package kugo_model

type Io_k8s_api_authentication_v1_TokenRequestSpec struct {
	Audiences         []string                                           `json:"audiences"`
	BoundObjectRef    *Io_k8s_api_authentication_v1_BoundObjectReference `json:"boundObjectRef,omitempty"`
	ExpirationSeconds int                                                `json:"expirationSeconds,omitempty"`
}


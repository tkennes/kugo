package kugo_model

type Io_k8s_api_networking_v1_IPBlock struct {
	Cidr   string   `json:"cidr"`
	Except []string `json:"except,omitempty"`
}


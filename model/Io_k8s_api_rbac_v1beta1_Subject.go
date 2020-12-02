package kugo_model

type Io_k8s_api_rbac_v1beta1_Subject struct {
	ApiGroup  string `json:"apiGroup,omitempty"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}


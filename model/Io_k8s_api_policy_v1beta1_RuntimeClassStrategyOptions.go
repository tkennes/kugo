package kugo_model

type Io_k8s_api_policy_v1beta1_RuntimeClassStrategyOptions struct {
	AllowedRuntimeClassNames []string `json:"allowedRuntimeClassNames"`
	DefaultRuntimeClassName  string   `json:"defaultRuntimeClassName,omitempty"`
}


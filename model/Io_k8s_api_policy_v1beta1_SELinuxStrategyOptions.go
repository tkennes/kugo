package kugo_model

type Io_k8s_api_policy_v1beta1_SELinuxStrategyOptions struct {
	Rule           string                             `json:"rule"`
	SeLinuxOptions *Io_k8s_api_core_v1_SELinuxOptions `json:"seLinuxOptions,omitempty"`
}


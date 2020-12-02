package kugo_model

type Io_k8s_api_networking_v1_NetworkPolicyIngressRule struct {
	From  []Io_k8s_api_networking_v1_NetworkPolicyPeer `json:"from,omitempty"`
	Ports []Io_k8s_api_networking_v1_NetworkPolicyPort `json:"ports,omitempty"`
}


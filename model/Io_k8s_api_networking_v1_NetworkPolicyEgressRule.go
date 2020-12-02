package kugo_model

type Io_k8s_api_networking_v1_NetworkPolicyEgressRule struct {
	Ports []Io_k8s_api_networking_v1_NetworkPolicyPort `json:"ports,omitempty"`
	To    []Io_k8s_api_networking_v1_NetworkPolicyPeer `json:"to,omitempty"`
}


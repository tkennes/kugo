package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSpec.go


// Affinity is a group of affinity scheduling rules.
type Io_k8s_api_core_v1_Affinity struct {
	// Describes node affinity scheduling rules for the pod.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeAffinity.go
	NodeAffinity    *Io_k8s_api_core_v1_NodeAffinity    `json:"nodeAffinity,omitempty"`

	// Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodAffinity.go
	PodAffinity     *Io_k8s_api_core_v1_PodAffinity     `json:"podAffinity,omitempty"`

	// Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other
	// pod(s)).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodAntiAffinity.go
	PodAntiAffinity *Io_k8s_api_core_v1_PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

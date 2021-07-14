package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodDNSConfig.go


// PodDNSConfigOption defines DNS resolver options of a pod.
type Io_k8s_api_core_v1_PodDNSConfigOption struct {
	// Required.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name  *string `json:"name,omitempty"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Value *string `json:"value,omitempty"`
}

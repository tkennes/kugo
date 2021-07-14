package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ServiceSpec.go


// SessionAffinityConfig represents the configurations of session affinity.
type Io_k8s_api_core_v1_SessionAffinityConfig struct {
	// clientIP contains the configurations of Client IP based session affinity.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ClientIPConfig.go
	ClientIP *Io_k8s_api_core_v1_ClientIPConfig `json:"clientIP,omitempty"`
}

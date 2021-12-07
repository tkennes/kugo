package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SessionAffinityConfig.go


// ClientIPConfig represents the configurations of Client IP based session affinity.
type Io_k8s_api_core_v1_ClientIPConfig struct {
	// timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if
	// ServiceAffinity == "ClientIP". Default value is 10800(for 3 hours).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DaemonSetSpec.go


// DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
type Io_k8s_api_apps_v1_DaemonSetUpdateStrategy struct {
	// Rolling update config params. Present only if type = "RollingUpdate".
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_RollingUpdateDaemonSet.go
	RollingUpdate *Io_k8s_api_apps_v1_RollingUpdateDaemonSet `json:"rollingUpdate,omitempty"`

	// Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type          *string                                    `json:"type,omitempty"`
}

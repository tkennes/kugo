package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DeploymentSpec.go


// DeploymentStrategy describes how to replace existing pods with new ones.
type Io_k8s_api_apps_v1_DeploymentStrategy struct {
	// Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_RollingUpdateDeployment.go
	RollingUpdate *Io_k8s_api_apps_v1_RollingUpdateDeployment `json:"rollingUpdate,omitempty"`

	// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type          *string                                     `json:"type,omitempty"`
}

package kugo_model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EnvVarSource.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_DownwardAPIVolumeFile.go


// ResourceFieldSelector represents container resources (cpu, memory) and their output format
type Io_k8s_api_core_v1_ResourceFieldSelector struct {
	// Container name: required for volumes, optional for env vars
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ContainerName *string                                        `json:"containerName,omitempty"`

	// Specifies the output format of the exposed resources, defaults to "1"
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_api_resource_Quantity.go
	Divisor       *Io_k8s_apimachinery_pkg_api_resource_Quantity `json:"divisor,omitempty"`

	// Required: resource to select
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Resource      *string                                        `json:"resource"`
}

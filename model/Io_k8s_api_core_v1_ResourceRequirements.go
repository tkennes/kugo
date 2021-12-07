package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Container.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go


// ResourceRequirements describes the compute resource requirements.
type Io_k8s_api_core_v1_ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed. More info:
	// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	Limits   *interface{} `json:"limits,omitempty"`

	// Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults
	// to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info:
	// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	Requests *interface{} `json:"requests,omitempty"`
}

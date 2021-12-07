package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeStatus.go


// AttachedVolume describes a volume attached to a node
type Io_k8s_api_core_v1_AttachedVolume struct {
	// DevicePath represents the device path where the volume should be available
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	DevicePath *string `json:"devicePath"`

	// Name of the attached volume
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name       *string `json:"name"`
}

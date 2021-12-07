package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Container.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go


// volumeDevice describes a mapping of a raw block device within a container.
type Io_k8s_api_core_v1_VolumeDevice struct {
	// devicePath is the path inside of the container that the device will be mapped to.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	DevicePath *string `json:"devicePath"`

	// name must match the name of a persistentVolumeClaim in the pod
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name       *string `json:"name"`
}

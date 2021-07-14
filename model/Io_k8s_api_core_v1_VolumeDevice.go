package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Container.go


// volumeDevice describes a mapping of a raw block device within a container.
type Io_k8s_api_core_v1_VolumeDevice struct {
	// devicePath is the path inside of the container that the device will be mapped to.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	DevicePath *string `json:"devicePath"`

	// name must match the name of a persistentVolumeClaim in the pod
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name       *string `json:"name"`
}

package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.
type Io_k8s_api_core_v1_HostPathVolumeSource struct {
	// Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info:
	// https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path *string `json:"path"`

	// Type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type *string `json:"type,omitempty"`
}

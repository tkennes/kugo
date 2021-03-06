package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_DownwardAPIVolumeSource.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_DownwardAPIProjection.go


// DownwardAPIVolumeFile represents information to create the file containing the pod field
type Io_k8s_api_core_v1_DownwardAPIVolumeFile struct {
	// Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ObjectFieldSelector.go
	FieldRef         *Io_k8s_api_core_v1_ObjectFieldSelector   `json:"fieldRef,omitempty"`

	// Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal
	// value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not
	// specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode,
	// like fsGroup, and the result can be other mode bits set.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Mode             *int                                      `json:"mode,omitempty"`

	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must
	// be utf-8 encoded. The first item of the relative path must not start with '..'
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Path             *string                                   `json:"path"`

	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and
	// requests.memory) are currently supported.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ResourceFieldSelector.go
	ResourceFieldRef *Io_k8s_api_core_v1_ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}

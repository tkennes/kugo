package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecretVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ConfigMapVolumeSource.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ConfigMapProjection.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecretProjection.go


// Maps a string key to a path within a volume.
type Io_k8s_api_core_v1_KeyToPath struct {
	// The key to project.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Key  *string `json:"key"`

	// Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal
	// value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not
	// specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode,
	// like fsGroup, and the result can be other mode bits set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Mode *int    `json:"mode,omitempty"`

	// The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May
	// not start with the string '..'.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Path *string `json:"path"`
}

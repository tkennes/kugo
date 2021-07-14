package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a projected volume source
type Io_k8s_api_core_v1_ProjectedVolumeSource struct {
	// Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal
	// value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// Directories within the path are not affected by this setting. This might be in conflict with other options that affect
	// the file mode, like fsGroup, and the result can be other mode bits set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	DefaultMode *int                                  `json:"defaultMode,omitempty"`

	// list of volume projections
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_VolumeProjection.go
	Sources     []Io_k8s_api_core_v1_VolumeProjection `json:"sources,omitempty"`
}

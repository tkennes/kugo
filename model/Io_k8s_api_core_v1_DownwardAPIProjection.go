package kugo_model


// Tree Depth: 6
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_VolumeProjection.go


// Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume
// source without the default mode.
type Io_k8s_api_core_v1_DownwardAPIProjection struct {
	// Items is a list of DownwardAPIVolume file
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_DownwardAPIVolumeFile.go
	Items []Io_k8s_api_core_v1_DownwardAPIVolumeFile `json:"items,omitempty"`
}

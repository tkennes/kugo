package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go


// VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
type Io_k8s_api_core_v1_VolumeNodeAffinity struct {
	// Required specifies hard node constraints that must be met.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeSelector.go
	Required *Io_k8s_api_core_v1_NodeSelector `json:"required,omitempty"`
}

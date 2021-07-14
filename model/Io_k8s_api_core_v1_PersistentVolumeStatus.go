package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolume.go


// PersistentVolumeStatus is the current status of a persistent volume.
type Io_k8s_api_core_v1_PersistentVolumeStatus struct {
	// A human-readable message indicating details about why the volume is in this state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message *string `json:"message,omitempty"`

	// Phase indicates if a volume is available, bound to a claim, or released by a claim. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Phase   *string `json:"phase,omitempty"`

	// Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the
	// CLI.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason  *string `json:"reason,omitempty"`
}

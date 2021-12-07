package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaim.go


// PersistentVolumeClaimStatus is the current status of a persistent volume claim.
type Io_k8s_api_core_v1_PersistentVolumeClaimStatus struct {
	// AccessModes contains the actual access modes the volume backing the PVC has. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	AccessModes []*string                                           `json:"accessModes,omitempty"`

	// Represents the actual resources of the underlying volume.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/interface{}.go
	Capacity    *interface{}                                        `json:"capacity,omitempty"`

	// Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will
	// be set to 'ResizeStarted'.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimCondition.go
	Conditions  []Io_k8s_api_core_v1_PersistentVolumeClaimCondition `json:"conditions,omitempty"`

	// Phase represents the current phase of PersistentVolumeClaim.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Phase       *string                                             `json:"phase,omitempty"`
}

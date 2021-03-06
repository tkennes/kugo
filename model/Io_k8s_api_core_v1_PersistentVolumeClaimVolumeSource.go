package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and
// mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of
// volume that is owned by someone else (the system).
type Io_k8s_api_core_v1_PersistentVolumeClaimVolumeSource struct {
	// ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ClaimName *string `json:"claimName"`

	// Will force the ReadOnly setting in VolumeMounts. Default false.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}

package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents an ephemeral volume that is handled by a normal storage driver.
type Io_k8s_api_core_v1_EphemeralVolumeSource struct {
	// Specifies a read-only configuration for the volume. Defaults to false (read/write).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnly            *bool                                             `json:"readOnly,omitempty"`

	// Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is
	// embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be
	// `<pod name>-<volume name>` where `<volume name>` is the name from the `PodSpec.Volumes` array entry. Pod validation will
	// reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name
	// that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the
	// pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the
	// PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it
	// may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by
	// Kubernetes to the PVC after it has been created.  Required, must not be nil.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimTemplate.go
	VolumeClaimTemplate *Io_k8s_api_core_v1_PersistentVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty"`
}

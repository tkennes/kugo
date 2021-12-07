package kugo_model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EphemeralVolumeSource.go


// PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
type Io_k8s_api_core_v1_PersistentVolumeClaimTemplate struct {
	// May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and
	// will be rejected during validation.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created
	// from this template. The same fields as in a PersistentVolumeClaim are also valid here.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimSpec.go
	Spec     Io_k8s_api_core_v1_PersistentVolumeClaimSpec     `json:"spec"`
}

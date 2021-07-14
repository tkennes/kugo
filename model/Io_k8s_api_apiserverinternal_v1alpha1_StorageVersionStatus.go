package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersion.go


// API server instances report the versions they can decode and the version they encode objects to when persisting objects
// in the backend.
type Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionStatus struct {
	// If all API server instances agree on the same encoding storage version, then this field is set to that version.
	// Otherwise this field is left empty. API servers should finish updating its storageVersionStatus entry before serving
	// write operations, so that this field will be in sync with the reality.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	CommonEncodingVersion *string                                                         `json:"commonEncodingVersion,omitempty"`

	// The latest available observations of the storageVersion's state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionCondition.go
	Conditions            []Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionCondition `json:"conditions,omitempty"`

	// The reported versions per API server instance.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_ServerStorageVersion.go
	StorageVersions       []Io_k8s_api_apiserverinternal_v1alpha1_ServerStorageVersion    `json:"storageVersions,omitempty"`
}

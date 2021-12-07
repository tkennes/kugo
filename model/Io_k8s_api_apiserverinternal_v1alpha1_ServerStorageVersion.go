package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionStatus.go


// An API server instance reports the version it can decode and the version it encodes objects to when persisting objects
// in the backend.
type Io_k8s_api_apiserverinternal_v1alpha1_ServerStorageVersion struct {
	// The ID of the reporting API server.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiServerID       *string   `json:"apiServerID,omitempty"`

	// The API server can decode objects encoded in these versions. The encodingVersion must be included in the
	// decodableVersions.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	DecodableVersions []*string `json:"decodableVersions,omitempty"`

	// The API server encodes the object to this version when persisting it in the backend (e.g., etcd).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	EncodingVersion   *string   `json:"encodingVersion,omitempty"`
}

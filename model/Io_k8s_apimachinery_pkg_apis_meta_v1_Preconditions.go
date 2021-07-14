package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_DeleteOptions.go


// Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
type Io_k8s_apimachinery_pkg_apis_meta_v1_Preconditions struct {
	// Specifies the target ResourceVersion
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ResourceVersion *string `json:"resourceVersion,omitempty"`

	// Specifies the target UID.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Uid             *string `json:"uid,omitempty"`
}

package kugo_model


// Event represents a single event to a watched resource.
type Io_k8s_apimachinery_pkg_apis_meta_v1_WatchEvent struct {
	// Object is:  * If Type is Added or Modified: the new state of the object.  * If Type is Deleted: the state of the object
	// immediately before deletion.  * If Type is Error: *Status is recommended; other types may make sense    depending on
	// context.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_runtime_RawExtension.go
	Object *Io_k8s_apimachinery_pkg_runtime_RawExtension `json:"object"`

	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Type   *string                                       `json:"type"`
}

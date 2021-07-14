package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_ControllerRevisionList.go


// ControllerRevision implements an immutable snapshot of state data. Clients are responsible for serializing and
// deserializing the objects that contain their internal state. Once a ControllerRevision has been successfully created, it
// can not be updated. The API Server will fail validation of all requests that attempt to mutate the Data field.
// ControllerRevisions may, however, be deleted. Note that, due to its use by both the DaemonSet and StatefulSet
// controllers for update and rollback, this object is beta. However, it may be subject to name and representation changes
// in future releases, and clients should not depend on its stability. It is primarily for internal use by controllers.
type Io_k8s_api_apps_v1_ControllerRevision struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                          `json:"apiVersion,omitempty"`

	// Data is the serialized representation of the state.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_runtime_RawExtension.go
	Data       *Io_k8s_apimachinery_pkg_runtime_RawExtension    `json:"data,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                          `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Revision indicates the revision of the state represented by Data.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Revision   *int                                             `json:"revision"`
}

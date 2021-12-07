package kugo_model


// Scale represents a scaling request for a resource.
type Io_k8s_api_autoscaling_v1_Scale struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                          `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                          `json:"kind,omitempty"`

	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#spec-and-status.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v1_ScaleSpec.go
	Spec       *Io_k8s_api_autoscaling_v1_ScaleSpec             `json:"spec,omitempty"`

	// current status of the scale. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#spec-and-status. Read-only.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v1_ScaleStatus.go
	Status     *Io_k8s_api_autoscaling_v1_ScaleStatus           `json:"status,omitempty"`
}

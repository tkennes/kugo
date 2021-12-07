package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerList.go


// HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica
// count of any resource implementing the scale subresource based on the metrics specified.
type Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscaler struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                                       `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                                       `json:"kind,omitempty"`

	// metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-
	// architecture/api-conventions.md#metadata
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta              `json:"metadata,omitempty"`

	// spec is the specification for the behaviour of the autoscaler. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerSpec.go
	Spec       *Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerSpec   `json:"spec,omitempty"`

	// status is the current information about the autoscaler.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerStatus.go
	Status     *Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

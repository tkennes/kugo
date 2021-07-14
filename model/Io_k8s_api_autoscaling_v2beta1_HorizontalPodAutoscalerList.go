package kugo_model


// HorizontalPodAutoscaler is a list of horizontal pod autoscaler objects.
type Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                                  `json:"apiVersion,omitempty"`

	// items is the list of horizontal pod autoscaler objects.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscaler.go
	Items      []Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscaler `json:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                                  `json:"kind,omitempty"`

	// metadata is the standard list metadata.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta           `json:"metadata,omitempty"`
}

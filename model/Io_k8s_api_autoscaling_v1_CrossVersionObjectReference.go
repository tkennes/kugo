package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v1_HorizontalPodAutoscalerSpec.go


// CrossVersionObjectReference contains enough information to let you identify the referred resource.
type Io_k8s_api_autoscaling_v1_CrossVersionObjectReference struct {
	// API version of the referent
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string `json:"apiVersion,omitempty"`

	// Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#types-kinds"
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string `json:"kind"`

	// Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name       *string `json:"name"`
}

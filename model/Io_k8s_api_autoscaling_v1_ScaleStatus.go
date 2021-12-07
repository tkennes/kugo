package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v1_Scale.go


// ScaleStatus represents the current status of a scale subresource.
type Io_k8s_api_autoscaling_v1_ScaleStatus struct {
	// actual number of observed instances of the scaled object.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Replicas *int    `json:"replicas"`

	// label query over pods that should match the replicas count. This is same as the label selector but in the string format
	// to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about
	// label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Selector *string `json:"selector,omitempty"`
}

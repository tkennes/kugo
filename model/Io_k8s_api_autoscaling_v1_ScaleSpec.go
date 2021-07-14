package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v1_Scale.go


// ScaleSpec describes the attributes of a scale subresource.
type Io_k8s_api_autoscaling_v1_ScaleSpec struct {
	// desired number of instances for the scaled object.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Replicas *int `json:"replicas,omitempty"`
}

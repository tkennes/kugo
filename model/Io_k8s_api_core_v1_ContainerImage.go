package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeStatus.go


// Describe a container image
type Io_k8s_api_core_v1_ContainerImage struct {
	// Names by which this image is known. e.g. ["k8s.gcr.io/hyperkube:v1.0.7",
	// "dockerhub.io/google_containers/hyperkube:v1.0.7"]
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Names     []*string `json:"names"`

	// The size of the image in bytes.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	SizeBytes *int      `json:"sizeBytes,omitempty"`
}

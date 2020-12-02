package kugo_model

type Io_k8s_api_node_v1alpha1_RuntimeClassSpec struct {
	Overhead       *Io_k8s_api_node_v1alpha1_Overhead   `json:"overhead,omitempty"`
	RuntimeHandler string                               `json:"runtimeHandler"`
	Scheduling     *Io_k8s_api_node_v1alpha1_Scheduling `json:"scheduling,omitempty"`
}


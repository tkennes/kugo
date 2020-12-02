package kugo_model

type Io_k8s_api_node_v1beta1_RuntimeClass struct {
	ApiVersion string                                           `json:"apiVersion,omitempty"`
	Handler    string                                           `json:"handler"`
	Kind       string                                           `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`
	Overhead   *Io_k8s_api_node_v1beta1_Overhead                `json:"overhead,omitempty"`
	Scheduling *Io_k8s_api_node_v1beta1_Scheduling              `json:"scheduling,omitempty"`
}


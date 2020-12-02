package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_ServiceReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Path      string `json:"path,omitempty"`
	Port      int    `json:"port,omitempty"`
}


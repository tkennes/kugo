package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_OwnerReference struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion,omitempty"`
	Controller         bool   `json:"controller,omitempty"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}


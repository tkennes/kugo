package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_APIResource struct {
	Categories         []string `json:"categories,omitempty"`
	Group              string   `json:"group,omitempty"`
	Kind               string   `json:"kind"`
	Name               string   `json:"name"`
	Namespaced         bool     `json:"namespaced"`
	ShortNames         []string `json:"shortNames,omitempty"`
	SingularName       string   `json:"singularName"`
	StorageVersionHash string   `json:"storageVersionHash,omitempty"`
	Verbs              []string `json:"verbs"`
	Version            string   `json:"version,omitempty"`
}


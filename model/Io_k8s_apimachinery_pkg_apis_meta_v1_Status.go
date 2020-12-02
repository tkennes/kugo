package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_Status struct {
	ApiVersion string                                              `json:"apiVersion,omitempty"`
	Code       int                                                 `json:"code,omitempty"`
	Details    *Io_k8s_apimachinery_pkg_apis_meta_v1_StatusDetails `json:"details,omitempty"`
	Kind       string                                              `json:"kind,omitempty"`
	Message    string                                              `json:"message,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta      `json:"metadata,omitempty"`
	Reason     string                                              `json:"reason,omitempty"`
	Status     string                                              `json:"status,omitempty"`
}


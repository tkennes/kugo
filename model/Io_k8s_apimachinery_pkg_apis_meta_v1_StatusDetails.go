package kugo_model

type Io_k8s_apimachinery_pkg_apis_meta_v1_StatusDetails struct {
	Causes            []Io_k8s_apimachinery_pkg_apis_meta_v1_StatusCause `json:"causes,omitempty"`
	Group             string                                             `json:"group,omitempty"`
	Kind              string                                             `json:"kind,omitempty"`
	Name              string                                             `json:"name,omitempty"`
	RetryAfterSeconds int                                                `json:"retryAfterSeconds,omitempty"`
	Uid               string                                             `json:"uid,omitempty"`
}


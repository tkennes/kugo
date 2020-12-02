package kugo_model

type Io_k8s_api_core_v1_GitRepoVolumeSource struct {
	Directory  string `json:"directory,omitempty"`
	Repository string `json:"repository"`
	Revision   string `json:"revision,omitempty"`
}


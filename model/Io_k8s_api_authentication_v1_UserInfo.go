package kugo_model

type Io_k8s_api_authentication_v1_UserInfo struct {
	Extra    *interface{} `json:"extra,omitempty"`
	Groups   []string     `json:"groups,omitempty"`
	Uid      string       `json:"uid,omitempty"`
	Username string       `json:"username,omitempty"`
}


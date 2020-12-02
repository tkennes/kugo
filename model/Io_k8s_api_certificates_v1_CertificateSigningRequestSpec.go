package kugo_model

type Io_k8s_api_certificates_v1_CertificateSigningRequestSpec struct {
	Extra      *interface{} `json:"extra,omitempty"`
	Groups     []string     `json:"groups,omitempty"`
	Request    string       `json:"request"`
	SignerName string       `json:"signerName"`
	Uid        string       `json:"uid,omitempty"`
	Usages     []string     `json:"usages,omitempty"`
	Username   string       `json:"username,omitempty"`
}


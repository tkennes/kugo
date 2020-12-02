package kugo_model

type Io_k8s_api_certificates_v1beta1_CertificateSigningRequestStatus struct {
	Certificate string                                                               `json:"certificate,omitempty"`
	Conditions  []Io_k8s_api_certificates_v1beta1_CertificateSigningRequestCondition `json:"conditions,omitempty"`
}


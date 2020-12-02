package kugo_model

type Io_k8s_api_certificates_v1_CertificateSigningRequest struct {
	ApiVersion string                                                      `json:"apiVersion,omitempty"`
	Kind       string                                                      `json:"kind,omitempty"`
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta            `json:"metadata,omitempty"`
	Spec       Io_k8s_api_certificates_v1_CertificateSigningRequestSpec    `json:"spec"`
	Status     *Io_k8s_api_certificates_v1_CertificateSigningRequestStatus `json:"status,omitempty"`
}


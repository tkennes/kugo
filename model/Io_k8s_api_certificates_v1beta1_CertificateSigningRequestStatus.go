package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_certificates_v1beta1_CertificateSigningRequest.go
type Io_k8s_api_certificates_v1beta1_CertificateSigningRequestStatus struct {
	// If request was approved, the controller will place the issued certificate here.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Certificate *string                                                              `json:"certificate,omitempty"`

	// Conditions applied to the request, such as approval or denial.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_certificates_v1beta1_CertificateSigningRequestCondition.go
	Conditions  []Io_k8s_api_certificates_v1beta1_CertificateSigningRequestCondition `json:"conditions,omitempty"`
}

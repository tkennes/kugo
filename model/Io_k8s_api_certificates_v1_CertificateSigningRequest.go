package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_certificates_v1_CertificateSigningRequestList.go


// CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing
// request, and having it asynchronously approved and issued.  Kubelets use this API to obtain:  1. client certificates to
// authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client-kubelet" signerName).  2. serving
// certificates for TLS endpoints kube-apiserver can connect to securely (with the "kubernetes.io/kubelet-serving"
// signerName).  This API can be used to request client certificates to authenticate to kube-apiserver (with the
// "kubernetes.io/kube-apiserver-client" signerName), or to obtain certificates from custom non-Kubernetes signers.
type Io_k8s_api_certificates_v1_CertificateSigningRequest struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion *string                                                     `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind       *string                                                     `json:"kind,omitempty"`

	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta            `json:"metadata,omitempty"`

	// spec contains the certificate request, and is immutable after creation. Only the request, signerName, and usages fields
	// can be set on creation. Other fields are derived by Kubernetes and cannot be modified by users.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_certificates_v1_CertificateSigningRequestSpec.go
	Spec       Io_k8s_api_certificates_v1_CertificateSigningRequestSpec    `json:"spec"`

	// status contains information about whether the request is approved or denied, and the certificate issued by the signer,
	// or the failure condition indicating signer failure.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_certificates_v1_CertificateSigningRequestStatus.go
	Status     *Io_k8s_api_certificates_v1_CertificateSigningRequestStatus `json:"status,omitempty"`
}

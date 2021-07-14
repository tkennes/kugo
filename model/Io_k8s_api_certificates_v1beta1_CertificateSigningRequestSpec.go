package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_certificates_v1beta1_CertificateSigningRequest.go


// This information is immutable after the request is created. Only the Request and Usages fields can be set on creation,
// other fields are derived by Kubernetes and cannot be modified by users.
type Io_k8s_api_certificates_v1beta1_CertificateSigningRequestSpec struct {
	// Extra information about the requesting user. See user.Info interface for details.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Extra      *interface{} `json:"extra,omitempty"`

	// Group information about the requesting user. See user.Info interface for details.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Groups     []*string    `json:"groups,omitempty"`

	// Base64-encoded PKCS#10 CSR data
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Request    *string      `json:"request"`

	// Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be
	// defaulted:  1. If it's a kubelet client certificate, it is assigned     "kubernetes.io/kube-apiserver-client-kubelet".
	// 2. If it's a kubelet serving certificate, it is assigned     "kubernetes.io/kubelet-serving".  3. Otherwise, it is
	// assigned "kubernetes.io/legacy-unknown". Distribution of trust for signers happens out of band. You can select on this
	// field using `spec.signerName`.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SignerName *string      `json:"signerName,omitempty"`

	// UID information about the requesting user. See user.Info interface for details.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Uid        *string      `json:"uid,omitempty"`

	// allowedUsages specifies a set of usage contexts the key will be valid for. See:
	// https://tools.ietf.org/html/rfc5280#section-4.2.1.3      https://tools.ietf.org/html/rfc5280#section-4.2.1.12 Valid
	// values are:  "signing",  "digital signature",  "content commitment",  "key encipherment",  "key agreement",  "data
	// encipherment",  "cert sign",  "crl sign",  "encipher only",  "decipher only",  "any",  "server auth",  "client auth",
	// "code signing",  "email protection",  "s/mime",  "ipsec end system",  "ipsec tunnel",  "ipsec user",  "timestamping",
	// "ocsp signing",  "microsoft sgc",  "netscape sgc"
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Usages     []*string    `json:"usages,omitempty"`

	// Information about the requesting user. See user.Info interface for details.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Username   *string      `json:"username,omitempty"`
}

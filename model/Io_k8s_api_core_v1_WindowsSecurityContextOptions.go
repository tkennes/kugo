package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSecurityContext.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecurityContext.go


// WindowsSecurityContextOptions contain Windows-specific options and credentials.
type Io_k8s_api_core_v1_WindowsSecurityContextOptions struct {
	// GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the
	// contents of the GMSA credential spec named by the GMSACredentialSpecName field.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	GmsaCredentialSpec     *string `json:"gmsaCredentialSpec,omitempty"`

	// GMSACredentialSpecName is the name of the GMSA credential spec to use.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	GmsaCredentialSpecName *string `json:"gmsaCredentialSpecName,omitempty"`

	// The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata
	// if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value
	// specified in SecurityContext takes precedence.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	RunAsUserName          *string `json:"runAsUserName,omitempty"`
}

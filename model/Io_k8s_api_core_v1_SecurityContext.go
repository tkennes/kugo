package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Container.go


// SecurityContext holds security configuration that will be applied to a container. Some fields are present in both
// SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.
type Io_k8s_api_core_v1_SecurityContext struct {
	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly
	// controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the
	// container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	AllowPrivilegeEscalation *bool                                             `json:"allowPrivilegeEscalation,omitempty"`

	// The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the
	// container runtime.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Capabilities.go
	Capabilities             *Io_k8s_api_core_v1_Capabilities                  `json:"capabilities,omitempty"`

	// Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host.
	// Defaults to false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Privileged               *bool                                             `json:"privileged,omitempty"`

	// procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the
	// container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be
	// enabled.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ProcMount                *string                                           `json:"procMount,omitempty"`

	// Whether this container has a read-only root filesystem. Default is false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnlyRootFilesystem   *bool                                             `json:"readOnlyRootFilesystem,omitempty"`

	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in
	// PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes
	// precedence.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	RunAsGroup               *int                                              `json:"runAsGroup,omitempty"`

	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to
	// ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such
	// validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	RunAsNonRoot             *bool                                             `json:"runAsNonRoot,omitempty"`

	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May
	// also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in
	// SecurityContext takes precedence.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	RunAsUser                *int                                              `json:"runAsUser,omitempty"`

	// The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux
	// context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
	// PodSecurityContext, the value specified in SecurityContext takes precedence.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SELinuxOptions.go
	SeLinuxOptions           *Io_k8s_api_core_v1_SELinuxOptions                `json:"seLinuxOptions,omitempty"`

	// The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the
	// container options override the pod options.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SeccompProfile.go
	SeccompProfile           *Io_k8s_api_core_v1_SeccompProfile                `json:"seccompProfile,omitempty"`

	// The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be
	// used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_WindowsSecurityContextOptions.go
	WindowsOptions           *Io_k8s_api_core_v1_WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}

package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicy.go


// PodSecurityPolicySpec defines the policy enforced.
type Io_k8s_api_policy_v1beta1_PodSecurityPolicySpec struct {
	// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to
	// true.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	AllowPrivilegeEscalation        *bool                                                       `json:"allowPrivilegeEscalation,omitempty"`

	// AllowedCSIDrivers is an allowlist of inline CSI drivers that must be explicitly set to be embedded within a pod spec. An
	// empty value indicates that any CSI driver can be used for inline ephemeral volumes. This is a beta field, and is only
	// honored if the API server enables the CSIInlineVolume feature gate.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_AllowedCSIDriver.go
	AllowedCSIDrivers               []Io_k8s_api_policy_v1beta1_AllowedCSIDriver                `json:"allowedCSIDrivers,omitempty"`

	// allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field
	// may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and
	// requiredDropCapabilities.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	AllowedCapabilities             []*string                                                   `json:"allowedCapabilities,omitempty"`

	// allowedFlexVolumes is an allowlist of Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This
	// parameter is effective only when the usage of the Flexvolumes is allowed in the "volumes" field.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_AllowedFlexVolume.go
	AllowedFlexVolumes              []Io_k8s_api_policy_v1beta1_AllowedFlexVolume               `json:"allowedFlexVolumes,omitempty"`

	// allowedHostPaths is an allowlist of host paths. Empty indicates that all host paths may be used.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_AllowedHostPath.go
	AllowedHostPaths                []Io_k8s_api_policy_v1beta1_AllowedHostPath                 `json:"allowedHostPaths,omitempty"`

	// AllowedProcMountTypes is an allowlist of allowed ProcMountTypes. Empty or nil indicates that only the
	// DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	AllowedProcMountTypes           []*string                                                   `json:"allowedProcMountTypes,omitempty"`

	// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain
	// sysctl name or ends in "*" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe
	// sysctls are allowed. Kubelet has to allowlist all allowed unsafe sysctls explicitly to avoid rejection.  Examples: e.g.
	// "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar", "foo.baz", etc.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	AllowedUnsafeSysctls            []*string                                                   `json:"allowedUnsafeSysctls,omitempty"`

	// defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec
	// specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and
	// requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the
	// allowedCapabilities list.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	DefaultAddCapabilities          []*string                                                   `json:"defaultAddCapabilities,omitempty"`

	// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its
	// parent process.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	DefaultAllowPrivilegeEscalation *bool                                                       `json:"defaultAllowPrivilegeEscalation,omitempty"`

	// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name
	// or ends in "*" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are
	// forbidden.  Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar", "foo.baz", etc.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ForbiddenSysctls                []*string                                                   `json:"forbiddenSysctls,omitempty"`

	// fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_FSGroupStrategyOptions.go
	FsGroup                         Io_k8s_api_policy_v1beta1_FSGroupStrategyOptions            `json:"fsGroup"`

	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	HostIPC                         *bool                                                       `json:"hostIPC,omitempty"`

	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	HostNetwork                     *bool                                                       `json:"hostNetwork,omitempty"`

	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	HostPID                         *bool                                                       `json:"hostPID,omitempty"`

	// hostPorts determines which host port ranges are allowed to be exposed.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_HostPortRange.go
	HostPorts                       []Io_k8s_api_policy_v1beta1_HostPortRange                   `json:"hostPorts,omitempty"`

	// privileged determines if a pod can request to be run as privileged.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	Privileged                      *bool                                                       `json:"privileged,omitempty"`

	// readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the
	// container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to
	// false the container may run with a read only root file system if it wishes but it will not be forced to.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	ReadOnlyRootFilesystem          *bool                                                       `json:"readOnlyRootFilesystem,omitempty"`

	// requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped
	// and cannot be added.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	RequiredDropCapabilities        []*string                                                   `json:"requiredDropCapabilities,omitempty"`

	// RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If this field is omitted,
	// the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_RunAsGroupStrategyOptions.go
	RunAsGroup                      *Io_k8s_api_policy_v1beta1_RunAsGroupStrategyOptions        `json:"runAsGroup,omitempty"`

	// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_RunAsUserStrategyOptions.go
	RunAsUser                       Io_k8s_api_policy_v1beta1_RunAsUserStrategyOptions          `json:"runAsUser"`

	// runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this field is omitted, the
	// pod's runtimeClassName field is unrestricted. Enforcement of this field depends on the RuntimeClass feature gate being
	// enabled.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_RuntimeClassStrategyOptions.go
	RuntimeClass                    *Io_k8s_api_policy_v1beta1_RuntimeClassStrategyOptions      `json:"runtimeClass,omitempty"`

	// seLinux is the strategy that will dictate the allowable labels that may be set.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_SELinuxStrategyOptions.go
	SeLinux                         Io_k8s_api_policy_v1beta1_SELinuxStrategyOptions            `json:"seLinux"`

	// supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_SupplementalGroupsStrategyOptions.go
	SupplementalGroups              Io_k8s_api_policy_v1beta1_SupplementalGroupsStrategyOptions `json:"supplementalGroups"`

	// volumes is an allowlist of volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use
	// '*'.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Volumes                         []*string                                                   `json:"volumes,omitempty"`
}

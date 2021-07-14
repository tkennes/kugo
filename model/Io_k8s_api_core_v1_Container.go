package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodSpec.go


// A single application container that you want to run within a pod.
type Io_k8s_api_core_v1_Container struct {
	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are
	// expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be
	// unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be
	// expanded, regardless of whether the variable exists or not. Cannot be updated. More info:
	// https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Args                     []*string                                `json:"args,omitempty"`

	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable
	// references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference
	// in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	// references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info:
	// https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Command                  []*string                                `json:"command,omitempty"`

	// List of environment variables to set in the container. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EnvVar.go
	Env                      []Io_k8s_api_core_v1_EnvVar              `json:"env,omitempty"`

	// List of sources to populate environment variables in the container. The keys defined within a source must be a
	// C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in
	// multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a
	// duplicate key will take precedence. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EnvFromSource.go
	EnvFrom                  []Io_k8s_api_core_v1_EnvFromSource       `json:"envFrom,omitempty"`

	// Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow
	// higher level config management to default or override container images in workload controllers like Deployments and
	// StatefulSets.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Image                    *string                                  `json:"image,omitempty"`

	// Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent
	// otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ImagePullPolicy          *string                                  `json:"imagePullPolicy,omitempty"`

	// Actions that the management system should take in response to container lifecycle events. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Lifecycle.go
	Lifecycle                *Io_k8s_api_core_v1_Lifecycle            `json:"lifecycle,omitempty"`

	// Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Probe.go
	LivenessProbe            *Io_k8s_api_core_v1_Probe                `json:"livenessProbe,omitempty"`

	// Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be
	// updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name                     *string                                  `json:"name"`

	// List of ports to expose from the container. Exposing a port here gives the system additional information about the
	// network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that
	// port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be
	// accessible from the network. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ContainerPort.go
	Ports                    []Io_k8s_api_core_v1_ContainerPort       `json:"ports,omitempty"`

	// Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails.
	// Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Probe.go
	ReadinessProbe           *Io_k8s_api_core_v1_Probe                `json:"readinessProbe,omitempty"`

	// Compute Resources required by this container. Cannot be updated. More info:
	// https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ResourceRequirements.go
	Resources                *Io_k8s_api_core_v1_ResourceRequirements `json:"resources,omitempty"`

	// Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More
	// info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecurityContext.go
	SecurityContext          *Io_k8s_api_core_v1_SecurityContext      `json:"securityContext,omitempty"`

	// StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this
	// completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be
	// used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load
	// data or warm a cache, than during steady-state operation. This cannot be updated. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Probe.go
	StartupProbe             *Io_k8s_api_core_v1_Probe                `json:"startupProbe,omitempty"`

	// Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin
	// in the container will always result in EOF. Default is false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Stdin                    *bool                                    `json:"stdin,omitempty"`

	// Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is
	// true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on
	// container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the
	// client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is
	// false, a container processes that reads from stdin will never receive an EOF. Default is false
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	StdinOnce                *bool                                    `json:"stdinOnce,omitempty"`

	// Optional: Path at which the file to which the container's termination message will be written is mounted into the
	// container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will
	// be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to
	// 12kb. Defaults to /dev/termination-log. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	TerminationMessagePath   *string                                  `json:"terminationMessagePath,omitempty"`

	// Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to
	// populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of
	// container log output if the termination message file is empty and the container exited with an error. The log output is
	// limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	TerminationMessagePolicy *string                                  `json:"terminationMessagePolicy,omitempty"`

	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Tty                      *bool                                    `json:"tty,omitempty"`

	// volumeDevices is the list of block devices to be used by the container.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_VolumeDevice.go
	VolumeDevices            []Io_k8s_api_core_v1_VolumeDevice        `json:"volumeDevices,omitempty"`

	// Pod volumes to mount into the container's filesystem. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_VolumeMount.go
	VolumeMounts             []Io_k8s_api_core_v1_VolumeMount         `json:"volumeMounts,omitempty"`

	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured
	// in the container image. Cannot be updated.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	WorkingDir               *string                                  `json:"workingDir,omitempty"`
}

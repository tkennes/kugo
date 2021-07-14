package model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_NodeStatus.go


// NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
type Io_k8s_api_core_v1_NodeSystemInfo struct {
	// The Architecture reported by the node
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Architecture            *string `json:"architecture"`

	// Boot ID reported by the node.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	BootID                  *string `json:"bootID"`

	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ContainerRuntimeVersion *string `json:"containerRuntimeVersion"`

	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	KernelVersion           *string `json:"kernelVersion"`

	// KubeProxy Version reported by the node.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	KubeProxyVersion        *string `json:"kubeProxyVersion"`

	// Kubelet Version reported by the node.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	KubeletVersion          *string `json:"kubeletVersion"`

	// MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more
	// from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	MachineID               *string `json:"machineID"`

	// The Operating System reported by the node
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	OperatingSystem         *string `json:"operatingSystem"`

	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	OsImage                 *string `json:"osImage"`

	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red
	// Hat hosts https://access.redhat.com/documentation/en-us/red_hat_subscription_management/1/html/rhsm/uuid
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SystemUUID              *string `json:"systemUUID"`
}

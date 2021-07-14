package model


// Tree Depth: 5
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EnvVar.go


// EnvVarSource represents a source for the value of an EnvVar.
type Io_k8s_api_core_v1_EnvVarSource struct {
	// Selects a key of a ConfigMap.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ConfigMapKeySelector.go
	ConfigMapKeyRef  *Io_k8s_api_core_v1_ConfigMapKeySelector  `json:"configMapKeyRef,omitempty"`

	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`,
	// `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ObjectFieldSelector.go
	FieldRef         *Io_k8s_api_core_v1_ObjectFieldSelector   `json:"fieldRef,omitempty"`

	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-
	// storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ResourceFieldSelector.go
	ResourceFieldRef *Io_k8s_api_core_v1_ResourceFieldSelector `json:"resourceFieldRef,omitempty"`

	// Selects a key of a secret in the pod's namespace
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_SecretKeySelector.go
	SecretKeyRef     *Io_k8s_api_core_v1_SecretKeySelector     `json:"secretKeyRef,omitempty"`
}

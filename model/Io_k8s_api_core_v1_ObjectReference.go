package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v1beta1_CronJobStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EndpointAddress.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_events_v1beta1_Event.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Event.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v2alpha1_CronJobStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_events_v1_Event.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ServiceAccount.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_Endpoint.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Binding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_StorageOSPersistentVolumeSource.go


// ObjectReference contains enough information to let you inspect or modify the referred object.
type Io_k8s_api_core_v1_ObjectReference struct {
	// API version of the referent.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion      *string `json:"apiVersion,omitempty"`

	// If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field
	// access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container
	// within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container
	// that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this
	// pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FieldPath       *string `json:"fieldPath,omitempty"`

	// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind            *string `json:"kind,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name            *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace       *string `json:"namespace,omitempty"`

	// Specific resourceVersion to which this reference is made, if any. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ResourceVersion *string `json:"resourceVersion,omitempty"`

	// UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Uid             *string `json:"uid,omitempty"`
}

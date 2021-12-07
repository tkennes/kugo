package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ReplicationControllerSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DeploymentSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DaemonSetSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_StatefulSetSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v1_JobSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_ReplicaSetSpec.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodTemplate.go


// PodTemplateSpec describes the data a pod should have when created from a template
type Io_k8s_api_core_v1_PodTemplateSpec struct {
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-
	// architecture/api-conventions.md#spec-and-status
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodSpec.go
	Spec     *Io_k8s_api_core_v1_PodSpec                      `json:"spec,omitempty"`
}

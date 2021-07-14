package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_Deployment.go


// DeploymentSpec is the specification of the desired behavior of the Deployment.
type Io_k8s_api_apps_v1_DeploymentSpec struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to
	// be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	MinReadySeconds         *int                                               `json:"minReadySeconds,omitempty"`

	// Indicates that the deployment is paused.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Paused                  *bool                                              `json:"paused,omitempty"`

	// The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment
	// controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be
	// surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused.
	// Defaults to 600s.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	ProgressDeadlineSeconds *int                                               `json:"progressDeadlineSeconds,omitempty"`

	// Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Replicas                *int                                               `json:"replicas,omitempty"`

	// The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and
	// not specified. Defaults to 10.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	RevisionHistoryLimit    *int                                               `json:"revisionHistoryLimit,omitempty"`

	// Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this
	// deployment. It must match the pod template's labels.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector                Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector"`

	// The deployment strategy to use to replace existing pods with new ones.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_DeploymentStrategy.go
	Strategy                *Io_k8s_api_apps_v1_DeploymentStrategy             `json:"strategy,omitempty"`

	// Template describes the pods that will be created.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodTemplateSpec.go
	Template                Io_k8s_api_core_v1_PodTemplateSpec                 `json:"template"`
}

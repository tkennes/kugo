package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_DaemonSet.go


// DaemonSetSpec is the specification of a daemon set.
type Io_k8s_api_apps_v1_DaemonSetSpec struct {
	// The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container
	// crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	MinReadySeconds      *int                                               `json:"minReadySeconds,omitempty"`

	// The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not
	// specified. Defaults to 10.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	RevisionHistoryLimit *int                                               `json:"revisionHistoryLimit,omitempty"`

	// A label query over pods that are managed by the daemon set. Must match in order to be controlled. It must match the pod
	// template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector             Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector"`

	// An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every
	// node that matches the template's node selector (or on every node if no node selector is specified). More info:
	// https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodTemplateSpec.go
	Template             Io_k8s_api_core_v1_PodTemplateSpec                 `json:"template"`

	// An update strategy to replace existing DaemonSet pods with new pods.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_DaemonSetUpdateStrategy.go
	UpdateStrategy       *Io_k8s_api_apps_v1_DaemonSetUpdateStrategy        `json:"updateStrategy,omitempty"`
}

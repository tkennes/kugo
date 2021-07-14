package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_StatefulSet.go


// A StatefulSetSpec is the specification of a StatefulSet.
type Io_k8s_api_apps_v1_StatefulSetSpec struct {
	// podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling
	// down. The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the
	// controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite
	// order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without
	// waiting, and on scale down will delete all pods at once.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	PodManagementPolicy  *string                                            `json:"podManagementPolicy,omitempty"`

	// replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are
	// instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults
	// to 1.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Replicas             *int                                               `json:"replicas,omitempty"`

	// revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet's revision history.
	// The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The
	// default value is 10.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	RevisionHistoryLimit *int                                               `json:"revisionHistoryLimit,omitempty"`

	// selector is a label query over pods that should match the replica count. It must match the pod template's labels. More
	// info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector.go
	Selector             Io_k8s_apimachinery_pkg_apis_meta_v1_LabelSelector `json:"selector"`

	// serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet,
	// and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-
	// string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ServiceName          *string                                            `json:"serviceName"`

	// template is the object that describes the pod that will be created if insufficient replicas are detected. Each pod
	// stamped out by the StatefulSet will fulfill this Template, but have a unique identity from the rest of the StatefulSet.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodTemplateSpec.go
	Template             Io_k8s_api_core_v1_PodTemplateSpec                 `json:"template"`

	// updateStrategy indicates the StatefulSetUpdateStrategy that will be employed to update Pods in the StatefulSet when a
	// revision is made to Template.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_StatefulSetUpdateStrategy.go
	UpdateStrategy       *Io_k8s_api_apps_v1_StatefulSetUpdateStrategy      `json:"updateStrategy,omitempty"`

	// volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible
	// for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must
	// have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence
	// over any volumes in the template, with the same name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaim.go
	VolumeClaimTemplates []Io_k8s_api_core_v1_PersistentVolumeClaim         `json:"volumeClaimTemplates,omitempty"`
}

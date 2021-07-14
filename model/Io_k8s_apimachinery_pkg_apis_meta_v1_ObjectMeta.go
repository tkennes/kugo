package model

import (
	"time"
)


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Service.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Endpoints.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v1_Job.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1alpha1_Role.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SelfSubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinition.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ServiceAccount.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v2alpha1_JobTemplateSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_CSINode.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_Eviction.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1beta1_RoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1beta1_RuntimeClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Pod.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1beta1_Role.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_discovery_v1beta1_EndpointSlice.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ComponentStatus.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_LocalSubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Secret.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_NetworkPolicy.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SelfSubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaim.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodTemplate.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v1_HorizontalPodAutoscaler.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_admissionregistration_v1beta1_MutatingWebhookConfiguration.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SelfSubjectRulesReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfiguration.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1beta1_TokenReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_Role.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfiguration.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSIDriver.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_IngressClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1alpha1_ClusterRole.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinition.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1_TokenRequest.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscaler.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ConfigMap.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_LimitRange.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_FlowSchema.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolume.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_flowcontrol_v1beta1_FlowSchema.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v1beta1_JobTemplateSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_PodDisruptionBudget.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscaler.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Node.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authentication_v1_TokenReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_ClusterRoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1alpha1_RoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_VolumeAttachment.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_Deployment.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_coordination_v1_Lease.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_StorageClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ReplicationController.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PodTemplateSpec.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersion.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIService.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v1beta1_CronJob.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1beta1_SelfSubjectRulesReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimTemplate.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_scheduling_v1beta1_PriorityClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_admissionregistration_v1_MutatingWebhookConfiguration.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_DaemonSet.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_admissionregistration_v1beta1_ValidatingWebhookConfiguration.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_ControllerRevision.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_scheduling_v1alpha1_PriorityClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_CSINode.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1beta1_ClusterRole.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1beta1_ClusterRoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIService.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_ClusterRole.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_events_v1beta1_Event.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_ResourceQuota.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_scheduling_v1_PriorityClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicy.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_batch_v2alpha1_CronJob.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_Ingress.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_coordination_v1beta1_Lease.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Namespace.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1_Ingress.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_autoscaling_v1_Scale.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Event.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_admissionregistration_v1_ValidatingWebhookConfiguration.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_certificates_v1_CertificateSigningRequest.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1alpha1_RuntimeClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1_RoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1_StorageClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_networking_v1beta1_IngressClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_ReplicaSet.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_apps_v1_StatefulSet.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_LocalSubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_extensions_v1beta1_Ingress.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_node_v1_RuntimeClass.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Binding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_rbac_v1alpha1_ClusterRoleBinding.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_certificates_v1beta1_CertificateSigningRequest.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_events_v1_Event.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_VolumeAttachment.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_authorization_v1_SubjectAccessReview.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1alpha1_VolumeAttachment.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_storage_v1beta1_CSIDriver.go


// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
type Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and
	// retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info:
	// http://kubernetes.io/docs/user-guide/annotations
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Annotations                *interface{}                                              `json:"annotations,omitempty"`

	// The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace
	// in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or
	// update request.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ClusterName                *string                                                   `json:"clusterName,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be
	// set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339
	// form and is in UTC.  Populated by the system. Read-only. Null for lists. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	CreationTimestamp          *time.Time                                                `json:"creationTimestamp,omitempty"`

	// Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set
	// when deletionTimestamp is also set. May only be shortened. Read-only.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	DeletionGracePeriodSeconds *int                                                      `json:"deletionGracePeriodSeconds,omitempty"`

	// DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when
	// a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be
	// deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the
	// finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp
	// is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may
	// be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will
	// react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will
	// send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the
	// presence of network partitions, this object may still exist after this timestamp, until an administrator or automated
	// process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been
	// requested.  Populated by the system when a graceful deletion is requested. Read-only. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	DeletionTimestamp          *time.Time                                                `json:"deletionTimestamp,omitempty"`

	// Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component
	// that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can
	// only be removed. Finalizers may be processed and removed in any order.  Order is NOT enforced because it introduces
	// significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the
	// finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first
	// finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component
	// responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to
	// order amongst themselves and are not vulnerable to ordering changes in the list.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Finalizers                 []*string                                                 `json:"finalizers,omitempty"`

	// GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been
	// provided. If this field is used, the name returned to the client will be different than the name passed. This value will
	// also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be
	// truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and
	// the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with
	// Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry
	// (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	GenerateName               *string                                                   `json:"generateName,omitempty"`

	// A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Generation                 *int                                                      `json:"generation,omitempty"`

	// Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match
	// selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/interface{}.go
	Labels                     *interface{}                                              `json:"labels,omitempty"`

	// ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for
	// internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's
	// name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version
	// that the workflow used when modifying the object.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ManagedFieldsEntry.go
	ManagedFields              []Io_k8s_apimachinery_pkg_apis_meta_v1_ManagedFieldsEntry `json:"managedFields,omitempty"`

	// Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client
	// to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and
	// configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name                       *string                                                   `json:"name,omitempty"`

	// Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the "default"
	// namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the
	// value of this field for those objects will be empty.  Must be a DNS_LABEL. Cannot be updated. More info:
	// http://kubernetes.io/docs/user-guide/namespaces
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Namespace                  *string                                                   `json:"namespace,omitempty"`

	// List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage
	// collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the
	// controller field set to true. There cannot be more than one managing controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_OwnerReference.go
	OwnerReferences            []Io_k8s_apimachinery_pkg_apis_meta_v1_OwnerReference     `json:"ownerReferences,omitempty"`

	// An opaque value that represents the internal version of this object that can be used by clients to determine when
	// objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or
	// set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be
	// valid for a particular resource or set of resources.  Populated by the system. Read-only. Value must be treated as
	// opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#concurrency-control-and-consistency
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ResourceVersion            *string                                                   `json:"resourceVersion,omitempty"`

	// SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop
	// propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SelfLink                   *string                                                   `json:"selfLink,omitempty"`

	// UID is the unique in time and space value for this object. It is typically generated by the server on successful
	// creation of a resource and is not allowed to change on PUT operations.  Populated by the system. Read-only. More info:
	// http://kubernetes.io/docs/user-guide/identifiers#uids
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Uid                        *string                                                   `json:"uid,omitempty"`
}

package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_extensions_v1beta1_IngressList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1beta1_RoleList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodSecurityPolicyList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_node_v1beta1_RuntimeClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1beta1_PriorityLevelConfigurationList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1alpha1_ClusterRoleList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_ControllerRevisionList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_admissionregistration_v1_MutatingWebhookConfigurationList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1beta1_RoleBindingList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1alpha1_RoleBindingList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_scheduling_v1beta1_PriorityClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PodTemplateList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NodeList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1alpha1_RoleList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_PriorityLevelConfigurationList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1beta1_VolumeAttachmentList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EndpointsList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1_RoleBindingList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_IngressList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1beta1_CSIDriverList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1_APIServiceList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_EventList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v1_HorizontalPodAutoscalerList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_NetworkPolicyList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1beta1_IngressList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1alpha1_ClusterRoleBindingList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_scheduling_v1_PriorityClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ServiceList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_policy_v1beta1_PodDisruptionBudgetList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_certificates_v1beta1_CertificateSigningRequestList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DeploymentList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_ReplicaSetList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1_CSIDriverList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1beta1_FlowSchemaList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_admissionregistration_v1beta1_MutatingWebhookConfigurationList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeClaimList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1_ClusterRoleList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1_VolumeAttachmentList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ReplicationControllerList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta1_HorizontalPodAutoscalerList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_DaemonSetList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_events_v1_EventList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_admissionregistration_v1beta1_ValidatingWebhookConfigurationList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_SecretList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1beta1_ClusterRoleBindingList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_node_v1alpha1_RuntimeClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_events_v1beta1_EventList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1_IngressClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v2alpha1_CronJobList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1_StorageClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1beta1_ClusterRoleList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v1beta1_CronJobList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1_RoleList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_scheduling_v1alpha1_PriorityClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_coordination_v1beta1_LeaseList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_PersistentVolumeList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_coordination_v1_LeaseList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_kube__aggregator_pkg_apis_apiregistration_v1beta1_APIServiceList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_rbac_v1_ClusterRoleBindingList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ServiceAccountList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1beta1_CSINodeList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_discovery_v1beta1_EndpointSliceList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_certificates_v1_CertificateSigningRequestList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1_CSINodeList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_NamespaceList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ComponentStatusList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v1_JobList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_LimitRangeList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_apps_v1_StatefulSetList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_node_v1_RuntimeClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1alpha1_VolumeAttachmentList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_storage_v1beta1_StorageClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_Status.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_flowcontrol_v1alpha1_FlowSchemaList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_networking_v1beta1_IngressClassList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ResourceQuotaList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_autoscaling_v2beta2_HorizontalPodAutoscalerList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_core_v1_ConfigMapList.go
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_admissionregistration_v1_ValidatingWebhookConfigurationList.go


// ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource
// may have only one of {ObjectMeta, ListMeta}.
type Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta struct {
	// continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data
	// available. The value is opaque and may be used to issue another request to the endpoint that served this list to
	// retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration
	// has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value
	// will be identical to the value in the first response, unless you have received this token from an error message.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Continue           *string `json:"continue,omitempty"`

	// remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the
	// list request contained label or field selectors, then the number of remaining items is unknown and the field will be
	// left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this
	// is the last chunk), then there are no more remaining items and this field will be left unset and omitted during
	// serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is
	// *estimating* the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	RemainingItemCount *int    `json:"remainingItemCount,omitempty"`

	// String that identifies the server's internal version of this object that can be used by clients to determine when
	// objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by
	// the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#concurrency-control-and-consistency
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ResourceVersion    *string `json:"resourceVersion,omitempty"`

	// selfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop
	// propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	SelfLink           *string `json:"selfLink,omitempty"`
}

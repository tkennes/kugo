package model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_policy_v1beta1_Eviction.go


// DeleteOptions may be provided when deleting an API object.
type Io_k8s_apimachinery_pkg_apis_meta_v1_DeleteOptions struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion         *string                                             `json:"apiVersion,omitempty"`

	// When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will
	// be processed
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	DryRun             []*string                                           `json:"dryRun,omitempty"`

	// The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero
	// indicates delete immediately. If this value is nil, the default grace period for the specified type will be used.
	// Defaults to a per object value if not specified. zero means delete immediately.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	GracePeriodSeconds *int                                                `json:"gracePeriodSeconds,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind               *string                                             `json:"kind,omitempty"`

	// Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be
	// orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list. Either this
	// field or PropagationPolicy may be set, but not both.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	OrphanDependents   *bool                                               `json:"orphanDependents,omitempty"`

	// Must be fulfilled before a deletion is carried out. If not possible, a 409 Conflict status will be returned.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_Preconditions.go
	Preconditions      *Io_k8s_apimachinery_pkg_apis_meta_v1_Preconditions `json:"preconditions,omitempty"`

	// Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both.
	// The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default
	// policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete
	// the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	PropagationPolicy  *string                                             `json:"propagationPolicy,omitempty"`
}

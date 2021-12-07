package kugo_model


// Tree Depth: 1
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_scheduling_v1alpha1_PriorityClassList.go


// DEPRECATED - This group version of PriorityClass is deprecated by scheduling.k8s.io/v1/PriorityClass. PriorityClass
// defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
type Io_k8s_api_scheduling_v1alpha1_PriorityClass struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	ApiVersion       *string                                          `json:"apiVersion,omitempty"`

	// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Description      *string                                          `json:"description,omitempty"`

	// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not
	// have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one
	// PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default
	// PriorityClasses will be used as the default priority.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	GlobalDefault    *bool                                            `json:"globalDefault,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Kind             *string                                          `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#metadata
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata         *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to
	// PreemptLowerPriority if unset. This field is beta-level, gated by the NonPreemptingPriority feature-gate.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	PreemptionPolicy *string                                          `json:"preemptionPolicy,omitempty"`

	// The value of this priority class. This is the actual priority that pods receive when they have the name of this class in
	// their pod spec.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Value            *int                                             `json:"value"`
}

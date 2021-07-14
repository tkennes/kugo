package model


// Status is a return value for calls that don't return other objects.
type Io_k8s_apimachinery_pkg_apis_meta_v1_Status struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas
	// to the latest internal value, and may reject unrecognized values. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                             `json:"apiVersion,omitempty"`

	// Suggested HTTP return code for this status, 0 if not set.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/int.go
	Code       *int                                                `json:"code,omitempty"`

	// Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and
	// the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_StatusDetails.go
	Details    *Io_k8s_apimachinery_pkg_apis_meta_v1_StatusDetails `json:"details,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint
	// the client submits requests to. Cannot be updated. In CamelCase. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind       *string                                             `json:"kind,omitempty"`

	// A human-readable description of the status of this operation.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Message    *string                                             `json:"message,omitempty"`

	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta.go
	Metadata   *Io_k8s_apimachinery_pkg_apis_meta_v1_ListMeta      `json:"metadata,omitempty"`

	// A machine-readable description of why this operation is in the "Failure" status. If this value is empty there is no
	// information available. A Reason clarifies an HTTP status code but does not override it.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Reason     *string                                             `json:"reason,omitempty"`

	// Status of the operation. One of: "Success" or "Failure". More info: https://git.k8s.io/community/contributors/devel/sig-
	// architecture/api-conventions.md#spec-and-status
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Status     *string                                             `json:"status,omitempty"`
}

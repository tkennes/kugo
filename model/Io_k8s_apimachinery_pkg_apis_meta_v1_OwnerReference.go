package model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go


// OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same
// namespace as the dependent, or be cluster-scoped, so there is no namespace field.
type Io_k8s_apimachinery_pkg_apis_meta_v1_OwnerReference struct {
	// API version of the referent.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion         *string `json:"apiVersion"`

	// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value
	// store until this reference is removed. Defaults to false. To set this field, a user needs "delete" permission of the
	// owner, otherwise 422 (Unprocessable Entity) will be returned.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	BlockOwnerDeletion *bool   `json:"blockOwnerDeletion,omitempty"`

	// If true, this reference points to the managing controller.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	Controller         *bool   `json:"controller,omitempty"`

	// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-
	// conventions.md#types-kinds
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Kind               *string `json:"kind"`

	// Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name               *string `json:"name"`

	// UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Uid                *string `json:"uid"`
}

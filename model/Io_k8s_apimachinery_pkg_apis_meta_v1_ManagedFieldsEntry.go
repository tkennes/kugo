package model

import (
	"time"
)


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go


// ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
type Io_k8s_apimachinery_pkg_apis_meta_v1_ManagedFieldsEntry struct {
	// APIVersion defines the version of this resource that this field set applies to. The format is "group/version" just like
	// the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically
	// converted.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	ApiVersion *string                                        `json:"apiVersion,omitempty"`

	// FieldsType is the discriminator for the different fields format and version. There is currently only one possible value:
	// "FieldsV1"
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	FieldsType *string                                        `json:"fieldsType,omitempty"`

	// FieldsV1 holds the first JSON version format as described in the "FieldsV1" type.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_FieldsV1.go
	FieldsV1   *Io_k8s_apimachinery_pkg_apis_meta_v1_FieldsV1 `json:"fieldsV1,omitempty"`

	// Manager is an identifier of the workflow managing these fields.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Manager    *string                                        `json:"manager,omitempty"`

	// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this
	// field are 'Apply' and 'Update'.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Operation  *string                                        `json:"operation,omitempty"`

	// Time is timestamp of when these fields were set. It should always be empty if Operation is 'Apply'
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/time.Time.go
	Time       *time.Time                                     `json:"time,omitempty"`
}

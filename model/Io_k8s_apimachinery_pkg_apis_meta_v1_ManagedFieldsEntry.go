package kugo_model

import (
	"time"
)

type Io_k8s_apimachinery_pkg_apis_meta_v1_ManagedFieldsEntry struct {
	ApiVersion string                                         `json:"apiVersion,omitempty"`
	FieldsType string                                         `json:"fieldsType,omitempty"`
	FieldsV1   *Io_k8s_apimachinery_pkg_apis_meta_v1_FieldsV1 `json:"fieldsV1,omitempty"`
	Manager    string                                         `json:"manager,omitempty"`
	Operation  string                                         `json:"operation,omitempty"`
	Time       *time.Time                                     `json:"time,omitempty"`
}


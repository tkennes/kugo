package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v2alpha1_CronJobSpec.go


// JobTemplateSpec describes the data a Job should have when created from a template
type Io_k8s_api_batch_v2alpha1_JobTemplateSpec struct {
	// Standard object's metadata of the jobs created from this template. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta.go
	Metadata *Io_k8s_apimachinery_pkg_apis_meta_v1_ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-
	// architecture/api-conventions.md#spec-and-status
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_api_batch_v1_JobSpec.go
	Spec     *Io_k8s_api_batch_v1_JobSpec                     `json:"spec,omitempty"`
}

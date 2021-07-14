package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_EphemeralContainer.go
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Container.go


// VolumeMount describes a mounting of a Volume within a container.
type Io_k8s_api_core_v1_VolumeMount struct {
	// Path within the container at which the volume should be mounted.  Must not contain ':'.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	MountPath        *string `json:"mountPath"`

	// mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set,
	// MountPropagationNone is used. This field is beta in 1.10.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	MountPropagation *string `json:"mountPropagation,omitempty"`

	// This must match the Name of a Volume.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Name             *string `json:"name"`

	// Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/bool.go
	ReadOnly         *bool   `json:"readOnly,omitempty"`

	// Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SubPath          *string `json:"subPath,omitempty"`

	// Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but
	// environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to "" (volume's
	// root). SubPathExpr and SubPath are mutually exclusive.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	SubPathExpr      *string `json:"subPathExpr,omitempty"`
}

package kugo_model


// Tree Depth: 3
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionSpec.go


// CustomResourceDefinitionVersion describes a version for CRD.
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceDefinitionVersion struct {
	// additionalPrinterColumns specifies additional columns returned in Table output. See
	// https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. Top-level and
	// per-version columns are mutually exclusive. Per-version columns must not all be set to identical values (top-level
	// columns should be used instead). If no top-level or per-version columns are specified, a single column displaying the
	// age of the custom resource is used.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceColumnDefinition.go
	AdditionalPrinterColumns []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceColumnDefinition `json:"additionalPrinterColumns,omitempty"`

	// deprecated indicates this version of the custom resource API is deprecated. When set to true, API requests to this
	// version receive a warning header in the server response. Defaults to false.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	Deprecated               *bool                                                                                           `json:"deprecated,omitempty"`

	// deprecationWarning overrides the default warning returned to API clients. May only be set when `deprecated` is true. The
	// default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater
	// stability, if one exists.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	DeprecationWarning       *string                                                                                         `json:"deprecationWarning,omitempty"`

	// name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at
	// `/apis/<group>/<version>/...` if `served` is true.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name                     *string                                                                                         `json:"name"`

	// schema describes the schema used for validation and pruning of this version of the custom resource. Top-level and per-
	// version schemas are mutually exclusive. Per-version schemas must not all be set to identical values (top-level
	// validation schema should be used instead).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceValidation.go
	Schema                   *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceValidation        `json:"schema,omitempty"`

	// served is a flag enabling/disabling this version from being served via REST APIs
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	Served                   *bool                                                                                           `json:"served"`

	// storage indicates this version should be used when persisting custom resources to storage. There must be exactly one
	// version with storage=true.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/bool.go
	Storage                  *bool                                                                                           `json:"storage"`

	// subresources specify what subresources this version of the defined custom resource have. Top-level and per-version
	// subresources are mutually exclusive. Per-version subresources must not all be set to identical values (top-level
	// subresources should be used instead).
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceSubresources.go
	Subresources             *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceSubresources      `json:"subresources,omitempty"`
}

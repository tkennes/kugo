package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionVersion.go


// CustomResourceColumnDefinition specifies a column for server side printing.
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceColumnDefinition struct {
	// description is a human readable description of this column.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Description *string `json:"description,omitempty"`

	// format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier
	// column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-
	// Specification/blob/master/versions/2.0.md#data-types for details.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Format      *string `json:"format,omitempty"`

	// jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the
	// value for this column.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	JsonPath    *string `json:"jsonPath"`

	// name is a human readable name for the column.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Name        *string `json:"name"`

	// priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered
	// higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/int.go
	Priority    *int    `json:"priority,omitempty"`

	// type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-
	// Specification/blob/master/versions/2.0.md#data-types for details.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Type        *string `json:"type"`
}

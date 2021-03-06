package kugo_model


// Tree Depth: 2
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apimachinery_pkg_apis_meta_v1_StatusDetails.go


// StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.
type Io_k8s_apimachinery_pkg_apis_meta_v1_StatusCause struct {
	// The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix
	// notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to
	// fields having multiple errors. Optional.  Examples:   "name" - the field "name" on the current resource
	// "items[0].name" - the field "name" on the first array entry in "items"
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Field   *string `json:"field,omitempty"`

	// A human-readable description of the cause of the error.  This field may be presented as-is to a reader.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Message *string `json:"message,omitempty"`

	// A machine-readable description of the cause of the error. If this value is empty there is no information available.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/string.go
	Reason  *string `json:"reason,omitempty"`
}

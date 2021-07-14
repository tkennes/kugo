package model


// IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it
// produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or
// number.
type Io_k8s_apimachinery_pkg_util_intstr_IntOrString struct {
}

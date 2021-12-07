package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceDefinitionVersion.go


// CustomResourceSubresources defines the status and scale subresources for CustomResources.
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceSubresources struct {
	// scale indicates the custom resource should serve a `/scale` subresource that returns an `autoscaling/v1` Scale object.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceSubresourceScale.go
	Scale  *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceSubresourceScale  `json:"scale,omitempty"`

	// status indicates the custom resource should serve a `/status` subresource. When enabled: 1. requests to the custom
	// resource primary endpoint ignore changes to the `status` stanza of the object. 2. requests to the custom resource
	// `/status` subresource ignore changes to anything other than the `status` stanza of the object.
	// See: file:///Users/tomkennes/Documents/Clients/_generic/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceSubresourceStatus.go
	Status *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_CustomResourceSubresourceStatus `json:"status,omitempty"`
}

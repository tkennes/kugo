package model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceSubresources.go


// CustomResourceSubresourceStatus defines how to serve the status subresource for CustomResources. Status is represented
// by the `.status` JSON path inside of a CustomResource. When set, * exposes a /status subresource for the custom resource
// * PUT requests to the /status subresource take a custom resource object, and ignore changes to anything except the
// status stanza * PUT/POST/PATCH requests to the custom resource ignore changes to the status stanza
type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1beta1_CustomResourceSubresourceStatus struct {
}

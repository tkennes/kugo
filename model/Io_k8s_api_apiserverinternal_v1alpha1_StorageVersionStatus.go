package kugo_model

type Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionStatus struct {
	CommonEncodingVersion string                                                          `json:"commonEncodingVersion,omitempty"`
	Conditions            []Io_k8s_api_apiserverinternal_v1alpha1_StorageVersionCondition `json:"conditions,omitempty"`
	StorageVersions       []Io_k8s_api_apiserverinternal_v1alpha1_ServerStorageVersion    `json:"storageVersions,omitempty"`
}


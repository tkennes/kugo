package kugo_model


// Tree Depth: 4
// REFERENCES:
// - file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/Io_k8s_api_core_v1_Volume.go


// Represents a volume that is populated with the contents of a git repository. Git repo volumes do not support ownership
// management. Git repo volumes support SELinux relabeling.  DEPRECATED: GitRepo is deprecated. To provision a container
// with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into
// the Pod's container.
type Io_k8s_api_core_v1_GitRepoVolumeSource struct {
	// Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git
	// repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given
	// name.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Directory  *string `json:"directory,omitempty"`

	// Repository URL
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Repository *string `json:"repository"`

	// Commit hash for the specified revision.
	// See: file:///Users/tomkennes/Clients/Volksbank/code/custom/kugo/model/string.go
	Revision   *string `json:"revision,omitempty"`
}

package version

var (
	gitMajor = "1"
	gitMinor = "36"
	gitVersion   = "v1.36.3-k3s1"
	gitCommit    = "aff26ddbbd423ebb6cf64f7bb46246368b459c4c"
	gitTreeState = "clean"
	buildDate = "2026-09-16T16:13:47Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.36"
)

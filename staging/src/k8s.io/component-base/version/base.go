package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.11-k3s1"
	gitCommit    = "c76c43869e8f7438c3d6ae109bb364567ef3ce0a"
	gitTreeState = "clean"
	buildDate = "2026-09-16T14:43:05Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

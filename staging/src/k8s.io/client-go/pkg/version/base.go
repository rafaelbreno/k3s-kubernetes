package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.2-k3s1"
	gitCommit    = "f5aad1d71b1559f5b50e6e5012c969043c39d1ec"
	gitTreeState = "clean"
	buildDate = "2025-11-12T22:43:56Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

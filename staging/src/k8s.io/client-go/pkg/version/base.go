package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.1-k3s1"
	gitCommit    = "f21d18dbb4df91cba0cc8f27b8a5daf21a988228"
	gitTreeState = "clean"
	buildDate = "2026-02-10T22:20:52Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)

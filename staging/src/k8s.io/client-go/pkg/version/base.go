package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.0-k3s1"
	gitCommit    = "4e5078419dfd7258d0d541b455fb536db0e6d8a0"
	gitTreeState = "clean"
	buildDate = "2025-12-18T00:45:19Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)

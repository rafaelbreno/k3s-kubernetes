package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.10-k3s1"
	gitCommit    = "1b35a1e74ab15e7ef4e35c47607c643c9890ac1b"
	gitTreeState = "clean"
	buildDate = "2026-07-22T22:20:55Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

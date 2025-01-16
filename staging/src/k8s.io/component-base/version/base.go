package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "32"
	gitVersion   = "v1.32.1-k3s1"
	gitCommit    = "6a02d26ffb39c6adf2f315be6f05c0d6c1aec1cb"
	gitTreeState = "clean"
	buildDate = "2025-01-16T01:09:20Z"
)

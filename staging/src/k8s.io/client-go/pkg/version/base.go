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
	gitVersion   = "v1.32.3-k3s1"
	gitCommit    = "e9ac656ad7336b0278eaf7374e0e4bbd0feb3bf1"
	gitTreeState = "clean"
	buildDate = "2025-03-12T17:25:21Z"
)

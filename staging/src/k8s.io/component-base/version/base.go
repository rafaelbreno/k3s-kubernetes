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
	gitVersion   = "v1.32.10-k3s1"
	gitCommit    = "a6efaa307f7d7f2380867a13de2bb753826d5242"
	gitTreeState = "clean"
	buildDate = "2025-11-12T23:22:42Z"
)

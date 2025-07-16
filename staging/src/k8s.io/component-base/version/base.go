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
	gitVersion   = "v1.32.7-k3s1"
	gitCommit    = "eb3a5e9a8657a4ac1c0c4c05dcfaf8c23bbb66c9"
	gitTreeState = "clean"
	buildDate = "2025-07-16T14:49:49Z"
)

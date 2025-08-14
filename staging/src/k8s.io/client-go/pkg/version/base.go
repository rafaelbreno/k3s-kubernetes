package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "31"
	gitVersion   = "v1.31.12-k3s1"
	gitCommit    = "33147d042aee51f2b32001e0f281a2123eee76cd"
	gitTreeState = "clean"
	buildDate = "2025-08-14T15:21:03Z"
)

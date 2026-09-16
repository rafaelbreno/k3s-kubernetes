package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.8-k3s1"
	gitCommit    = "032617bea7afd3e7465e7e138e90476dc6591722"
	gitTreeState = "clean"
	buildDate = "2026-09-16T15:37:53Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)

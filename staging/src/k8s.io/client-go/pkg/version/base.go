package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.9-k3s1"
	gitCommit    = "bfa66495b00f575273d7b3493cffb61457d438c0"
	gitTreeState = "clean"
	buildDate = "2026-07-21T19:37:52Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

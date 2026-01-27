package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "33"
	gitVersion   = "v1.33.7-k3s2"
	gitCommit    = "c3ac10a179bc64b5474ac6349cadb7f718ec8f66"
	gitTreeState = "clean"
	buildDate = "2026-01-27T16:19:21Z"
)

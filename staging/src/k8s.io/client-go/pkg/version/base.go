package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.3-k3s2"
	gitCommit    = "ff993c304719975122c9acfe216a2dada3443630"
	gitTreeState = "clean"
	buildDate = "2026-01-27T16:21:30Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

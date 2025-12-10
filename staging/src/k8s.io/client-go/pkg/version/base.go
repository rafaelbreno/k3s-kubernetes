package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.3-k3s1"
	gitCommit    = "ff993c304719975122c9acfe216a2dada3443630"
	gitTreeState = "clean"
	buildDate = "2025-12-10T15:09:01Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

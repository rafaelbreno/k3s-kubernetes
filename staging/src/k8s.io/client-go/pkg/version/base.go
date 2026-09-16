package version

var (
	gitMajor = "1"
	gitMinor = "36"
	gitVersion   = "v1.36.4-k3s1"
	gitCommit    = "8bbb3ebd09c8644a231134d9b464a20088d0610b"
	gitTreeState = "clean"
	buildDate = "2026-09-16T15:45:49Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.36"
)

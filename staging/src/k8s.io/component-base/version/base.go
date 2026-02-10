package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.4-k3s1"
	gitCommit    = "80cff1ea62db567b8f433a6caf710df141fed7cd"
	gitTreeState = "clean"
	buildDate = "2026-02-10T22:18:49Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

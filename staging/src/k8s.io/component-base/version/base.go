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
	gitVersion   = "v1.33.11-k3s1"
	gitCommit    = "41067dbbd2b4bec726348a5e10dd4ebc4487d939"
	gitTreeState = "clean"
	buildDate = "2026-04-15T23:05:54Z"
)

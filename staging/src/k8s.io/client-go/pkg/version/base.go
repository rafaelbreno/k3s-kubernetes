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
	gitVersion   = "v1.33.9-k3s1"
	gitCommit    = "e8bc4ef208527a872877c664b0786901aa9a821d"
	gitTreeState = "clean"
	buildDate = "2026-02-27T13:40:09Z"
)

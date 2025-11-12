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
	gitVersion   = "v1.31.14-k3s1"
	gitCommit    = "26588d2dcbd2590686fb73958e31556e7db0ee5c"
	gitTreeState = "clean"
	buildDate = "2025-11-12T23:34:08Z"
)

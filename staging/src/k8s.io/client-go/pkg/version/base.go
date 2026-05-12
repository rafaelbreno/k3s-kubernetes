package version

var (
	gitMajor = "1"
	gitMinor = "36"
	gitVersion   = "v1.36.1-k3s1"
	gitCommit    = "4c8b574a932129c84e2d13115606191e6486db8d"
	gitTreeState = "clean"
	buildDate = "2026-05-12T20:09:58Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.36"
)

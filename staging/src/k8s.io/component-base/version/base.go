package version

var (
	gitMajor = "1"
	gitMinor = "36"
	gitVersion   = "v1.36.2-k3s1"
	gitCommit    = "364ba6303ec5e5225c3a81592418fa7ef74ff5ad"
	gitTreeState = "clean"
	buildDate = "2026-07-21T19:37:33Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.36"
)

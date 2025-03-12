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
	gitVersion   = "v1.31.7-k3s1"
	gitCommit    = "c69a796ed103e243013b14f0f1956c6d7365ed75"
	gitTreeState = "clean"
	buildDate = "2025-03-12T17:24:44Z"
)

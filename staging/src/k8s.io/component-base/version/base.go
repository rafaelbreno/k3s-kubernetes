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
	gitVersion   = "v1.33.5-k3s1"
	gitCommit    = "8e817a1e48a93696d44b590230be236edc259fcc"
	gitTreeState = "clean"
	buildDate = "2025-09-10T19:09:53Z"
)

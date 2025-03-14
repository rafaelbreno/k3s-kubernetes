package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.32"
)

var (
	gitMajor = "1"
	gitMinor = "32"
	gitVersion   = "v1.32.3-k3s2"
	gitCommit    = "f1a97139d5d197b065b0f7c09b0c1475135a65ea"
	gitTreeState = "clean"
	buildDate = "2025-03-14T17:20:44Z"
)

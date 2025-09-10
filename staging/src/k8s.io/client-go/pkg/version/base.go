package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "32"
	gitVersion   = "v1.32.9-k3s1"
	gitCommit    = "c8e5a7dacaa4ece174359ad536400b2ce1d792fa"
	gitTreeState = "clean"
	buildDate = "2025-09-10T19:13:00Z"
)

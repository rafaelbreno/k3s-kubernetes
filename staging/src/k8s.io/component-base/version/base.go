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
	gitVersion   = "v1.32.11-k3s1"
	gitCommit    = "3db2ef09f9a1a0cb029031c5979a069bae35bcc8"
	gitTreeState = "clean"
	buildDate = "2025-12-16T23:32:23Z"
)

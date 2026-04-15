package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.7-k3s1"
	gitCommit    = "dfb425923e11651c74348c73a3c2f90cff276702"
	gitTreeState = "clean"
	buildDate = "2026-04-15T23:07:56Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

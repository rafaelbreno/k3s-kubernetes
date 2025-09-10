package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.1-k3s1"
	gitCommit    = "85ce3b0f0d16eea6651098eeb4c20d4402577287"
	gitTreeState = "clean"
	buildDate = "2025-09-10T19:06:13Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

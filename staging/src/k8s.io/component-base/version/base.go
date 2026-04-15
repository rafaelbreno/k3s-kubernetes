package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.4-k3s1"
	gitCommit    = "23810ed924466760a3ed10959f367a6a17af4f70"
	gitTreeState = "clean"
	buildDate = "2026-04-15T23:10:55Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)

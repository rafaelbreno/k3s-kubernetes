package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.3-k3s1"
	gitCommit    = "838972d9c9333510bca636b785d6cc57c8ae839d"
	gitTreeState = "clean"
	buildDate = "2026-03-19T14:22:53Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)

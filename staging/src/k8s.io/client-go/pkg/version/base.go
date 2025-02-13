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
	gitVersion   = "v1.31.6-k3s1"
	gitCommit    = "7e4445dba0c7c2721dae774cc006f7bcdadbced2"
	gitTreeState = "clean"
	buildDate = "2025-02-13T18:03:15Z"
)

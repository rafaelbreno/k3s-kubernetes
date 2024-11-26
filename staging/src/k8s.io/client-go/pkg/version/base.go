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
	gitVersion   = "v1.31.3-k3s1"
	gitCommit    = "59bb7688ff2a39c3e57df6cbead9783a281e1bf7"
	gitTreeState = "clean"
	buildDate = "2024-11-26T05:15:14Z"
)

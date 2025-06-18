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
	gitVersion   = "v1.33.2-k3s1"
	gitCommit    = "a69c0427d324d065a42b07df68157d7766eee037"
	gitTreeState = "clean"
	buildDate = "2025-06-18T19:25:24Z"
)

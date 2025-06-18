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
	gitVersion   = "v1.32.6-k3s1"
	gitCommit    = "0c40e10146069af63f1922770fd69dc78ccab2cc"
	gitTreeState = "clean"
	buildDate = "2025-06-18T20:16:50Z"
)

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
	gitVersion   = "v1.33.1-k3s1"
	gitCommit    = "477c3c458d27f70dfc44556a7558d9e9290a8349"
	gitTreeState = "clean"
	buildDate = "2025-05-15T19:08:03Z"
)

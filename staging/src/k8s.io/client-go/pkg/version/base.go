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
	gitVersion   = "v1.31.10-k3s1"
	gitCommit    = "10feaa2139778d937dcef39a16119f6983c58505"
	gitTreeState = "clean"
	buildDate = "2025-06-18T20:57:16Z"
)

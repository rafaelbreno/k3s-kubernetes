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
	gitVersion   = "v1.32.13-k3s1"
	gitCommit    = "7f7dfd49fcacd680502b826fd88ead4efc23890b"
	gitTreeState = "clean"
	buildDate = "2026-02-27T13:35:10Z"
)

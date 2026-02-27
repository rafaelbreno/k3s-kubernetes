package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.5-k3s1"
	gitCommit    = "72b9d1ba2dfd50681bacc23d9eb0492af8af001a"
	gitTreeState = "clean"
	buildDate = "2026-02-27T13:42:03Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.6-k3s1"
	gitCommit    = "15d96ccf1241b56a4ed7eadea701ec5416582f67"
	gitTreeState = "clean"
	buildDate = "2026-07-21T19:37:50Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)

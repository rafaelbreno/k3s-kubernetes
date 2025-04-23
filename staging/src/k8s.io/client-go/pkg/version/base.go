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
	gitVersion   = "v1.31.8-k3s1"
	gitCommit    = "95ad165f7554ab614d7ac0dc88f732d963e15391"
	gitTreeState = "clean"
	buildDate = "2025-04-23T16:20:38Z"
)

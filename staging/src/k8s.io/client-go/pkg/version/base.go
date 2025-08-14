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
	gitVersion   = "v1.33.4-k3s1"
	gitCommit    = "1e3a564d6217e4cee50c1f3be74407f469d4674b"
	gitTreeState = "clean"
	buildDate = "2025-08-14T14:29:56Z"
)

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
	gitVersion   = "v1.32.8-k3s1"
	gitCommit    = "e98bec04626bb8f52952f855ed73546022660dc9"
	gitTreeState = "clean"
	buildDate = "2025-08-14T14:39:38Z"
)

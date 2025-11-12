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
	gitVersion   = "v1.33.6-k3s1"
	gitCommit    = "5c66b56e6607486012b33ecace0777f8c67bc7dc"
	gitTreeState = "clean"
	buildDate = "2025-11-12T22:46:19Z"
)

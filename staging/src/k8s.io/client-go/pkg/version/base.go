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
	gitVersion   = "v1.33.13-k3s2"
	gitCommit    = "a2eaeaa8f6ae554acc83ba950fc10413c63019eb"
	gitTreeState = "clean"
	buildDate = "2026-07-22T22:35:04Z"
)

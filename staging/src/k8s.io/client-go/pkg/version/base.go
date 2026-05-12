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
	gitVersion   = "v1.33.12-k3s1"
	gitCommit    = "0cfb4ec1037be4ec17c45fd87dc2e0c3b23fc7a3"
	gitTreeState = "clean"
	buildDate = "2026-05-12T19:58:44Z"
)

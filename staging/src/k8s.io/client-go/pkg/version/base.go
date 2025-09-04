package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.0-k3s1"
	gitCommit    = "dce16ceb9bcf2d3ae5ed35078a01cb002a46697d"
	gitTreeState = "clean"
	buildDate = "2025-09-04T20:00:35Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)

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
	gitVersion   = "v1.31.9-k3s1"
	gitCommit    = "85a87cd9697e530a8157f6a518144bab35d3540b"
	gitTreeState = "clean"
	buildDate = "2025-05-15T20:23:39Z"
)

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
	gitVersion   = "v1.31.11-k3s1"
	gitCommit    = "a552d5eaa080ef48d2687169156e4993733dd87e"
	gitTreeState = "clean"
	buildDate = "2025-07-16T14:56:00Z"
)

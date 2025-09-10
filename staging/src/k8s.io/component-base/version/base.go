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
	gitVersion   = "v1.31.13-k3s1"
	gitCommit    = "84b24e8228669424bb6e871eb85162ee569cbeb0"
	gitTreeState = "clean"
	buildDate = "2025-09-10T19:14:44Z"
)

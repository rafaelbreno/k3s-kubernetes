package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.2-k3s1"
	gitCommit    = "34af2030ec6fa20992ded7c497020b9514a57bdb"
	gitTreeState = "clean"
	buildDate = "2026-02-27T13:45:55Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
